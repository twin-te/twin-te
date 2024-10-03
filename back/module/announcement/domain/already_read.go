package announcementdomain

import (
	"fmt"
	"time"

	"github.com/twin-te/twin-te/back/module/shared/domain/idtype"
)

// AlreadyRead is identified by one of the following fields.
//   - ID
//   - UserID and AnnouncementID
type AlreadyRead struct {
	ID             idtype.AlreadyReadID
	UserID         idtype.UserID
	AnnouncementID idtype.AnnouncementID
	ReadAt         time.Time
}

func ConstructAlreadyRead(fn func(ar *AlreadyRead) (err error)) (*AlreadyRead, error) {
	ar := new(AlreadyRead)
	if err := fn(ar); err != nil {
		return nil, err
	}

	if ar.ID.IsZero() ||
		ar.UserID.IsZero() ||
		ar.AnnouncementID.IsZero() ||
		ar.ReadAt.IsZero() {
		return nil, fmt.Errorf("failed to construct %+v", ar)
	}

	return ar, nil
}
