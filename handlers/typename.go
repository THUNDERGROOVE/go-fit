package handlers

import (
	"github.com/THUNDERGROOVE/go-fit/db"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func TypeName(rw http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	TypeID, _ := strconv.Atoi(vars["TypeID"])
	t, err := db.SDE.GetType(TypeID)
	if err != nil {
		rw.Write([]byte("No such type."))
		return
	}
	t.GetAttribute("mDisplayName")
	rw.Write([]byte(t.GetName()))
}
