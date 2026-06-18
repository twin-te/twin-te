package timetablequery

import (
	"context"

	"github.com/twin-te/twin-te/back/base"
	"github.com/twin-te/twin-te/back/module/shared/domain/idtype"
	timetabledbmodel "github.com/twin-te/twin-te/back/module/timetable/dbmodel"
	"gorm.io/gorm"
)

func (q *impl) ListTagIDsByUserID(ctx context.Context, userID idtype.UserID, ids []idtype.TagID) ([]idtype.TagID, error) {
	if len(ids) == 0 {
		return nil, nil
	}

	var dbTags []*timetabledbmodel.Tag
	if err := q.gormTransaction(ctx, func(tx *gorm.DB) error {
		return tx.
			Where("user_id = ?", userID.String()).
			Where("id IN ?", base.MapByString(ids)).
			Find(&dbTags).
			Error
	}); err != nil {
		return nil, err
	}

	return base.MapWithErr(dbTags, func(dbTag *timetabledbmodel.Tag) (idtype.TagID, error) {
		return idtype.ParseTagID(dbTag.ID)
	})
}
