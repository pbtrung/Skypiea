package api

import (
	"fmt"
	"net/http"
	"os"

	"github.com/pbtrung/Skypiea/server/app"
)

func GetBookById(w http.ResponseWriter, r *http.Request) {
	var db = app.InitDB(os.Getenv("DB_PATH"))

	rows, err := db.Query("SELECT * FROM books WHERE id < 10")
	checkErr(err)
	fmt.Println(rows)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
