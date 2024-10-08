package announcementv1svc

import (
	"context"

	"github.com/bufbuild/connect-go"

	"github.com/twin-te/twin-te/back/base"
	announcementv1conv "github.com/twin-te/twin-te/back/handler/api/rpc/announcement/v1/conv"
	sharedconv "github.com/twin-te/twin-te/back/handler/api/rpc/shared/conv"
	announcementv1 "github.com/twin-te/twin-te/back/handler/api/rpcgen/announcement/v1"
	"github.com/twin-te/twin-te/back/handler/api/rpcgen/announcement/v1/announcementv1connect"
	announcementmodule "github.com/twin-te/twin-te/back/module/announcement"
	"github.com/twin-te/twin-te/back/module/shared/domain/idtype"
)

var _ announcementv1connect.AnnouncementServiceHandler = (*impl)(nil)

type impl struct {
	uc announcementmodule.UseCase
}

func (svc *impl) ListAnnouncements(ctx context.Context, req *connect.Request[announcementv1.ListAnnouncementsRequest]) (res *connect.Response[announcementv1.ListAnnouncementsResponse], err error) {
	announcements, idToReadFlag, err := svc.uc.ListAnnouncements(ctx)
	if err != nil {
		return
	}

	pbAnnouncements, err := base.MapWithArgAndErr(announcements, idToReadFlag, announcementv1conv.ToPBAnnouncement)
	if err != nil {
		return
	}

	res = connect.NewResponse(&announcementv1.ListAnnouncementsResponse{
		Announcements: pbAnnouncements,
	})

	return
}

func (svc *impl) ReadAnnouncements(ctx context.Context, req *connect.Request[announcementv1.ReadAnnouncementsRequest]) (res *connect.Response[announcementv1.ReadAnnouncementsResponse], err error) {
	ids, err := base.MapWithArgAndErr(req.Msg.Ids, idtype.ParseAnnouncementID, sharedconv.FromPBUUID)
	if err != nil {
		return
	}

	if err = svc.uc.ReadAnnouncements(ctx, ids); err != nil {
		return
	}

	res = connect.NewResponse(&announcementv1.ReadAnnouncementsResponse{})

	return
}

func New(uc announcementmodule.UseCase) *impl {
	return &impl{uc: uc}
}
