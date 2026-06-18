package calendarusecase

import (
	"bytes"
	"context"

	"github.com/samber/mo"
	calendardomain "github.com/twin-te/twin-te/back/module/calendar/domain"
	shareddomain "github.com/twin-te/twin-te/back/module/shared/domain"
	"github.com/twin-te/twin-te/back/module/shared/domain/idtype"
	timetableappdto "github.com/twin-te/twin-te/back/module/timetable/appdto"
	timetableport "github.com/twin-te/twin-te/back/module/timetable/port"
)

func filterCoursesByTags(courses []*timetableappdto.RegisteredCourse, tagIDs []idtype.TagID) []*timetableappdto.RegisteredCourse {
	if len(tagIDs) == 0 {
		return courses
	}
	m := make(map[idtype.TagID]struct{}, len(tagIDs))
	for _, t := range tagIDs {
		m[t] = struct{}{}
	}
	filtered := make([]*timetableappdto.RegisteredCourse, 0, len(courses))
	for _, c := range courses {
		for _, t := range c.TagIDs {
			if _, ok := m[t]; ok {
				filtered = append(filtered, c)
				break
			}
		}
	}
	return filtered
}

func (uc *impl) ExportTimetableToICal(ctx context.Context, year shareddomain.AcademicYear, tagIDs []idtype.TagID, isRdateSupported bool) ([]byte, error) {
	userID, err := uc.a.Authenticate(ctx)
	if err != nil {
		return nil, err
	}

	return uc.exportTimetableToICalByUserID(ctx, userID, year, tagIDs, nil, isRdateSupported)
}

func (uc *impl) ExportTimetableToICalByIcalSubscriptionID(ctx context.Context, id idtype.IcalSubscriptionID, year shareddomain.AcademicYear, tagIDs []idtype.TagID, isRdateSupported bool) ([]byte, error) {
	sub, err := uc.resolveIcalSubscriptionByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return uc.exportTimetableToICalByUserID(ctx, sub.UserID, year, tagIDs, sub, isRdateSupported)
}

func (uc *impl) exportTimetableToICalByUserID(ctx context.Context, userID idtype.UserID, year shareddomain.AcademicYear, tagIDs []idtype.TagID, sub *calendardomain.IcalSubscription, isRdateSupported bool) ([]byte, error) {
	modules, err := uc.buildSchoolCalendarModules(ctx, year)
	if err != nil {
		return nil, err
	}

	courses, err := uc.timetable.ListRegisteredCourses(ctx, timetableport.ListRegisteredCoursesConds{
		UserID: mo.Some(userID),
		Year:   mo.Some(year),
	})
	if err != nil {
		return nil, err
	}

	// ?tags[]= クエリパラメータは包含フィルタとして先に適用する。
	courses = filterCoursesByTags(courses, tagIDs)

	// 連携のモード（EXCLUDE / TRANSPARENT）は包含フィルタの後段で適用する。
	courses, transparentCourseIDs := applyIcalSubscriptionMode(courses, sub)

	var buf bytes.Buffer
	if err := uc.writeICalendar(&buf, modules, courses, transparentCourseIDs, isRdateSupported); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// applyIcalSubscriptionMode は連携のモードをコース一覧に適用する。
//   - EXCLUDE: タグが対象タグと交差するコースを除外する。
//   - TRANSPARENT: 該当コースは残しつつ、TRANSP:TRANSPARENT を付与できるよう返却用の集合に記録する。
//   - SYNC / 連携が nil（認証ユーザーによる直接エクスポート）: コースをそのまま返す。
func applyIcalSubscriptionMode(courses []*timetableappdto.RegisteredCourse, sub *calendardomain.IcalSubscription) ([]*timetableappdto.RegisteredCourse, map[idtype.RegisteredCourseID]struct{}) {
	if sub == nil || sub.Mode == calendardomain.IcalSubscriptionModeSync {
		return courses, nil
	}

	transparent := make(map[idtype.RegisteredCourseID]struct{})
	filtered := make([]*timetableappdto.RegisteredCourse, 0, len(courses))
	for _, c := range courses {
		if sub.IsExcluded(c.TagIDs) {
			continue
		}
		if sub.IsTransparent(c.TagIDs) {
			transparent[c.ID] = struct{}{}
		}
		filtered = append(filtered, c)
	}
	return filtered, transparent
}
