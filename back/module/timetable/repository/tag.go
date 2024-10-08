package timetablerepository

import (
	"context"

	"github.com/samber/lo"
	"github.com/twin-te/twin-te/back/base"
	dbhelper "github.com/twin-te/twin-te/back/db/helper"
	"github.com/twin-te/twin-te/back/module/shared/domain/idtype"
	sharedport "github.com/twin-te/twin-te/back/module/shared/port"
	timetabledbmodel "github.com/twin-te/twin-te/back/module/timetable/dbmodel"
	timetabledomain "github.com/twin-te/twin-te/back/module/timetable/domain"
	timetableport "github.com/twin-te/twin-te/back/module/timetable/port"
	"gorm.io/gorm/clause"
)

func (r *impl) FindTag(ctx context.Context, conds timetableport.FindTagConds, lock sharedport.Lock) (*timetabledomain.Tag, error) {
	db := r.db.WithContext(ctx).
		Where("id = ?", conds.ID.String())

	if userID, ok := conds.UserID.Get(); ok {
		db = db.Where("user_id = ?", userID.String())
	}

	if lock != sharedport.LockNone {
		db = db.Clauses(clause.Locking{
			Strength: lo.Ternary(lock == sharedport.LockExclusive, "UPDATE", "SHARE"),
			Table:    clause.Table{Name: clause.CurrentTable},
		})
	}

	dbTag := new(timetabledbmodel.Tag)
	if err := db.Take(&dbTag).Error; err != nil {
		return nil, dbhelper.ConvertErrRecordNotFound(err)
	}

	return fromDBTag(dbTag)
}

func (r *impl) ListTags(ctx context.Context, conds timetableport.ListTagsConds, lock sharedport.Lock) ([]*timetabledomain.Tag, error) {
	db := r.db.WithContext(ctx)

	if ids, ok := conds.IDs.Get(); ok {
		db = db.Where("id IN ?", base.MapByString(ids))
	}

	if userID, ok := conds.UserID.Get(); ok {
		db = db.Where("user_id = ?", userID.String())
	}

	if lock != sharedport.LockNone {
		db = db.Clauses(clause.Locking{
			Strength: lo.Ternary(lock == sharedport.LockExclusive, "UPDATE", "SHARE"),
			Table:    clause.Table{Name: clause.CurrentTable},
		})
	}

	var dbTags []*timetabledbmodel.Tag
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
	before := tag.BeforeUpdated.MustGet()
	columns := make([]string, 0)

	if tag.UserID != before.UserID {
		columns = append(columns, "user_id")
	}

	if tag.Name != before.Name {
		columns = append(columns, "name")
	}

	if tag.Position != before.Position {
		columns = append(columns, "order")
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

	if id, ok := conds.ID.Get(); ok {
		db = db.Where("id = ?", id.String())
	}

	if userID, ok := conds.UserID.Get(); ok {
		db = db.Where("user_id = ?", userID.String())
	}

	return int(db.Delete(&timetabledbmodel.Tag{}).RowsAffected), db.Error
}

func fromDBTag(dbTag *timetabledbmodel.Tag) (*timetabledomain.Tag, error) {
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

		t.Position, err = timetabledomain.ParsePosition(int(dbTag.Order))
		if err != nil {
			return err
		}

		return nil
	})
}

func toDBTag(tag *timetabledomain.Tag) *timetabledbmodel.Tag {
	return &timetabledbmodel.Tag{
		ID:     tag.ID.String(),
		UserID: tag.UserID.String(),
		Name:   tag.Name.String(),
		Order:  tag.Position.Int(),
	}
}
