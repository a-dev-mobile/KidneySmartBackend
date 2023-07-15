package migrations

import (
	envConst "KidneySmartBackend/internal/env"
	"database/sql"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func PerformMigration(env string, dbName string, db *sql.DB) error {

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

	// Проверка текущей среды
	if env == envConst.Local || env == envConst.Dev {
		// Если мы в локальной среде, сперва откатываем миграции
		err = m.Down()
		if err != nil && err != migrate.ErrNoChange {
			return err
		}

	}
	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}
