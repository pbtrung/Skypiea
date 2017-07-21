package api

import (
	"net/http"
)

func GetBookById(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get a book!"))
}
