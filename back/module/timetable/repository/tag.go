package timetablerepository

import (
	"context"

	"github.com/samber/mo"
	"github.com/twin-te/twin-te/back/base"
	dbhelper "github.com/twin-te/twin-te/back/db/helper"
	sharedport "github.com/twin-te/twin-te/back/module/shared/port"
	timetabledbmodel "github.com/twin-te/twin-te/back/module/timetable/dbmodel"
	timetabledomain "github.com/twin-te/twin-te/back/module/timetable/domain"
	timetableport "github.com/twin-te/twin-te/back/module/timetable/port"
	"gorm.io/gorm"
)

func (r *impl) FindTag(ctx context.Context, filter timetableport.TagFilter, lock sharedport.Lock) (mo.Option[*timetabledomain.Tag], error) {
	db := r.db.WithContext(ctx)
	db = applyTagFilter(db, filter)
	db = dbhelper.ApplyLock(db, lock)

	dbTag := new(timetabledbmodel.Tag)
	if err := db.Take(&dbTag).Error; err != nil {
		return dbhelper.ConvertErrRecordNotFound[*timetabledomain.Tag](err)
	}

	return base.SomeWithErr(timetabledbmodel.FromDBTag(dbTag))
}

func (r *impl) ListTags(ctx context.Context, filter timetableport.TagFilter, limitOffset sharedport.LimitOffset, lock sharedport.Lock) ([]*timetabledomain.Tag, error) {
	db := r.db.WithContext(ctx)
	db = applyTagFilter(db, filter)
	db = dbhelper.ApplyLimitOffset(db, limitOffset)
	db = dbhelper.ApplyLock(db, lock)

	var dbTags []*timetabledbmodel.Tag
	if err := db.Find(&dbTags).Error; err != nil {
		return nil, err
	}

	return base.MapWithErr(dbTags, timetabledbmodel.FromDBTag)
}

func (r *impl) CreateTags(ctx context.Context, tags ...*timetabledomain.Tag) error {
	dbTags := base.Map(tags, timetabledbmodel.ToDBTag)
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

	dbTag := timetabledbmodel.ToDBTag(tag)
	return r.db.WithContext(ctx).
		Select(columns).
		Updates(dbTag).
		Error
}

func (r *impl) DeleteTags(ctx context.Context, filter timetableport.TagFilter) (rowsAffected int, err error) {
	db := r.db.WithContext(ctx)
	db = applyTagFilter(db, filter)
	return int(db.Delete(&timetabledbmodel.Tag{}).RowsAffected), db.Error
}

func applyTagFilter(db *gorm.DB, filter timetableport.TagFilter) *gorm.DB {
	if id, ok := filter.ID.Get(); ok {
		db = db.Where("id = ?", id.String())
	}

	if ids, ok := filter.IDs.Get(); ok {
		db = db.Where("id IN ?", base.MapByString(ids))
	}

	if userID, ok := filter.UserID.Get(); ok {
		db = db.Where("user_id = ?", userID.String())
	}

	return db
}
