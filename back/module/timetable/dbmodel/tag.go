package timetabledbmodel

import (
	"time"

	"github.com/twin-te/twin-te/back/module/shared/domain/idtype"
	timetabledomain "github.com/twin-te/twin-te/back/module/timetable/domain"
)

type Tag struct {
	ID     string
	UserID string
	Name   string
	Order  int

	CreatedAt time.Time
	UpdatedAt time.Time
}

func FromDBTag(dbTag *Tag) (*timetabledomain.Tag, error) {
	return timetabledomain.ConstructTag(func(t *timetabledomain.Tag) (err error) {
		t.ID, err = idtype.ParseTagID(dbTag.ID)
		if err != nil {
			return err
		}

		t.UserID, err = idtype.ParseUserID(dbTag.UserID)
		if err != nil {
			return err
		}

		t.Name, err = timetabledomain.ParseName(dbTag.Name)
		if err != nil {
			return err
		}

		t.Order, err = timetabledomain.ParseOrder(int(dbTag.Order))
		if err != nil {
			return err
		}

		return nil
	})
}

func ToDBTag(tag *timetabledomain.Tag) *Tag {
	return &Tag{
		ID:     tag.ID.String(),
		UserID: tag.UserID.String(),
		Name:   tag.Name.String(),
		Order:  tag.Order.Int(),
	}
}
