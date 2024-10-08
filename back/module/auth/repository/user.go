package authrepository

import (
	"context"
	"errors"
	"time"

	"github.com/samber/lo"
	"github.com/samber/mo"
	"github.com/twin-te/twin-te/back/base"
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
	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
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

		return tx.
			Where("deleted_at IS NULL").
			Clauses(clause.Locking{
				Strength: lo.Ternary(lock == sharedport.LockExclusive, "UPDATE", "SHARE"),
				Table:    clause.Table{Name: clause.CurrentTable},
			}).
			Preload("UserAuthentications").
			Take(dbUser).
			Error
	}, nil)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return mo.None[*authdomain.User](), nil
		}
		return mo.None[*authdomain.User](), err
	}

	return base.SomeWithErr(authdbmodel.FromDBUser(dbUser))
}

func (r *impl) ListUsers(ctx context.Context, conds authport.ListUsersConds, lock sharedport.Lock) ([]*authdomain.User, error) {
	var dbUsers []*authdbmodel.User

	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return tx.
			Where("deleted_at IS NULL").
			Clauses(clause.Locking{
				Strength: lo.Ternary(lock == sharedport.LockExclusive, "UPDATE", "SHARE"),
				Table:    clause.Table{Name: clause.CurrentTable},
			}).
			Preload("UserAuthentications").
			Find(&dbUsers).
			Error
	})
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
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Omit(clause.Associations).Create(dbUsers).Error; err != nil {
			return err
		}
		if len(dbUserAuthentications) > 0 {
			if err := tx.Create(dbUserAuthentications).Error; err != nil {
				return err
			}
		}
		return nil
	}, nil)
}

func (r *impl) UpdateUser(ctx context.Context, user *authdomain.User) error {
	before := user.BeforeUpdated.MustGet()
	columns := make([]string, 0)

	if !user.CreatedAt.Equal(before.CreatedAt) {
		columns = append(columns, "created_at")
	}

	dbUser := authdbmodel.ToDBUser(user, false)

	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if len(columns) > 0 {
			if err := tx.Select(columns).Updates(dbUser).Error; err != nil {
				return err
			}
		}
		return r.updateUserAuthentications(tx, user)
	}, nil)
}

func (r *impl) DeleteUsers(ctx context.Context, conds authport.DeleteUserConds) (rowsAffected int, err error) {
	err = r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
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
	}, nil)
	return
}
