package migrations

import (
	"database/sql"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func PerformMigration(dbName string, db *sql.DB) error {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return err
	}

	// Путь к файлу миграции
	migrationsDir := "file://../internal/migrations/"

	m, err := migrate.NewWithDatabaseInstance(
		migrationsDir, // Источник файлов миграции
		dbName,        // Имя базы данных
		driver,        // Driver конфигурации базы данных
	)

	if err != nil {
		return err
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}
