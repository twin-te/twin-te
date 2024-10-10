package dbhelper

import (
	"database/sql"
	"errors"
	"time"

	"github.com/samber/lo"
	"github.com/samber/mo"
	"github.com/twin-te/twin-te/back/appenv"
	sharedport "github.com/twin-te/twin-te/back/module/shared/port"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
)

func NewDB() (*gorm.DB, error) {
	return gorm.Open(postgres.New(postgres.Config{
		DSN:                  appenv.DB_URL,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{
		SkipDefaultTransaction: true,
		Logger:                 logger.Default.LogMode((logger.Info)),
		NowFunc:                func() time.Time { return time.Now().Truncate(time.Microsecond) },
	})
}

func ConvertErrRecordNotFound[T any](err error) (mo.Option[T], error) {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return mo.None[T](), nil
	}
	return mo.None[T](), err
}

func ApplyLock(db *gorm.DB, lock sharedport.Lock) *gorm.DB {
	if lock != sharedport.LockNone {
		db = db.Clauses(clause.Locking{
			Strength: lo.Ternary(lock == sharedport.LockExclusive, "UPDATE", "SHARE"),
			Table:    clause.Table{Name: clause.CurrentTable},
		})
	}
	return db
}

func ApplyLimitOffset(db *gorm.DB, limitOffset sharedport.LimitOffset) *gorm.DB {
	if limitOffset.Limit > 0 {
		db = db.Limit(limitOffset.Limit)
	}

	if limitOffset.Offset > 0 {
		db = db.Offset(limitOffset.Offset)
	}

	return db
}

func OptionToNull[T any](o mo.Option[T]) sql.Null[T] {
	v, valid := o.Get()
	return sql.Null[T]{
		V:     v,
		Valid: valid,
	}
}

func NullToOption[T any](null sql.Null[T]) mo.Option[T] {
	if null.Valid {
		return mo.Some(null.V)
	}
	return mo.None[T]()
}
