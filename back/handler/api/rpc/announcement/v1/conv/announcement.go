package announcementv1conv

import (
	"github.com/samber/lo"
	"github.com/twin-te/twin-te/back/base"
	sharedconv "github.com/twin-te/twin-te/back/handler/api/rpc/shared/conv"
	announcementv1 "github.com/twin-te/twin-te/back/handler/api/rpcgen/announcement/v1"
	announcementdomain "github.com/twin-te/twin-te/back/module/announcement/domain"
	"github.com/twin-te/twin-te/back/module/shared/domain/idtype"
)

func ToPBAnnouncement(announcement *announcementdomain.Announcement, idToReadFlag map[idtype.AnnouncementID]bool) (*announcementv1.Announcement, error) {
	pbAnnouncementTag, err := base.MapWithErr(announcement.Tags, ToPBAnnouncementTag)
	if err != nil {
		return nil, err
	}

	pbAnnouncement := &announcementv1.Announcement{
		Id:          sharedconv.ToPBUUID(announcement.ID),
		Tags:        pbAnnouncementTag,
		Title:       announcement.Title.String(),
		Content:     announcement.Content.String(),
		PublishedAt: sharedconv.ToPBRFC3339DateTime(announcement.PublishedAt),
		IsRead:      lo.Ternary(idToReadFlag == nil, nil, lo.ToPtr(idToReadFlag[announcement.ID])),
	}

	return pbAnnouncement, nil
}
