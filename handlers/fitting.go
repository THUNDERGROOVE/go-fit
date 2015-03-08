package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/THUNDERGROOVE/SDETool2/log"
	"github.com/THUNDERGROOVE/SDETool2/sde"
	"github.com/THUNDERGROOVE/go-fit/db"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
)

type Fit struct {
	ID       int
	Fit      db.Fit
	BaseType sde.SDEType
}

func Fitting(rw http.ResponseWriter, req *http.Request) {
	var g Global
	v := mux.Vars(req)
	g.Fit.FitID, _ = strconv.Atoi(v["FitID"])
	fmt.Println(g.Fit.FitID)
	if g.Fit.FitID == 0 {
		var err error
		g.Fit, err = db.GetFitFromDB(g.Fit.FitID)
		if err != nil {
			log.LogError(err.Error())
			g.Errors = append(g.Errors, err)
		}
	}
	_, user, _ := GetSession(rw, req)
	g.User = user
	if t, ok := Templates["fit.tmpl"]; ok {
		if err := t.Execute(rw, g); err != nil {
			log.LogError(err.Error())
		}
	} else {
		log.LogError("Template for fit was not found.")
	}
}

func NewFit(rw http.ResponseWriter, req *http.Request) {
	g := GetGlobal(rw, req)
	loggedIn, _, _ := GetSession(rw, req)
	if !loggedIn {
		WriteError(fmt.Errorf("You must be logged in"), rw, req)
		return
	}
	switch req.Method {
	case "GET":
	case "POST":
		if err := req.ParseMultipartForm(2056); err != nil {
			log.LogError(err.Error())
			if err := req.ParseForm(); err != nil {
				log.LogError(err.Error())
			}
		}
		idst := req.Form.Get("TypeID")
		id, _ := strconv.Atoi(idst)
		fid, err := db.NewFit(g.User.Username, id)
		if err != nil {
			log.LogError(err.Error())
			g.Errors = append(g.Errors, err)
		}
		http.Redirect(rw, req, fmt.Sprintf("/fit/%v", fid), 301)
	}
	if t, ok := Templates["newfit.tmpl"]; ok {
		if err := t.Execute(rw, g); err != nil {
			log.LogError(err.Error())
		}
	} else {
		log.LogError("Template for fit was not found.")
	}
}

func RawFit(rw http.ResponseWriter, req *http.Request) {
	v := mux.Vars(req)
	f := db.Fit{}
	f.FitID, _ = strconv.Atoi(v["FitID"])
	if f.FitID != 0 {
		var err error
		f, err = db.GetFitFromDB(f.FitID)
		if err != nil {
			log.LogError(err.Error())
		}
	}
	data, err := json.Marshal(f)
	if err != nil {
		log.LogError(err.Error())
	}
	rw.Write(data)
}

func SaveFit(rw http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	ID, _ := strconv.Atoi(vars["FitID"])
	data, _ := ioutil.ReadAll(req.Body)
	fit := db.Fit{}
	if err := json.Unmarshal(data, &fit); err != nil {
		rw.Write([]byte("Error unmarshaling JSON " + err.Error()))
		return
	}
	logged, user, _ := GetSession(rw, req)
	if oldfit, err := db.GetFitFromDB(ID); err == nil {
		if !logged {
			rw.Write([]byte("You must be logged in to do this"))
			return
		}
		if oldfit.Owner != user.Username && user.UserLevel != "admin" {
			rw.Write([]byte("You are not the owner of this fit the owner is " + fit.Owner))
			return
		}
		if err := db.Fittings.Update(db.GetFitDocFromID(ID), db.UnMapFit(fit)); err != nil {
			rw.Write([]byte("DB error: " + err.Error()))
			return
		}
		rw.Write([]byte("success"))
	} else {
		rw.Write([]byte("DB error while getting original fit: " + err.Error()))
	}
}

func DeleteFit(rw http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	logged, user, _ := GetSession(rw, req)
	if !logged {
		WriteError(fmt.Errorf("You must be logged in to do this"), rw, req)
		return
	}
	f := db.Fit{}
	f.FitID, _ = strconv.Atoi(vars["FitID"])
	if f.FitID != 0 {
		var err error
		f, err = db.GetFitFromDB(f.FitID)
		if err != nil {
			WriteError(err, rw, req)
			return
		}
	} else {
		WriteError(fmt.Errorf("No such fit"), rw, req)
		return
	}

	if user.Username != f.Owner && user.UserLevel != "admin" {
		WriteError(fmt.Errorf("You are not the owner of this fit."), rw, req)
		return
	}
	docid := db.GetFitDocFromID(f.FitID)
	if err := db.Fittings.Delete(docid); err != nil {
		WriteError(fmt.Errorf("DB error: %v", err.Error()), rw, req)
		return
	}
	rw.Write([]byte("success"))
}

func YourFits(rw http.ResponseWriter, req *http.Request) {
	g := GetGlobal(rw, req)
	loggedIn, _, _ := GetSession(rw, req)
	if !loggedIn {
		WriteError(fmt.Errorf("You must be logged in"), rw, req)
		return
	}
	g.Fits = db.GetFitsFromUsername(g.User.Username)
	if t, ok := Templates["fits.tmpl"]; ok {
		if err := t.Execute(rw, g); err != nil {
			log.LogError(err.Error())
		}
	} else {
		log.LogError("Template for fits was not found.")
	}
}

func RenameFit(rw http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	ID, _ := strconv.Atoi(vars["FitID"])
	data, _ := ioutil.ReadAll(req.Body)
	logged, user, _ := GetSession(rw, req)
	if oldfit, err := db.GetFitFromDB(ID); err == nil {
		if !logged {
			rw.Write([]byte("You must be logged in to do this"))
			return
		}
		if oldfit.Owner != user.Username && user.UserLevel != "admin" {
			rw.Write([]byte("You are not the owner of this fit the owner is " + oldfit.Owner))
			return
		}
		if len(data) >= 20 {
			rw.Write([]byte("Name too long.  Must be <= 20 characters was: " + strconv.Itoa(len(data))))
			return
		}
		oldfit.Name = string(data)
		if err := db.Fittings.Update(db.GetFitDocFromID(ID), db.UnMapFit(oldfit)); err != nil {
			rw.Write([]byte("DB error: " + err.Error()))
			return
		}
		rw.Write([]byte("success"))
	} else {
		rw.Write([]byte("DB error while getting original fit: " + err.Error()))
	}
}
