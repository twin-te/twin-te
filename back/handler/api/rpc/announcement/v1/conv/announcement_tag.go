package announcementv1conv

import (
	"fmt"

	announcementv1 "github.com/twin-te/twin-te/back/handler/api/rpcgen/announcement/v1"
	announcementdomain "github.com/twin-te/twin-te/back/module/announcement/domain"
)

func FromPBAnnouncementTag(pbAnnouncementTag announcementv1.AnnouncementTag) (announcementdomain.AnnouncementTag, error) {
	switch pbAnnouncementTag {
	case announcementv1.AnnouncementTag_ANNOUNCEMENT_TAG_INFORMATION:
		return announcementdomain.AnnouncementTagInformation, nil
	case announcementv1.AnnouncementTag_ANNOUNCEMENT_TAG_NOTIFICATION:
		return announcementdomain.AnnouncementTagNotification, nil
	}
	return 0, fmt.Errorf("invalid %v", pbAnnouncementTag)
}

func ToPBAnnouncementTag(announcementTag announcementdomain.AnnouncementTag) (announcementv1.AnnouncementTag, error) {
	switch announcementTag {
	case announcementdomain.AnnouncementTagInformation:
		return announcementv1.AnnouncementTag_ANNOUNCEMENT_TAG_INFORMATION, nil
	case announcementdomain.AnnouncementTagNotification:
		return announcementv1.AnnouncementTag_ANNOUNCEMENT_TAG_NOTIFICATION, nil
	}
	return 0, fmt.Errorf("invalid %v", announcementTag)
}
