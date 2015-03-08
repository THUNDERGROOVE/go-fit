package db

import (
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/HouzuoGuo/tiedot/db"
	"github.com/THUNDERGROOVE/SDETool2/log"
	"github.com/THUNDERGROOVE/SDETool2/sde"
	"github.com/THUNDERGROOVE/go-fit/args"
	"sort"
)

var DB *db.DB
var SDE sde.SDE
var Users *db.Col
var Fittings *db.Col

func init() {
	var err error
	DB, err = db.OpenDB("database/")
	if err != nil {
		log.LogError(err.Error())
		return
	}
	var err1 error
	SDE, err1 = sde.Open(*args.Version)
	if err != nil {
		log.LogError("Error in SDE initialization", err1.Error())
		return
	}

	if !contains("Users", DB.AllCols()) {
		log.Info("Users didn't exist yet.  Creating..")
		if err := DB.Create("Users"); err != nil {
			panic(err)
		}
	}

	Users = DB.Use("Users")
	if err := Users.Index([]string{"username"}); err != nil {
		log.LogError(err.Error())
	}
	gob.Register(User{})

	if !contains("Fittings", DB.AllCols()) {
		log.Info("Fittings didn't exist yet.  Creating..")
		if err := DB.Create("Fittings"); err != nil {
			panic(err)
		}
	}

	Fittings = DB.Use("Fittings")
	if err := Fittings.Index([]string{"FitID"}); err != nil {
		log.LogError(err.Error())
	}
	if err := Fittings.Index([]string{"Owner"}); err != nil {
		log.LogError(err.Error())
	}
	SetNames()
}

func NewUser(username, password, userlevel string) error {
	hash := PassHashFunc(password)
	_, err := Users.Insert(map[string]interface{}{
		"username":    username,
		"password":    hash,
		"accountType": userlevel,
	})

	return err
}

func GetUsersOrdered() []User {
	out := make([]User, 0)
	u := make(map[string]User, 0)
	names := make([]string, 0)
	Users.ForEachDoc(func(id int, doc []byte) bool {
		if data, err := Users.Read(id); err == nil {
			us := MapUser(data)
			u[us.Username] = us
			names = append(names, us.Username)
		} else {
			log.LogError(err.Error())
		}
		return true
	})
	sort.Strings(names)
	for _, v := range names {
		out = append(out, u[v])
	}
	return out
}

func VerifyCredientialsHash(username, password string) (User, bool) {
	return VerifyCredientials(username, PassHashFunc(password))
}

func VerifyCredientials(username, password string) (User, bool) {
	user := &User{}
	var query interface{}
	json.Unmarshal([]byte(fmt.Sprintf(`[{"eq": "%v", "in": ["username"]}]`, username)), &query)
	res := make(map[int]struct{})
	if err := db.EvalQuery(query, Users, &res); err != nil {
		log.LogError(err.Error())
		return *user, false
	}

	for id := range res {
		if re, err := Users.Read(id); err == nil {
			user.PassHash = re["password"].(string)
			user.Username = re["username"].(string)
			user.UserLevel = re["accountType"].(string)
			if user.PassHash == password {
				return *user, true
			} else {
				return *user, false
			}
		} else {
			log.LogError(err.Error())
		}
	}
	return *user, false
}

// Helpers

func contains(s string, ss []string) bool {
	for _, v := range ss {
		if v == s {
			return true
		}
	}
	return false
}

func PassHashFunc(password string) string {
	s := sha256.New()
	s.Write([]byte(password))
	return hex.EncodeToString(s.Sum([]byte("")))
}

func SetNames() {
	Fittings.ForEachDoc(func(id int, doc []byte) bool {
		data, err := Fittings.Read(id)
		if err != nil {
			log.LogError(err.Error())
			return false
		}
		if _, ok := data["Name"]; ok {

		} else {
			t, _ := SDE.GetType(int(data["TypeID"].(float64)))
			fmt.Println("Fixing blank fit name")
			data["Name"] = t.GetName()
			if err := Fittings.Update(id, data); err != nil {
				fmt.Println(err.Error())
			}
		}
		return true
	})
}
