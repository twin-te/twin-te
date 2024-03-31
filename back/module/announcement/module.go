package announcementmodule

import (
	"context"

	announcementdomain "github.com/twin-te/twinte-back/module/announcement/domain"
	"github.com/twin-te/twinte-back/module/shared/domain/idtype"
)

// UseCase represents application specific business rules.
//
// The following error codes are not stated explicitly in the each method, but may be returned.
//   - shared.InvalidArgument
//   - shared.Unauthenticated
//   - shared.Unauthorized
type UseCase interface {
	// GetAnnouncements returns all published announcements.
	// If authenticated, idToReadFlag is also be returned.
	//
	// [Authentication] optional
	GetAnnouncements(ctx context.Context) (announcements []*announcementdomain.Announcement, idToReadFlag map[idtype.AnnouncementID]bool, err error)

	// ReadAnnouncements means that the user read the announcements specified by the given ids.
	//
	// [Error Code]
	//   - announcement.AnnouncementNotFound
	//
	// [Authentication] required
	ReadAnnouncements(ctx context.Context, ids []idtype.AnnouncementID) error
}
