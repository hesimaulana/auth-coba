package usersdb

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func connectDB() error {
	db, err := sql.Open("sqlite3", "./backend/datasource/sqlite/users_db/users")

	if err != nil {
		return err
	}

	DB = db
	return nil
}
