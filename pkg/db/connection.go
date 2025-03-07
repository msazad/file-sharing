package db

import (
	"file-sharing/pkg/config"
	"file-sharing/pkg/domain"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB(cfg config.Config) (*gorm.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s", cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBPort, cfg.DBName)

	db, dbErr := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{
		SkipDefaultTransaction: true,
	})

	if err := db.AutoMigrate(domain.User{}); err != nil {
		return db, err
	}

	if err := db.AutoMigrate(domain.File{}); err != nil {
		return db, err
	}

	return db, dbErr
}
