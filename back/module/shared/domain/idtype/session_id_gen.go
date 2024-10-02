// Code generated by codegen/idtype/generate.py; DO NOT EDIT.

package idtype

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/samber/lo"
)

type SessionID uuid.UUID

func (id SessionID) String() string {
	return uuid.UUID(id).String()
}

func (id SessionID) IsZero() bool {
	return uuid.UUID(id) == uuid.Nil
}

func (id SessionID) Less(other SessionID) bool {
	for i := 0; i < 16; i++ {
		if id[i] == other[i] {
			continue
		}
		return id[i] < other[i]
	}
	return false
}

func (id *SessionID) StringPtr() *string {
	if id == nil {
		return nil
	}
	return lo.ToPtr(id.String())
}

func NewSessionID() SessionID {
	return SessionID(uuid.New())
}

func ParseSessionID(s string) (SessionID, error) {
	id, err := uuid.Parse(s)
	if err != nil {
		return SessionID{}, fmt.Errorf("failed to parse SessionID %v", s)
	}
	return SessionID(id), nil
}
