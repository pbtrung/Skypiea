package app

import (
	"database/sql"
	"fmt"

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

	rows, err := db.Query("SELECT * FROM books WHERE id < 10")
	fmt.Println(rows)

	return db
}
