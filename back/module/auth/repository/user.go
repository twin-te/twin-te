package authrepository

import (
	"context"
	"time"

	"github.com/samber/lo"
	"github.com/samber/mo"
	"github.com/twin-te/twin-te/back/base"
	dbhelper "github.com/twin-te/twin-te/back/db/helper"
	authdbmodel "github.com/twin-te/twin-te/back/module/auth/dbmodel"
	authdomain "github.com/twin-te/twin-te/back/module/auth/domain"
	authport "github.com/twin-te/twin-te/back/module/auth/port"
	sharedport "github.com/twin-te/twin-te/back/module/shared/port"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (r *impl) FindUser(ctx context.Context, conds authport.FindUserConds, lock sharedport.Lock) (mo.Option[*authdomain.User], error) {
	if err := conds.Validate(); err != nil {
		return mo.None[*authdomain.User](), err
	}

	dbUser := new(authdbmodel.User)
	err := r.transaction(ctx, func(tx *gorm.DB) error {
		if id, ok := conds.ID.Get(); ok {
			tx = tx.Where("id = ?", id.String())
		}

		if userAuthentication, ok := conds.UserAuthentication.Get(); ok {
			tx = tx.Where(
				"id = ( ? )",
				tx.Select("user_id").Where("provider = ? AND social_id = ?",
					userAuthentication.Provider.String(),
					userAuthentication.SocialID.String(),
				).Table("user_authentications"),
			)
		}

		tx = dbhelper.ApplyLock(tx, lock)

		return tx.
			Where("deleted_at IS NULL").
			Preload("UserAuthentications").
			Take(dbUser).
			Error
	}, true)
	if err != nil {
		return dbhelper.ConvertErrRecordNotFound[*authdomain.User](err)
	}

	return base.SomeWithErr(authdbmodel.FromDBUser(dbUser))
}

func (r *impl) ListUsers(ctx context.Context, conds authport.ListUsersConds, lock sharedport.Lock) ([]*authdomain.User, error) {
	var dbUsers []*authdbmodel.User

	err := r.transaction(ctx, func(tx *gorm.DB) error {
		tx = dbhelper.ApplyLock(tx, lock)
		return tx.
			Where("deleted_at IS NULL").
			Preload("UserAuthentications").
			Find(&dbUsers).
			Error
	}, true)
	if err != nil {
		return nil, err
	}

	return base.MapWithErr(dbUsers, authdbmodel.FromDBUser)
}

func (r *impl) CreateUsers(ctx context.Context, users ...*authdomain.User) error {
	dbUsers := base.MapWithArg(users, true, authdbmodel.ToDBUser)
	dbUserAuthentications := lo.Flatten(base.Map(dbUsers, func(dbUser *authdbmodel.User) []authdbmodel.UserAuthentication {
		return dbUser.UserAuthentications
	}))
	return r.transaction(ctx, func(tx *gorm.DB) error {
		if err := tx.Omit(clause.Associations).Create(dbUsers).Error; err != nil {
			return err
		}
		if len(dbUserAuthentications) > 0 {
			if err := tx.Create(dbUserAuthentications).Error; err != nil {
				return err
			}
		}
		return nil
	}, false)
}

func (r *impl) UpdateUser(ctx context.Context, user *authdomain.User) error {
	before := user.BeforeUpdated.MustGet()
	columns := make([]string, 0)

	if !user.CreatedAt.Equal(before.CreatedAt) {
		columns = append(columns, "created_at")
	}

	dbUser := authdbmodel.ToDBUser(user, false)

	return r.transaction(ctx, func(tx *gorm.DB) error {
		if len(columns) > 0 {
			if err := tx.Select(columns).Updates(dbUser).Error; err != nil {
				return err
			}
		}
		return r.updateUserAuthentications(tx, user)
	}, false)
}

func (r *impl) DeleteUsers(ctx context.Context, conds authport.DeleteUserConds) (rowsAffected int, err error) {
	err = r.transaction(ctx, func(tx *gorm.DB) error {
		var dbUsers []*authdbmodel.User
		tx = tx.Model(&dbUsers)

		if id, ok := conds.ID.Get(); ok {
			tx.Where("id = ?", id.String())
		}

		if err := tx.Clauses(clause.Returning{Columns: []clause.Column{{Name: "id"}}}).
			Update("deleted_at", time.Now()).
			Error; err != nil {
			return err
		}

		if rowsAffected = int(tx.RowsAffected); rowsAffected == 0 {
			return nil
		}

		return r.db.
			Where("user_id IN ?", base.Map(dbUsers, func(dbUser *authdbmodel.User) string {
				return dbUser.ID
			})).
			Delete(&authdbmodel.UserAuthentication{}).
			Error
	}, false)
	return
}
