package dao

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func getDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./db.sqlite3")
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(
		`create table if not exists todo (id INTEGER PRIMARY KEY AUTOINCREMENT, member_code VARCHAR(255), task_name VARCHAR(255), completed INTEGER)`,
	)
	if err != nil {
		return nil, err
	}
	return db, nil
}
