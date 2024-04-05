package timetabledomain

import (
	"fmt"

	"github.com/samber/lo"
	shareddomain "github.com/twin-te/twin-te/back/module/shared/domain"
	"github.com/twin-te/twin-te/back/module/shared/domain/idtype"
)

var ParsePosition = shareddomain.NewNonNegativeIntParser("position")

// Tag is identified by one of the following fields.
//   - ID
//   - UserID and Position
type Tag struct {
	ID       idtype.TagID
	UserID   idtype.UserID
	Name     shareddomain.RequiredString
	Position shareddomain.NonNegativeInt

	EntityBeforeUpdated *Tag
}

func (t *Tag) Clone() *Tag {
	ret := lo.ToPtr(*t)
	return ret
}

func (t *Tag) BeforeUpdateHook() {
	t.EntityBeforeUpdated = t.Clone()
}

type TagDataToUpdate struct {
	Name     *shareddomain.RequiredString
	Position *shareddomain.NonNegativeInt
}

func (t *Tag) Update(data TagDataToUpdate) {
	if data.Name != nil {
		t.Name = *data.Name
	}

	if data.Position != nil {
		t.Position = *data.Position
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
	idToNewPosition := make(map[idtype.TagID]shareddomain.NonNegativeInt, len(ids))
	for i, id := range ids {
		idToNewPosition[id] = shareddomain.NonNegativeInt(i)
	}

	for _, tag := range tags {
		tag.Update(TagDataToUpdate{
			Position: lo.ToPtr(idToNewPosition[tag.ID]),
		})
	}
}
