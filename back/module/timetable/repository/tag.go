package timetablerepository

import (
	"context"

	"github.com/samber/lo"
	"github.com/twin-te/twin-te/back/base"
	"github.com/twin-te/twin-te/back/db/gen/model"
	dbhelper "github.com/twin-te/twin-te/back/db/helper"
	"github.com/twin-te/twin-te/back/module/shared/domain/idtype"
	sharedport "github.com/twin-te/twin-te/back/module/shared/port"
	timetabledomain "github.com/twin-te/twin-te/back/module/timetable/domain"
	timetableport "github.com/twin-te/twin-te/back/module/timetable/port"
	"gorm.io/gorm/clause"
)

func (r *impl) FindTag(ctx context.Context, conds timetableport.FindTagConds, lock sharedport.Lock) (*timetabledomain.Tag, error) {
	db := r.db.WithContext(ctx).
		Where("id = ?", conds.ID.String())

	if conds.UserID != nil {
		db = db.Where("user_id = ?", conds.UserID.String())
	}

	if lock != sharedport.LockNone {
		db = db.Clauses(clause.Locking{
			Strength: lo.Ternary(lock == sharedport.LockExclusive, "UPDATE", "SHARE"),
			Table:    clause.Table{Name: clause.CurrentTable},
		})
	}

	dbTag := new(model.Tag)
	if err := db.Take(&dbTag).Error; err != nil {
		return nil, dbhelper.ConvertErrRecordNotFound(err)
	}

	return fromDBTag(dbTag)
}

func (r *impl) ListTags(ctx context.Context, conds timetableport.ListTagsConds, lock sharedport.Lock) ([]*timetabledomain.Tag, error) {
	db := r.db.WithContext(ctx)

	if conds.IDs != nil {
		db = db.Where("id IN ?", base.MapByString(*conds.IDs))
	}

	if conds.UserID != nil {
		db = db.Where("user_id = ?", conds.UserID.String())
	}

	if lock != sharedport.LockNone {
		db = db.Clauses(clause.Locking{
			Strength: lo.Ternary(lock == sharedport.LockExclusive, "UPDATE", "SHARE"),
			Table:    clause.Table{Name: clause.CurrentTable},
		})
	}

	var dbTags []*model.Tag
	if err := db.Find(&dbTags).Error; err != nil {
		return nil, err
	}

	return base.MapWithErr(dbTags, fromDBTag)
}

func (r *impl) CreateTags(ctx context.Context, tags ...*timetabledomain.Tag) error {
	dbTags := base.Map(tags, toDBTag)
	return r.db.WithContext(ctx).Create(dbTags).Error
}

func (r *impl) UpdateTag(ctx context.Context, tag *timetabledomain.Tag) error {
	before := tag.EntityBeforeUpdated.MustGet()
	columns := make([]string, 0)

	if tag.UserID != before.UserID {
		columns = append(columns, "user_id")
	}

	if tag.Name != before.Name {
		columns = append(columns, "name")
	}

	if tag.Position != before.Position {
		columns = append(columns, "position")
	}

	if len(columns) == 0 {
		return nil
	}

	dbTag := toDBTag(tag)
	return r.db.WithContext(ctx).
		Select(columns).
		Updates(dbTag).
		Error
}

func (r *impl) DeleteTags(ctx context.Context, conds timetableport.DeleteTagsConds) (rowsAffected int, err error) {
	db := r.db.WithContext(ctx)

	if conds.ID != nil {
		db = db.Where("id = ?", conds.ID.String())
	}

	if conds.UserID != nil {
		db = db.Where("user_id = ?", conds.UserID.String())
	}

	return int(db.Delete(&model.Tag{}).RowsAffected), db.Error
}

func fromDBTag(dbTag *model.Tag) (*timetabledomain.Tag, error) {
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

		t.Position, err = timetabledomain.ParsePosition(int(dbTag.Position))
		if err != nil {
			return err
		}

		return nil
	})
}

func toDBTag(tag *timetabledomain.Tag) *model.Tag {
	return &model.Tag{
		ID:       tag.ID.String(),
		UserID:   tag.UserID.String(),
		Name:     tag.Name.String(),
		Position: int32(tag.Position.Int()),
	}
}
