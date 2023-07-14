package postgres

import (
	"KidneySmartBackend/internal/config"
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func GetDB(cfg config.Config) (*sql.DB, error) {

	dbUser := cfg.Database.User
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := cfg.Database.Name
	sslMode := cfg.Database.SslMode

	connStr := fmt.Sprintf(
		"user=%s password=%s dbname=%s sslmode=%s",
		dbUser, dbPassword, dbName, sslMode)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}

	// Проверка подключения
	if err := db.Ping(); err != nil {
		db.Close() // Закрываем соединение в случае ошибки
		return nil, fmt.Errorf("error connecting to database: %w", err)
	}
	return db, nil
}
