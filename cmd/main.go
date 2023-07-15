package main

import (
	"KidneySmartBackend/internal/config"
	"KidneySmartBackend/internal/lib/logger/sl"
	"KidneySmartBackend/internal/migrations"
	"KidneySmartBackend/internal/storage/postgres"

	"log"

	"fmt"
	"os"

	"golang.org/x/exp/slog"
)

func main() {
	appEnv, err := getAppEnv()
	if err != nil {
		log.Fatalf("Error getting app environment: %s", err)
	}

	cfg, err := config.GetConfig(appEnv)
	if err != nil {
		log.Fatalf("Error loading config: %s", err)
	}
	lg := sl.SetupLogger(appEnv, cfg.Logging.Level)

	db, err := postgres.GetDB(cfg)
	if err != nil {
		log.Fatalf("Error initializing database: %s", err)
	}
	_ = db

	// Выполнение миграций
	err = migrations.PerformMigration(appEnv, cfg.Database.Name, db)
	if err != nil {
		log.Fatalf("Error performing migrations: %s", err)
	}

	lg.Info("start KidneySmartBackend", slog.String("env", appEnv))
	lg.Info("Loaded config file", slog.Any("config_json", cfg))
	lg.Debug("debug msg are enabled")
}

func getAppEnv() (string, error) {
	appEnv := os.Getenv("APP_ENV")
	if appEnv == "" {
		return "", fmt.Errorf("APP_ENV is not set")
	}

	return appEnv, nil
}
