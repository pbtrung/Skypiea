package api

import (
	"net/http"
	"os"

	"github.com/pbtrung/Skypiea/server/app"
)

var db = app.InitDB(os.Getenv("DB_PATH"))

func GetBookById(w http.ResponseWriter, r *http.Request) {
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
