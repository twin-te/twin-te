package timetableusecase

import (
	"context"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/samber/lo"
	"github.com/samber/mo"
	"github.com/twin-te/twin-te/back/appenv"
	"github.com/twin-te/twin-te/back/base"
	authdomain "github.com/twin-te/twin-te/back/module/auth/domain"
	shareddomain "github.com/twin-te/twin-te/back/module/shared/domain"
	sharedport "github.com/twin-te/twin-te/back/module/shared/port"
	timetablemodule "github.com/twin-te/twin-te/back/module/timetable"
	timetabledomain "github.com/twin-te/twin-te/back/module/timetable/domain"
	timetableport "github.com/twin-te/twin-te/back/module/timetable/port"
)

func (uc *impl) ListCoursesByCodes(ctx context.Context, year shareddomain.AcademicYear, codes []timetabledomain.Code) ([]*timetabledomain.Course, error) {
	return uc.r.ListCourses(ctx, timetableport.ListCoursesConds{
		Year:  mo.Some(year),
		Codes: mo.Some(codes),
	}, sharedport.LockNone)
}

func (uc *impl) SearchCourses(ctx context.Context, conds timetablemodule.SearchCoursesIn) ([]*timetabledomain.Course, error) {
	courses, err := uc.getCoursesWithCache(ctx, conds.Year)
	if err != nil {
		return nil, err
	}

	// Filter by keywords
	if len(conds.Keywords) != 0 {
		courses = lo.Filter(courses, func(course *timetabledomain.Course, _ int) bool {
			return lo.EveryBy(conds.Keywords, func(keyword string) bool {
				return strings.Contains(course.Name.String(), keyword)
			})
		})
	}

	// Filter by code prefixes
	if len(conds.CodePrefixes.Included) != 0 {
		courses = lo.Filter(courses, func(course *timetabledomain.Course, _ int) bool {
			return lo.EveryBy(conds.CodePrefixes.Included, func(code string) bool {
				return strings.HasPrefix(course.Code.String(), code)
			})
		})
	}
	if len(conds.CodePrefixes.Excluded) != 0 {
		courses = lo.Filter(courses, func(course *timetabledomain.Course, _ int) bool {
			return lo.EveryBy(conds.CodePrefixes.Excluded, func(code string) bool {
				return !strings.HasPrefix(course.Code.String(), code)
			})
		})
	}

	// Filter by schedules
	if len(conds.Schedules.FullyIncluded) != 0 {
		courses = lo.Filter(courses, func(course *timetabledomain.Course, _ int) bool {
			return lo.EveryBy(course.Schedules, func(s1 timetabledomain.Schedule) bool {
				return lo.SomeBy(conds.Schedules.FullyIncluded, func(s2 timetabledomain.Schedule) bool {
					return s1.Module == s2.Module && s1.Day == s2.Day && s1.Period == s2.Period
				})
			})
		})
	}
	if len(conds.Schedules.PartiallyOverlapped) != 0 {
		courses = lo.Filter(courses, func(course *timetabledomain.Course, _ int) bool {
			return lo.SomeBy(course.Schedules, func(s1 timetabledomain.Schedule) bool {
				return lo.SomeBy(conds.Schedules.PartiallyOverlapped, func(s2 timetabledomain.Schedule) bool {
					return s1.Module == s2.Module && s1.Day == s2.Day && s1.Period == s2.Period
				})
			})
		})
	}

	// Sort by code
	sort.Slice(courses, func(i, j int) bool {
		return courses[i].Code.String() < courses[j].Code.String()
	})

	// Apply offset
	courses = courses[lo.Clamp(conds.Offset, 0, len(courses)):]

	// Apply limit
	courses = courses[:lo.Clamp(conds.Limit, 0, len(courses))]

	return base.MapByClone(courses), nil
}

func (uc *impl) UpdateCoursesBasedOnKdB(ctx context.Context, year shareddomain.AcademicYear) error {
	if err := uc.a.Authorize(ctx, authdomain.PermissionExecuteBatchJob); err != nil {
		return err
	}

	courseWithoutIDs, err := uc.i.ListCourseWithoutIDsFromKdB(ctx, year)
	if err != nil {
		return err
	}

	for _, courseWithoutID := range courseWithoutIDs {
		err = uc.r.Transaction(ctx, func(rtx timetableport.Repository) error {
			courseOption, err := rtx.FindCourse(ctx, timetableport.FindCourseConds{
				Year: year,
				Code: courseWithoutID.Code,
			}, sharedport.LockExclusive)
			if err != nil {
				return err
			}

			course, found := courseOption.Get()
			if !found {
				newCourse, err := uc.f.NewCourse(courseWithoutID)
				if err != nil {
					return err
				}
				return rtx.CreateCourses(ctx, newCourse)
			}

			if courseWithoutID.LastUpdatedAt.After(course.LastUpdatedAt) {
				course.BeforeUpdateHook()
				course.Update(timetabledomain.CourseDataToUpdate{
					Name:              mo.Some(courseWithoutID.Name),
					Instructors:       mo.Some(courseWithoutID.Instructors),
					Credit:            mo.Some(courseWithoutID.Credit),
					Overview:          mo.Some(courseWithoutID.Overview),
					Remarks:           mo.Some(courseWithoutID.Remarks),
					LastUpdatedAt:     mo.Some(courseWithoutID.LastUpdatedAt),
					HasParseError:     mo.Some(courseWithoutID.HasParseError),
					IsAnnual:          mo.Some(courseWithoutID.IsAnnual),
					RecommendedGrades: mo.Some(courseWithoutID.RecommendedGrades),
					Methods:           mo.Some(courseWithoutID.Methods),
					Schedules:         mo.Some(courseWithoutID.Schedules),
				})
				return rtx.UpdateCourse(ctx, course)
			}

			return nil
		})
		if err != nil {
			return err
		}
	}

	return nil
}

var (
	courseCache         = make(map[shareddomain.AcademicYear][]*timetabledomain.Course)
	courseCacheMutex    sync.Mutex
	courseCacheTime     time.Duration = time.Duration(appenv.COURSE_CACHE_HOURS) * time.Hour
	courseCacheCapacity               = 100_000
)

func (uc *impl) getCoursesWithCache(ctx context.Context, year shareddomain.AcademicYear) (courses []*timetabledomain.Course, err error) {
	courseCacheMutex.Lock()
	defer courseCacheMutex.Unlock()

	courses, ok := courseCache[year]
	if ok {
		return
	}

	courses, err = uc.r.ListCourses(ctx, timetableport.ListCoursesConds{
		Year: mo.Some(year),
	}, sharedport.LockNone)
	if err != nil {
		return
	}

	courseCache[year] = courses

	for len(courseCache) != 0 {
		totalNumCourses := len(lo.Flatten(lo.Values(courseCache)))

		if totalNumCourses <= courseCacheCapacity {
			break
		}

		oldestYear := lo.Min(lo.Keys(courseCache))
		delete(courseCache, oldestYear)
	}

	go func() {
		time.Sleep(courseCacheTime)
		courseCacheMutex.Lock()
		delete(courseCache, year)
		courseCacheMutex.Unlock()
	}()

	return
}
