// Code generated by codegen/idtype/generate.py; DO NOT EDIT.

package idtype

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/samber/lo"
)

type TagID uuid.UUID

func (id TagID) String() string {
	return uuid.UUID(id).String()
}

func (id TagID) IsZero() bool {
	return uuid.UUID(id) == uuid.Nil
}

func (id TagID) Less(other TagID) bool {
	for i := 0; i < 16; i++ {
		if id[i] == other[i] {
			continue
		}
		return id[i] < other[i]
	}
	return false
}

func (id *TagID) StringPtr() *string {
	if id == nil {
		return nil
	}
	return lo.ToPtr(id.String())
}

func NewTagID() TagID {
	return TagID(uuid.New())
}

func ParseTagID(s string) (TagID, error) {
	id, err := uuid.Parse(s)
	if err != nil {
		return TagID{}, fmt.Errorf("failed to parse TagID %v", s)
	}
	return TagID(id), nil
}
