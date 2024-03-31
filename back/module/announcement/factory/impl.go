package announcementfactory

import (
	"time"

	announcementdomain "github.com/twin-te/twinte-back/module/announcement/domain"
	announcementport "github.com/twin-te/twinte-back/module/announcement/port"
	"github.com/twin-te/twinte-back/module/shared/domain/idtype"
)

var _ announcementport.Factory = (*impl)(nil)

type impl struct {
	nowFunc func() time.Time
}

func (f *impl) NewAlreadyRead(userID idtype.UserID, announcementID idtype.AnnouncementID) (*announcementdomain.AlreadyRead, error) {
	return announcementdomain.ConstructAlreadyRead(func(ar *announcementdomain.AlreadyRead) (err error) {
		ar.ID = idtype.NewAlreadyReadID()
		ar.UserID = userID
		ar.AnnouncementID = announcementID
		ar.ReadAt = f.nowFunc()
		return nil
	})
}

func New(nowFunc func() time.Time) *impl {
	return &impl{nowFunc: nowFunc}
}
