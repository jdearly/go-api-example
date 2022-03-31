package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func CreateDB() (*sql.DB, error) {
	serverName := "db_mysql:3306"
	user := "admin"
	password := "password"
	dbName := "example"
	connectionString :=
		fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&collation=utf8_unicode_ci&parseTime=true&multiStatements=true",
			user, password, serverName, dbName)
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		return nil, err
	}
	// TODO: this should work..
	//if err := MigrateDB(db); err != nil {
	//	return db, err
	//}

	return db, nil
}
