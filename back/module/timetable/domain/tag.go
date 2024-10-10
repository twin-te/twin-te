package timetabledomain

import (
	"fmt"

	"github.com/samber/lo"
	"github.com/samber/mo"
	shareddomain "github.com/twin-te/twin-te/back/module/shared/domain"
	"github.com/twin-te/twin-te/back/module/shared/domain/idtype"
)

var ParseOrder = shareddomain.NewNonNegativeIntParser("order")

// Tag is identified by one of the following fields.
//   - ID
//   - UserID and Order
type Tag struct {
	ID     idtype.TagID
	UserID idtype.UserID
	Name   shareddomain.RequiredString
	Order  shareddomain.NonNegativeInt

	BeforeUpdated mo.Option[*Tag]
}

func (t *Tag) Clone() *Tag {
	ret := lo.ToPtr(*t)
	return ret
}

func (t *Tag) BeforeUpdateHook() {
	t.BeforeUpdated = mo.Some(t.Clone())
}

type TagDataToUpdate struct {
	Name mo.Option[shareddomain.RequiredString]
}

func (t *Tag) Update(data TagDataToUpdate) {
	if name, ok := data.Name.Get(); ok {
		t.Name = name
	}
}

func ConstructTag(fn func(t *Tag) (err error)) (*Tag, error) {
	t := new(Tag)
	if err := fn(t); err != nil {
		return nil, err
	}

	if t.ID.IsZero() || t.UserID.IsZero() || t.Name.IsZero() {
		return nil, fmt.Errorf("failed to construct %+v", t)
	}

	return t, nil
}

func RearrangeTags(tags []*Tag, ids []idtype.TagID) {
	idToNewOrder := make(map[idtype.TagID]shareddomain.NonNegativeInt, len(ids))
	for i, id := range ids {
		idToNewOrder[id] = shareddomain.NonNegativeInt(i)
	}

	for _, tag := range tags {
		tag.Order = idToNewOrder[tag.ID]
	}
}
