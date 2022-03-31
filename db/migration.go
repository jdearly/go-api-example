package db

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
)

func MigrateDB(db *sql.DB) error {
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		return err
	}

	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	// TODO: allow docker to find migration files
	migration, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file:///%s/db/migrations", dir),
		"mysql",
		driver,
	)
	if err != nil {
		return err
	}

	err = migration.Up()

	if err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}
