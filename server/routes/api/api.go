package api

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/pbtrung/Skypiea/server/app"
)

var db = app.InitDB(os.Getenv("DB_PATH"))

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	w.Write([]byte("Get a book " + id))
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
