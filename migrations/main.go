package main

import (
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

	if err := dbConn.AutoMigrate(); err != nil {
		logger.Log.Fatal().Msg("Migration failed")
	}

	dbConn.Exec(`
		ALTER TABLE absences
		MODIFY clock_in TIME,
		MODIFY clock_out TIME NULL;
	`)

	dbConn.Exec(`
		ALTER TABLE booking_rooms
		MODIFY start TIME,
		MODIFY end TIME;
	`)

	logger.Log.Info().Msg("Migration success")
}
