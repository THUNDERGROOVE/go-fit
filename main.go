// GoFit is a online fitting tool for DUST514 that takes advantage of SDETool2's sde package.
package main

import (
	"bufio"
	"fmt"
	"github.com/THUNDERGROOVE/SDETool2/log"
	"github.com/THUNDERGROOVE/go-fit/args"
	"github.com/THUNDERGROOVE/go-fit/db"
	"github.com/THUNDERGROOVE/go-fit/handlers"
	ghandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func main() {

	log.InfoLog = true
	log.ErrLog = true

	if *args.NewAccount {
		fmt.Print("Username:\n> ")
		username := RawInput()
		fmt.Print("Password:\n> ")
		password := RawInput()
		fmt.Println("AccoutnType:\n> ")
		accountType := RawInput()
		if err := db.NewUser(username, password, accountType); err != nil {
			log.LogError(err.Error())
		}
		return
	}

	handlers.LoadTemplates()

	log.Info("Starting http server on localhost:" + strconv.Itoa(*args.Port))

	r := mux.NewRouter()
	r.NotFoundHandler = http.HandlerFunc(handlers.FourOhFour)
	// Handlers
	r.HandleFunc("/login", handlers.Login)
	r.HandleFunc("/logout", handlers.Logout)

	r.HandleFunc("/fit", handlers.Fitting)
	r.HandleFunc("/fits", handlers.YourFits)
	r.HandleFunc("/newfit", handlers.NewFit)
	r.HandleFunc("/rawfit/{FitID:[0-9]+}", handlers.RawFit)
	r.HandleFunc("/fit/{FitID:[0-9]+}", handlers.Fitting)
	r.HandleFunc("/fit/save/{FitID:[0-9]+}", handlers.SaveFit)
	r.HandleFunc("/fit/delete/{FitID:[0-9]+}", handlers.DeleteFit)
	r.HandleFunc("/fit/rename/{FitID:[0-9]+}", handlers.RenameFit)

	// Raw shit
	r.HandleFunc("/typename/{TypeID:[0-9]+}", handlers.TypeName)

	// Admin shit
	r.HandleFunc("/admin/reload", handlers.AdminReload)
	r.HandleFunc("/admin/manageusers/{Page:[0-9]+}", handlers.AdminManageAccounts)
	r.HandleFunc("/admin/fits", handlers.AdminFits)

	r.HandleFunc("/", handlers.Index)
	http.Handle("/", ghandlers.CombinedLoggingHandler(ioutil.Discard, r))

	// Static Content
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("static/")))
	fmt.Println("Server starting..")
	http.ListenAndServe(fmt.Sprintf(":%v", *args.Port), nil)
}

func RawInput() string {
	in := bufio.NewReader(os.Stdin)
	line, _ := in.ReadString('\n')
	line = strings.Trim(line, "\n\t\r")
	return line
}
