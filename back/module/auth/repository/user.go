package authrepository

import (
	"context"
	"fmt"

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

func (r *impl) FindUser(ctx context.Context, filter authport.UserFilter, lock sharedport.Lock) (mo.Option[*authdomain.User], error) {
	if !filter.IsUniqueFilter() {
		return mo.None[*authdomain.User](), fmt.Errorf("%v is not unique", filter)
	}

	dbUser := new(authdbmodel.User)
	err := r.transaction(ctx, func(tx *gorm.DB) error {
		tx = applyUserFilter(tx, filter)
		tx = dbhelper.ApplyLock(tx, lock)
		return tx.
			Preload("UserAuthentications").
			Take(dbUser).
			Error
	}, true)
	if err != nil {
		return dbhelper.ConvertErrRecordNotFound[*authdomain.User](err)
	}

	return base.SomeWithErr(authdbmodel.FromDBUser(dbUser))
}

func (r *impl) ListUsers(ctx context.Context, filter authport.UserFilter, limitOffset sharedport.LimitOffset, lock sharedport.Lock) ([]*authdomain.User, error) {
	var dbUsers []*authdbmodel.User
	err := r.transaction(ctx, func(tx *gorm.DB) error {
		tx = applyUserFilter(tx, filter)
		tx = dbhelper.ApplyLimitOffset(tx, limitOffset)
		tx = dbhelper.ApplyLock(tx, lock)
		return tx.
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
	dbUser := authdbmodel.ToDBUser(user, false)
	return r.transaction(ctx, func(tx *gorm.DB) error {
		if err := tx.Updates(dbUser).Error; err != nil {
			return err
		}
		return r.updateUserAuthentications(tx, user)
	}, false)
}

func (r *impl) DeleteUsers(ctx context.Context, filter authport.UserFilter) (rowsAffected int, err error) {
	err = r.transaction(ctx, func(tx *gorm.DB) error {
		tx = applyUserFilter(tx, filter)
		rowsAffected = int(tx.Delete(&authdbmodel.User{}).RowsAffected)
		return tx.Error
	}, false)
	return
}

func applyUserFilter(db *gorm.DB, filter authport.UserFilter) *gorm.DB {
	if id, ok := filter.ID.Get(); ok {
		db.Where("id = ?", id.String())
	}

	if userAuthentication, ok := filter.UserAuthentication.Get(); ok {
		db = db.Where(
			"id = ( ? )",
			db.Select("user_id").Where("provider = ? AND social_id = ?",
				userAuthentication.Provider.String(),
				userAuthentication.SocialID.String(),
			).Table("user_authentications"),
		)
	}

	return db
}
