package repository

import (
	"mose/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func CreateConnectionPostreSql(config config.Config) (*gorm.DB, error) {

	return gorm.Open(postgres.New(postgres.Config{
		DSN:                  config.GetDBURL(),
		PreferSimpleProtocol: true,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
}
