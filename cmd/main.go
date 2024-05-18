package main

import (
	"Tatarinhack_backend/internal/delivery"
	"Tatarinhack_backend/pkg/config"
	"Tatarinhack_backend/pkg/database"
	"Tatarinhack_backend/pkg/logger"
)

func main() {
	log, loggerInfoFile, loggerErrorFile := logger.InitLogger()

	defer loggerInfoFile.Close()
	defer loggerErrorFile.Close()

	config.InitConfig()
	log.Info("Config initialized")

	db := database.GetDB()
	log.Info("Database initialized")

	delivery.Start(db, log)

}
