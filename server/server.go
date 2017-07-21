package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/pbtrung/Skypiea/server/routes/callback"
	"github.com/pbtrung/Skypiea/server/routes/home"
	"github.com/urfave/negroni"
)

func StartServer() {
	r := mux.NewRouter()

	r.HandleFunc("/", home.HomeHandler)
	r.HandleFunc("/callback", callback.CallbackHandler)
	r.Handle("/user", negroni.New(
		negroni.HandlerFunc(middlewares.IsAuthenticated),
		negroni.Wrap(http.HandlerFunc(user.UserHandler)),
	))

	r.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("public/"))))
	var port string
	if os.Getenv("PORT") == "" {
		port = "3000"
	} else {
		port = os.Getenv("PORT")
	}
	log.Fatalln(http.ListenAndServe(":"+port, r))
}
