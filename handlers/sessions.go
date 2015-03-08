package handlers

import (
	"github.com/THUNDERGROOVE/SDETool2/log"
	"github.com/THUNDERGROOVE/go-fit/config"
	"github.com/THUNDERGROOVE/go-fit/db"
	"github.com/gorilla/sessions"
	"net/http"
)

var Sessions = sessions.NewCookieStore([]byte(config.Config.SessionSecret))

func GetSession(res http.ResponseWriter, r *http.Request) (bool, db.User, *sessions.Session) {
	user := db.User{}
	var loggedin bool
	var ses *sessions.Session
	if s, err := Sessions.Get(r, "go-fit"); err == nil {
		ses = s
		if user, ok := s.Values["user"].(db.User); ok {
			if newuser, ok := db.VerifyCredientials(user.Username, user.PassHash); ok {
				user = newuser
				return true, newuser, s
			} else {
				log.LogError("Someone logged in with malformed user data?", user.Username)
			}
		} else {
			loggedin = false
		}

	} else {
		log.LogError(err.Error())
	}
	return loggedin, user, ses
}
