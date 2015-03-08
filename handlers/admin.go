package handlers

import (
	"errors"
	"fmt"
	"github.com/THUNDERGROOVE/SDETool2/log"
	"github.com/THUNDERGROOVE/go-fit/db"
	"github.com/gorilla/mux"
	"math"
	"net/http"
	"strconv"
)

func AdminReload(rw http.ResponseWriter, req *http.Request) {
	logged, user, _ := GetSession(rw, req)
	if logged {
		if user.UserLevel == "admin" {
			go LoadTemplates() // Run in goroutine so response is faster.
			rw.Write([]byte("success"))
		} else {
			WriteError(errors.New("You must be an admin to do that"), rw, req)
		}
	} else {
		WriteError(errors.New("You must be logged in to do that."), rw, req)
	}
}

type AdminManage struct {
	Page    int
	Pages   []int
	Total   int
	UserSet []db.User
}

func (a AdminManage) Next() int {
	return a.Page + 1
}

func (a AdminManage) Previus() int {
	return a.Page - 1
}

const (
	AccountsPerPage = 10
)

func Iter(i int) []int {
	out := make([]int, 0)
	for i := 0; i < i; i++ {
		out = append(out, i)
	}
	return out
}

func AdminManageAccounts(rw http.ResponseWriter, req *http.Request) {
	g := GetGlobal(rw, req)
	logged, user, _ := GetSession(rw, req)
	if logged {
		if user.UserLevel == "admin" {
			vars := mux.Vars(req)
			page, _ := strconv.Atoi(vars["Page"])
			users := db.GetUsersOrdered()
			pages := int(math.Ceil(float64(len(users)) / float64(10)))
			if page < pages {
				page = 0
			}
			var userset []db.User
			fmt.Println(len(users))
			if len(users) < AccountsPerPage {
				userset = users
			} else if len(users)-1 > page+1*AccountsPerPage {
				fmt.Printf("[%v:%v]\n", page*AccountsPerPage, page*AccountsPerPage+AccountsPerPage)
				userset = users[page*AccountsPerPage : page*AccountsPerPage+AccountsPerPage]
			} else {
				fmt.Printf("[%v:%v]\n", page*AccountsPerPage, len(users)-1-(page*AccountsPerPage)+AccountsPerPage)
				userset = users[page*AccountsPerPage : len(users)-1-(page*AccountsPerPage)+AccountsPerPage]
			}
			g.AdminManage.Page = page
			g.AdminManage.Pages = Iter(pages)
			g.AdminManage.Total = pages
			g.AdminManage.UserSet = userset
			if t, ok := Templates["usermanage.tmpl"]; ok {
				if err := t.Execute(rw, g); err != nil {
					log.LogError(err.Error())
				}
			} else {
				log.LogError("Template for usermanage was not found.")
			}
		} else {
			WriteError(errors.New("You must be an admin to do that"), rw, req)
		}
	} else {
		WriteError(errors.New("You must be logged in to do that."), rw, req)
	}
}

func AdminFits(rw http.ResponseWriter, req *http.Request) {
	g := GetGlobal(rw, req)
	loggedIn, _, _ := GetSession(rw, req)
	if !loggedIn {
		WriteError(fmt.Errorf("You must be logged in"), rw, req)
		return
	}
	if g.User.UserLevel != "admin" {
		WriteError(fmt.Errorf("You must be an admin to use this"), rw, req)
	}

	g.Fits = db.GetAllFits()

	if t, ok := Templates["fits.tmpl"]; ok {
		if err := t.Execute(rw, g); err != nil {
			log.LogError(err.Error())
		}
	} else {
		log.LogError("Template for fits was not found.")
	}
}
