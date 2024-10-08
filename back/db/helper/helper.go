package dbhelper

import (
	"database/sql"
	"errors"

	"github.com/samber/mo"
	"github.com/twin-te/twin-te/back/appenv"
	sharedport "github.com/twin-te/twin-te/back/module/shared/port"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// ConvertErrRecordNotFound converts gorm.ErrRecordNotFound into sharedport.ErrNotFound,
func ConvertErrRecordNotFound(err error) error {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return sharedport.ErrNotFound
	}
	return err
}

func NewDB() (*gorm.DB, error) {
	return gorm.Open(postgres.New(postgres.Config{
		DSN:                  appenv.DB_URL,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{
		SkipDefaultTransaction:   true,
		Logger:                   logger.Default.LogMode((logger.Info)),
		DisableNestedTransaction: true,
	})
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
