package handlers

import (
	"github.com/THUNDERGROOVE/SDETool2/log"
	"github.com/THUNDERGROOVE/go-fit/db"
	"net/http"
)

type Global struct {
	User        db.User
	Fit         db.Fit
	Errors      []error
	Method      string
	UhOh        error
	AdminManage AdminManage
	Fits        []db.Fit
}

func GetGlobal(rw http.ResponseWriter, req *http.Request) Global {
	_, user, _ := GetSession(rw, req)
	var g Global
	g.Errors = make([]error, 0)
	g.User = user
	g.Method = req.Method
	g.AdminManage = AdminManage{}
	g.AdminManage.UserSet = make([]db.User, 0)
	g.Fits = make([]db.Fit, 0)
	return g
}

func Index(rw http.ResponseWriter, req *http.Request) {
	var g Global
	_, user, _ := GetSession(rw, req)
	g.User = user
	if t, ok := Templates["index.tmpl"]; ok {
		if err := t.Execute(rw, g); err != nil {
			log.LogError(err.Error())
		}
	} else {
		log.LogError("Template for index was not found.")
	}
}
