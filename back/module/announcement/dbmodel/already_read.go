package announcementdbmodel

import (
	"time"

	announcementdomain "github.com/twin-te/twin-te/back/module/announcement/domain"
	"github.com/twin-te/twin-te/back/module/shared/domain/idtype"
)

type AlreadyRead struct {
	ID             string
	UserID         string
	AnnouncementID string
	ReadAt         time.Time

	CreatedAt time.Time
	UpdatedAt time.Time
}

func FromDBAlreadyRead(dbAlreadyRead *AlreadyRead) (*announcementdomain.AlreadyRead, error) {
	return announcementdomain.ConstructAlreadyRead(func(ar *announcementdomain.AlreadyRead) (err error) {
		ar.ID, err = idtype.ParseAlreadyReadID(dbAlreadyRead.ID)
		if err != nil {
			return
		}

		ar.UserID, err = idtype.ParseUserID(dbAlreadyRead.UserID)
		if err != nil {
			return
		}

		ar.AnnouncementID, err = idtype.ParseAnnouncementID(dbAlreadyRead.AnnouncementID)
		if err != nil {
			return
		}

		ar.ReadAt = dbAlreadyRead.ReadAt

		return
	})
}

func ToDBAlreadyRead(alreadyRead *announcementdomain.AlreadyRead) *AlreadyRead {
	return &AlreadyRead{
		ID:             alreadyRead.ID.String(),
		UserID:         alreadyRead.UserID.String(),
		AnnouncementID: alreadyRead.AnnouncementID.String(),
		ReadAt:         alreadyRead.ReadAt,
	}
}
