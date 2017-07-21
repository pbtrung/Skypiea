package app

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB(filePath string) *sql.DB {
	db, err := sql.Open("sqlite3", filePath)
	if err != nil {
		panic(err)
	}
	if db == nil {
		panic("db nil")
	}
	return db
}
