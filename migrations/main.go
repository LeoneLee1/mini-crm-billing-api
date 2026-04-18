package main

import (
	"mini-crm-billing-api/source/common/models"
	"mini-crm-billing-api/source/config"
	"mini-crm-billing-api/source/pkg/db"
	"mini-crm-billing-api/source/pkg/logger"
)

func main() {
	logger.Init(true)

	cfg := config.Load()

	dbConn, err := db.Database(cfg)
	if err != nil {
		logger.Log.Fatal().Msg("Database connection failed")
	}
	logger.Log.Info().Msg("Running database migration...")

	if err := dbConn.AutoMigrate(
		&models.UserModel{},
		&models.RefreshTokenModel{},
		&models.LogActivityModel{},
		&models.CustomerModel{},
		&models.ProductModel{},
		&models.TransactionModel{},
		&models.TransactionItemModel{},
		&models.InvoiceModel{},
		&models.PaymentModel{},
	); err != nil {
		logger.Log.Fatal().Msg("Migration failed")
	}

	logger.Log.Info().Msg("Migration success")
}
