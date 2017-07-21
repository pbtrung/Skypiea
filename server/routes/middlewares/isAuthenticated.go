package middlewares

import (
	"net/http"

	"github.com/pbtrung/Skypiea/server/app"
)

func IsAuthenticated(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	session, err := app.Store.Get(r, "auth0-session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if _, ok := session.Values["is_authenticated"]; !ok {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else {
		next(w, r)
	}
}
