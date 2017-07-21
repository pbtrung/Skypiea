package app

import (
	"encoding/gob"

	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
)

var (
	Store *sessions.CookieStore
)

func Init() error {
	Store = sessions.NewCookieStore(securecookie.GenerateRandomKey(32), securecookie.GenerateRandomKey(32))
	gob.Register(map[string]interface{}{})
	return nil
}
