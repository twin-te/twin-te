package timetablev1svc

import (
	"context"

	"github.com/bufbuild/connect-go"
	"github.com/samber/mo"
	"github.com/twin-te/twin-te/back/base"
	sharedconv "github.com/twin-te/twin-te/back/handler/api/rpc/shared/conv"
	timetablev1conv "github.com/twin-te/twin-te/back/handler/api/rpc/timetable/v1/conv"
	timetablev1 "github.com/twin-te/twin-te/back/handler/api/rpcgen/timetable/v1"
	"github.com/twin-te/twin-te/back/handler/api/rpcgen/timetable/v1/timetablev1connect"
	shareddomain "github.com/twin-te/twin-te/back/module/shared/domain"
	"github.com/twin-te/twin-te/back/module/shared/domain/idtype"
	timetablemodule "github.com/twin-te/twin-te/back/module/timetable"
	timetabledomain "github.com/twin-te/twin-te/back/module/timetable/domain"
)

var _ timetablev1connect.TimetableServiceHandler = (*impl)(nil)

type impl struct {
	uc timetablemodule.UseCase
}

func (svc *impl) ListCoursesByCodes(ctx context.Context, req *connect.Request[timetablev1.ListCoursesByCodesRequest]) (res *connect.Response[timetablev1.ListCoursesByCodesResponse], err error) {
	year, err := sharedconv.FromPBAcadimicYear(req.Msg.Year)
	if err != nil {
		return
	}

	codes, err := base.MapWithErr(req.Msg.Codes, timetabledomain.ParseCode)
	if err != nil {
		return
	}

	courses, err := svc.uc.ListCoursesByCodes(ctx, year, codes)
	if err != nil {
		return
	}

	pbCourses, err := base.MapWithErr(courses, timetablev1conv.ToPBCourse)
	if err != nil {
		return
	}

	res = connect.NewResponse(&timetablev1.ListCoursesByCodesResponse{
		Courses: pbCourses,
	})

	return
}

func (svc *impl) SearchCourses(ctx context.Context, req *connect.Request[timetablev1.SearchCoursesRequest]) (res *connect.Response[timetablev1.SearchCoursesResponse], err error) {
	year, err := sharedconv.FromPBAcadimicYear(req.Msg.Year)
	if err != nil {
		return
	}

	in := timetablemodule.SearchCoursesIn{
		Year:     year,
		Keywords: req.Msg.Keywords,
		CodePrefixes: struct {
			Included []string
			Excluded []string
		}{
			Included: req.Msg.CodePrefixesIncluded,
			Excluded: req.Msg.CodePrefixesExcluded,
		},
		Limit:  int(req.Msg.Limit),
		Offset: int(req.Msg.Offset),
	}

	in.Schedules.FullyIncluded, err = base.MapWithErr(req.Msg.SchedulesFullyIncluded, timetablev1conv.FromPBSchedule)
	if err != nil {
		return
	}

	in.Schedules.PartiallyOverlapped, err = base.MapWithErr(req.Msg.SchedulesPartiallyOverlapped, timetablev1conv.FromPBSchedule)
	if err != nil {
		return
	}

	courses, err := svc.uc.SearchCourses(ctx, in)
	if err != nil {
		return
	}

	pbCourses, err := base.MapWithErr(courses, timetablev1conv.ToPBCourse)
	if err != nil {
		return
	}

	res = connect.NewResponse(&timetablev1.SearchCoursesResponse{
		Courses: pbCourses,
	})

	return
}

func (svc *impl) CreateRegisteredCoursesByCodes(ctx context.Context, req *connect.Request[timetablev1.CreateRegisteredCoursesByCodesRequest]) (res *connect.Response[timetablev1.CreateRegisteredCoursesByCodesResponse], err error) {
	year, err := sharedconv.FromPBAcadimicYear(req.Msg.Year)
	if err != nil {
		return
	}

	codes, err := base.MapWithErr(req.Msg.Codes, timetabledomain.ParseCode)
	if err != nil {
		return
	}

	registeredCourses, err := svc.uc.CreateRegisteredCoursesByCodes(ctx, year, codes)
	if err != nil {
		return
	}

	pbRegisteredCourses, err := base.MapWithErr(registeredCourses, timetablev1conv.ToPBRegisteredCourse)
	if err != nil {
		return
	}

	res = connect.NewResponse(&timetablev1.CreateRegisteredCoursesByCodesResponse{
		RegisteredCourses: pbRegisteredCourses,
	})

	return
}

func (svc *impl) CreateRegisteredCourseManually(ctx context.Context, req *connect.Request[timetablev1.CreateRegisteredCourseManuallyRequest]) (res *connect.Response[timetablev1.CreateRegisteredCourseManuallyResponse], err error) {
	year, err := sharedconv.FromPBAcadimicYear(req.Msg.Year)
	if err != nil {
		return
	}

	name, err := timetabledomain.ParseName(req.Msg.Name)
	if err != nil {
		return
	}

	credit, err := timetabledomain.ParseCredit(req.Msg.Credit)
	if err != nil {
		return
	}

	methods, err := base.MapWithErr(req.Msg.Methods, timetablev1conv.FromPBCourseMethod)
	if err != nil {
		return
	}

	schedules, err := base.MapWithErr(req.Msg.Schedules, timetablev1conv.FromPBSchedule)
	if err != nil {
		return
	}

	in := timetablemodule.CreateRegisteredCourseManuallyIn{
		Year:        year,
		Name:        name,
		Instructors: req.Msg.Instructors,
		Credit:      credit,
		Methods:     methods,
		Schedules:   schedules,
	}

	registeredCourse, err := svc.uc.CreateRegisteredCourseManually(ctx, in)
	if err != nil {
		return
	}

	pbRegisteredCourse, err := timetablev1conv.ToPBRegisteredCourse(registeredCourse)
	if err != nil {
		return
	}

	res = connect.NewResponse(&timetablev1.CreateRegisteredCourseManuallyResponse{
		RegisteredCourse: pbRegisteredCourse,
	})

	return
}

func (svc *impl) ListRegisteredCourses(ctx context.Context, req *connect.Request[timetablev1.ListRegisteredCoursesRequest]) (res *connect.Response[timetablev1.ListRegisteredCoursesResponse], err error) {
	var year mo.Option[shareddomain.AcademicYear]
	if req.Msg.Year != nil {
		year, err = base.OptionMapWithErr(mo.Some(req.Msg.Year), sharedconv.FromPBAcadimicYear)
		if err != nil {
			return
		}
	}

	registeredCourses, err := svc.uc.ListRegisteredCourses(ctx, year)
	if err != nil {
		return
	}

	pbRegisteredCourses, err := base.MapWithErr(registeredCourses, timetablev1conv.ToPBRegisteredCourse)
	if err != nil {
		return
	}

	res = connect.NewResponse(&timetablev1.ListRegisteredCoursesResponse{
		RegisteredCourses: pbRegisteredCourses,
	})

	return
}

func (svc *impl) UpdateRegisteredCourse(ctx context.Context, req *connect.Request[timetablev1.UpdateRegisteredCourseRequest]) (res *connect.Response[timetablev1.UpdateRegisteredCourseResponse], err error) {
	in := timetablemodule.UpdateRegisteredCourseIn{
		Instructors: mo.PointerToOption(req.Msg.Instructors),
		Memo:        mo.PointerToOption(req.Msg.Memo),
	}

	in.ID, err = sharedconv.FromPBUUID(req.Msg.Id, idtype.ParseRegisteredCourseID)
	if err != nil {
		return
	}

	if req.Msg.Name != nil {
		in.Name, err = base.SomeWithErr(timetabledomain.ParseName(*req.Msg.Name))
		if err != nil {
			return
		}
	}

	if req.Msg.Credit != nil {
		in.Credit, err = base.SomeWithErr(timetabledomain.ParseCredit(*req.Msg.Credit))
		if err != nil {
			return
		}
	}

	if req.Msg.Methods != nil {
		in.Methods, err = base.SomeWithErr(base.MapWithErr(req.Msg.Methods.Values, timetablev1conv.FromPBCourseMethod))
		if err != nil {
			return
		}
	}

	if req.Msg.Schedules != nil {
		in.Schedules, err = base.SomeWithErr(base.MapWithErr(req.Msg.Schedules.Values, timetablev1conv.FromPBSchedule))
		if err != nil {
			return
		}
	}

	if req.Msg.Attendance != nil {
		in.Attendance, err = base.SomeWithErr(timetabledomain.ParseAttendance(int(*req.Msg.Attendance)))
		if err != nil {
			return
		}
	}

	if req.Msg.Late != nil {
		in.Late, err = base.SomeWithErr(timetabledomain.ParseLate(int(*req.Msg.Late)))
		if err != nil {
			return
		}
	}

	if req.Msg.Absence != nil {
		in.Absence, err = base.SomeWithErr(timetabledomain.ParseAbsence(int(*req.Msg.Absence)))
		if err != nil {
			return
		}
	}

	if req.Msg.TagIds != nil {
		in.TagIDs, err = base.SomeWithErr(base.MapWithArgAndErr(req.Msg.TagIds.Values, idtype.ParseTagID, sharedconv.FromPBUUID[idtype.TagID]))
		if err != nil {
			return
		}
	}

	registeredCourse, err := svc.uc.UpdateRegisteredCourse(ctx, in)
	if err != nil {
		return
	}

	pbRegisteredCourse, err := timetablev1conv.ToPBRegisteredCourse(registeredCourse)
	if err != nil {
		return
	}

	res = connect.NewResponse(&timetablev1.UpdateRegisteredCourseResponse{
		RegisteredCourse: pbRegisteredCourse,
	})

	return
}

func (svc *impl) DeleteRegisteredCourse(ctx context.Context, req *connect.Request[timetablev1.DeleteRegisteredCourseRequest]) (res *connect.Response[timetablev1.DeleteRegisteredCourseResponse], err error) {
	id, err := sharedconv.FromPBUUID(req.Msg.Id, idtype.ParseRegisteredCourseID)
	if err != nil {
		return
	}

	if err = svc.uc.DeleteRegisteredCourse(ctx, id); err != nil {
		return
	}

	res = connect.NewResponse(&timetablev1.DeleteRegisteredCourseResponse{})

	return
}

func (svc *impl) CreateTag(ctx context.Context, req *connect.Request[timetablev1.CreateTagRequest]) (res *connect.Response[timetablev1.CreateTagResponse], err error) {
	name, err := timetabledomain.ParseName(req.Msg.Name)
	if err != nil {
		return
	}

	tag, err := svc.uc.CreateTag(ctx, name)
	if err != nil {
		return
	}

	pbTag := timetablev1conv.ToPBTag(tag)

	res = connect.NewResponse(&timetablev1.CreateTagResponse{
		Tag: pbTag,
	})

	return
}

func (svc *impl) ListTags(ctx context.Context, req *connect.Request[timetablev1.ListTagsRequest]) (res *connect.Response[timetablev1.ListTagsResponse], err error) {
	tags, err := svc.uc.ListTags(ctx)
	if err != nil {
		return
	}

	pbTags := base.Map(tags, timetablev1conv.ToPBTag)

	res = connect.NewResponse(&timetablev1.ListTagsResponse{
		Tags: pbTags,
	})

	return
}

func (svc *impl) UpdateTag(ctx context.Context, req *connect.Request[timetablev1.UpdateTagRequest]) (res *connect.Response[timetablev1.UpdateTagResponse], err error) {
	in := timetablemodule.UpdateTagIn{}

	in.ID, err = sharedconv.FromPBUUID(req.Msg.Id, idtype.ParseTagID)
	if err != nil {
		return
	}

	if req.Msg.Name != nil {
		in.Name, err = base.SomeWithErr(timetabledomain.ParseName(*req.Msg.Name))
		if err != nil {
			return
		}
	}

	tag, err := svc.uc.UpdateTag(ctx, in)
	if err != nil {
		return
	}

	pbTag := timetablev1conv.ToPBTag(tag)

	res = connect.NewResponse(&timetablev1.UpdateTagResponse{
		Tag: pbTag,
	})

	return
}

func (svc *impl) DeleteTag(ctx context.Context, req *connect.Request[timetablev1.DeleteTagRequest]) (res *connect.Response[timetablev1.DeleteTagResponse], err error) {
	id, err := sharedconv.FromPBUUID(req.Msg.Id, idtype.ParseTagID)
	if err != nil {
		return
	}

	if err = svc.uc.DeleteTag(ctx, id); err != nil {
		return
	}

	res = connect.NewResponse(&timetablev1.DeleteTagResponse{})

	return
}

func (svc *impl) RearrangeTags(ctx context.Context, req *connect.Request[timetablev1.RearrangeTagsRequest]) (res *connect.Response[timetablev1.RearrangeTagsResponse], err error) {
	ids, err := base.MapWithArgAndErr(req.Msg.Ids, idtype.ParseTagID, sharedconv.FromPBUUID[idtype.TagID])
	if err != nil {
		return
	}

	tags, err := svc.uc.RearrangeTags(ctx, ids)
	if err != nil {
		return
	}

	pbTags := base.Map(tags, timetablev1conv.ToPBTag)

	res = connect.NewResponse(&timetablev1.RearrangeTagsResponse{
		Tags: pbTags,
	})

	return
}

func New(uc timetablemodule.UseCase) *impl {
	return &impl{uc: uc}
}
