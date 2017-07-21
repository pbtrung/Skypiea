package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/pbtrung/Skypiea/server/routes/api"
	"github.com/pbtrung/Skypiea/server/routes/callback"
	"github.com/pbtrung/Skypiea/server/routes/home"
	"github.com/pbtrung/Skypiea/server/routes/middlewares"
	"github.com/urfave/negroni"
)

func StartServer() {
	r := mux.NewRouter().StrictSlash(false)

	r.HandleFunc("/", home.HomeHandler)
	r.HandleFunc("/callback", callback.CallbackHandler)

	r.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("public/"))))

	apiRouter := r.PathPrefix("/api").Subrouter()
	apiRouter.HandleFunc("/book/{id}", api.GetBookById)

	muxGroup := http.NewServeMux()
	muxGroup.Handle("/", r)
	muxGroup.Handle("/api/", negroni.New(
		negroni.HandlerFunc(middlewares.IsAuthenticated),
		negroni.Wrap(r),
	))

	var port string
	if os.Getenv("PORT") == "" {
		port = "3000"
	} else {
		port = os.Getenv("PORT")
	}
	n := negroni.Classic()
	n.UseHandler(muxGroup)
	log.Fatalln(http.ListenAndServe(":"+port, n))
}
