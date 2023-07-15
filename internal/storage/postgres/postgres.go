package postgres

import (
	"KidneySmartBackend/internal/config"
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/lib/pq"
)

func GetDB(cfg config.Config) (*sql.DB, error) {
	connStr := buildConnString(cfg)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}

	err = configureDBPool(db)
	if err != nil {
		return nil, err
	}

	// Проверка подключения
	if _, err := db.Exec("SELECT 1"); err != nil {
		closeErr := db.Close()
		if closeErr != nil {
			return nil, fmt.Errorf("error connecting to database: %w, error closing database: %v", err, closeErr)
		}
		return nil, fmt.Errorf("error connecting to database: %w", err)
	}

	return db, nil
}

func buildConnString(cfg config.Config) string {
	dbUser := cfg.Database.User
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := cfg.Database.Name
	sslMode := cfg.Database.SslMode
	dbHost := cfg.Database.Host
	dbPort := cfg.Database.Port

	return fmt.Sprintf(
		"user=%s password=%s dbname=%s sslmode=%s host=%s port=%d connect_timeout=10",
		dbUser, dbPassword, dbName, sslMode, dbHost, dbPort)
}

func configureDBPool(db *sql.DB) error {
	// Настройка пула соединений
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return nil
}
