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
	"github.com/twin-te/twin-te/back/module/shared/domain/idtype"
	sharedport "github.com/twin-te/twin-te/back/module/shared/port"
	timetablemodule "github.com/twin-te/twin-te/back/module/timetable"
	timetableappdto "github.com/twin-te/twin-te/back/module/timetable/appdto"
	timetabledomain "github.com/twin-te/twin-te/back/module/timetable/domain"
	timetableport "github.com/twin-te/twin-te/back/module/timetable/port"
)

func (uc *impl) ListCoursesByCodes(ctx context.Context, year shareddomain.AcademicYear, codes []timetabledomain.Code) ([]*timetabledomain.Course, error) {
	return uc.r.ListCourses(ctx, timetableport.CourseFilter{
		Year:  mo.Some(year),
		Codes: mo.Some(codes),
	}, sharedport.LimitOffset{}, sharedport.LockNone)
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

func (uc *impl) UpdateCoursesBasedOnKdB(ctx context.Context, year shareddomain.AcademicYear) ([]timetabledomain.Code, error) {
	if err := uc.a.Authorize(ctx, authdomain.PermissionExecuteBatchJob); err != nil {
		return nil, err
	}

	courseWithoutIDs, err := uc.i.ListCourseWithoutIDsFromKdB(ctx, year)
	if err != nil {
		return nil, err
	}

	importedCodes := make([]timetabledomain.Code, 0, len(courseWithoutIDs))

	for _, courseWithoutID := range courseWithoutIDs {
		err = uc.r.Transaction(ctx, func(rtx timetableport.Repository) error {
			courseOption, err := rtx.FindCourse(ctx, timetableport.CourseFilter{
				Year: mo.Some(year),
				Code: mo.Some(courseWithoutID.Code),
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
		}, false)
		if err != nil {
			return nil, err
		}
		importedCodes = append(importedCodes, courseWithoutID.Code)
	}

	return importedCodes, nil
}

func (uc *impl) CopyCoursesToFutureYears(ctx context.Context, sourceYear shareddomain.AcademicYear, maxFutureYears int) error {
	if err := uc.a.Authorize(ctx, authdomain.PermissionExecuteBatchJob); err != nil {
		return err
	}

	sourceCourses, err := uc.r.ListCourses(ctx, timetableport.CourseFilter{
		Year: mo.Some(sourceYear),
	}, sharedport.LimitOffset{}, sharedport.LockNone)
	if err != nil {
		return err
	}

	for futureOffset := 1; futureOffset <= maxFutureYears; futureOffset++ {
		targetYear, err := shareddomain.ParseAcademicYear(sourceYear.Int() + futureOffset)
		if err != nil {
			return err
		}

		for _, sourceCourse := range sourceCourses {
			err = uc.r.Transaction(ctx, func(rtx timetableport.Repository) error {
				courseOption, err := rtx.FindCourse(ctx, timetableport.CourseFilter{
					Year: mo.Some(targetYear),
					Code: mo.Some(sourceCourse.Code),
				}, sharedport.LockExclusive)
				if err != nil {
					return err
				}

				_, found := courseOption.Get()
				if !found {
					courseWithoutID := timetableappdto.CourseWithoutID{
						Year:              targetYear,
						Code:              sourceCourse.Code,
						Name:              sourceCourse.Name,
						Instructors:       sourceCourse.Instructors,
						Credit:            sourceCourse.Credit,
						Overview:          sourceCourse.Overview,
						Remarks:           sourceCourse.Remarks,
						LastUpdatedAt:     sourceCourse.LastUpdatedAt,
						HasParseError:     sourceCourse.HasParseError,
						IsAnnual:          sourceCourse.IsAnnual,
						RecommendedGrades: sourceCourse.RecommendedGrades,
						Methods:           sourceCourse.Methods,
						Schedules:         sourceCourse.Schedules,
					}
					newCourse, err := uc.f.NewCourse(courseWithoutID)
					if err != nil {
						return err
					}
					return rtx.CreateCourses(ctx, newCourse)
				}

				return nil
			}, false)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (uc *impl) ListMissingCourses(ctx context.Context, year shareddomain.AcademicYear, importedCodes []timetabledomain.Code) ([]*timetabledomain.Course, error) {
	if err := uc.a.Authorize(ctx, authdomain.PermissionExecuteBatchJob); err != nil {
		return nil, err
	}

	allCourses, err := uc.r.ListCourses(ctx, timetableport.CourseFilter{
		Year: mo.Some(year),
	}, sharedport.LimitOffset{}, sharedport.LockNone)
	if err != nil {
		return nil, err
	}

	importedCodeSet := lo.SliceToMap(importedCodes, func(code timetabledomain.Code) (timetabledomain.Code, struct{}) {
		return code, struct{}{}
	})

	missingCourses := lo.Filter(allCourses, func(course *timetabledomain.Course, _ int) bool {
		_, exists := importedCodeSet[course.Code]
		return !exists
	})

	return missingCourses, nil
}

func (uc *impl) MigrateMissingCourses(ctx context.Context, year shareddomain.AcademicYear, importedCodes []timetabledomain.Code) error {
	missingCourses, err := uc.ListMissingCourses(ctx, year, importedCodes)
	if err != nil {
		return err
	}

	for _, course := range missingCourses {
		err = uc.r.Transaction(ctx, func(rtx timetableport.Repository) error {
			registeredCourses, err := rtx.ListRegisteredCourses(ctx, timetableport.RegisteredCourseFilter{
				CourseIDs: mo.Some([]idtype.CourseID{course.ID}),
			}, sharedport.LimitOffset{}, sharedport.LockExclusive)
			if err != nil {
				return err
			}

			for _, rc := range registeredCourses {
				rc.BeforeUpdateHook()

				deprecatedName, err := timetabledomain.ParseName("【移行失敗】" + course.Name.String())
				if err != nil {
					return err
				}

				rc.CourseID = mo.None[idtype.CourseID]()
				rc.Name = mo.Some(deprecatedName)
				rc.Instructors = mo.Some(course.Instructors)
				rc.Credit = mo.Some(course.Credit)
				rc.Methods = mo.Some(base.CopySlice(course.Methods))
				rc.Schedules = mo.Some(base.CopySlice(course.Schedules))
				if rc.Memo != "" {
					rc.Memo = rc.Memo + "\n"
				}
				rc.Memo = rc.Memo + "元の科目番号: " + course.Code.String()

				if err := rtx.UpdateRegisteredCourse(ctx, rc); err != nil {
					return err
				}
			}

			_, err = rtx.DeleteCourses(ctx, timetableport.CourseFilter{
				ID: mo.Some(course.ID),
			})
			return err
		}, false)
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

	courses, err = uc.r.ListCourses(ctx, timetableport.CourseFilter{
		Year: mo.Some(year),
	}, sharedport.LimitOffset{}, sharedport.LockNone)
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
