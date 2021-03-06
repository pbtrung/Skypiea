package api

import (
	"fmt"
	"net/http"
	"os"

	"github.com/pbtrung/Skypiea/server/app"
)

var db = app.InitDB(os.Getenv("DB_PATH"))

func GetBookById(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT * FROM books WHERE id < 10")
	checkErr(err)
	fmt.Println(rows)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
