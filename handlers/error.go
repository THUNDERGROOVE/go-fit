package handlers

import (
	"github.com/THUNDERGROOVE/SDETool2/log"
	"net/http"
)

func WriteError(err error, rw http.ResponseWriter, req *http.Request) {
	if err == nil {
		return
	}
	g := GetGlobal(rw, req)
	g.UhOh = err
	if t, ok := Templates["error.tmpl"]; ok {
		if err := t.ExecuteTemplate(rw, "error.tmpl", g); err != nil {
			log.LogError(err.Error())
		}
	} else {
		log.LogError("Template for error was not found.")
	}
}
