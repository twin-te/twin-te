package announcementport

import (
	announcementdomain "github.com/twin-te/twinte-back/module/announcement/domain"
	"github.com/twin-te/twinte-back/module/shared/domain/idtype"
)

type Factory interface {
	NewAlreadyRead(userID idtype.UserID, announcementID idtype.AnnouncementID) (*announcementdomain.AlreadyRead, error)
}
