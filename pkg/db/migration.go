package db

import (
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log"
)

func migrateSQLFiles(db *sql.DB, folder string) error {
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		return err
	}
	m, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file:///%s", folder),
		"mysql",
		driver,
	)

	if err != nil {
		log.Println(24)
		return err
	}

	m.Steps(2)
	return nil
}
