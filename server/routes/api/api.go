package api

import (
	"net/http"
)

func GetBookById(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	w.Write([]byte("Get a book!"))
}
