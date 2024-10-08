// Code generated by codegen/idtype/generate.py; DO NOT EDIT.

package idtype

import (
	"database/sql/driver"	
	"fmt"

	"github.com/google/uuid"
)

type AnnouncementID uuid.UUID

func (id AnnouncementID) String() string {
	return uuid.UUID(id).String()
}

func (id AnnouncementID) IsZero() bool {
	return uuid.UUID(id) == uuid.Nil
}

func (id AnnouncementID) Less(other AnnouncementID) bool {
	for i := 0; i < 16; i++ {
		if id[i] == other[i] {
			continue
		}
		return id[i] < other[i]
	}
	return false
}

func (id *AnnouncementID) Scan(src interface{}) error {
	uuid := new(uuid.UUID)
	if err := uuid.Scan(src); err != nil {
		return err
	}
	*id = AnnouncementID(*uuid)
	return nil
}

func (id AnnouncementID) Value() (driver.Value, error) {
	return id.String(), nil
}

func NewAnnouncementID() AnnouncementID {
	return AnnouncementID(uuid.New())
}

func ParseAnnouncementID(s string) (AnnouncementID, error) {
	id, err := uuid.Parse(s)
	if err != nil {
		return AnnouncementID{}, fmt.Errorf("failed to parse AnnouncementID %v", s)
	}
	return AnnouncementID(id), nil
}
