package db

import (
	"fmt"
	"mini-crm-billing-api/source/config"
	"mini-crm-billing-api/source/pkg/logger"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Database(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		cfg.DBHost,
		cfg.DBUser,
		cfg.DBPass,
		cfg.DBName,
		cfg.DBPort,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Error().Msg("Failed connect db")
		return nil, err
	}

	return db, nil
}
