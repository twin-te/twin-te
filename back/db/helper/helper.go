package dbhelper

import (
	"errors"

	"github.com/twin-te/twinte-back/appenv"
	sharedport "github.com/twin-te/twinte-back/module/shared/port"
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
