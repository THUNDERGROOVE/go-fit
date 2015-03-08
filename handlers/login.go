package handlers

import (
	"fmt"
	"github.com/THUNDERGROOVE/SDETool2/log"
	"github.com/THUNDERGROOVE/go-fit/db"
	"net/http"
)

func Login(rw http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		if err := req.ParseMultipartForm(2056); err != nil {
			log.LogError(err.Error())
			if err := req.ParseForm(); err != nil {
				log.LogError(err.Error())
			}
		}
		username := req.Form.Get("username")
		password := req.Form.Get("password")
		log.Info("LoginPOST", username)

		if user, ok := db.VerifyCredientialsHash(username, password); ok {
			_, _, s := GetSession(rw, req)
			s.Values["user"] = interface{}(user)
			if err := s.Save(req, rw); err != nil {
				log.LogError(err.Error())
			} else {
				log.Info("Sucessfully logged user", user.Username, "in")
			}
		} else {
			fmt.Fprintf(rw, "Error logging in %v\n", "Username or password incorrect.")
		}
		http.Redirect(rw, req, "/", 301)
		return
	}

	if t, ok := Templates["login.tmpl"]; ok {
		if err := t.ExecuteTemplate(rw, "login.tmpl", nil); err != nil {
			log.LogError(err.Error())
		}
	} else {
		log.LogError("Template for login was not found.")
	}
}

func Logout(rw http.ResponseWriter, req *http.Request) {
	_, u, ses := GetSession(rw, req)
	fmt.Println("User", u.Username, "requesting logout.")
	if ses != nil {
		delete(ses.Values, "user")
		if err := ses.Save(req, rw); err != nil {
			log.LogError(err.Error())
		}
	}
	http.Redirect(rw, req, "/", 301)
}
