package db

import (
	"encoding/json"
	"fmt"
	"github.com/HouzuoGuo/tiedot/db"
	"github.com/THUNDERGROOVE/SDETool2/log"
)

func GetFitsFromUsername(username string) []Fit {
	fits := make([]Fit, 0)
	var query interface{}
	json.Unmarshal([]byte(fmt.Sprintf(`[{"eq": "%v", "in": ["Owner"]}]`, username)), &query)
	res := make(map[int]struct{})
	if err := db.EvalQuery(query, Fittings, &res); err != nil {
		log.LogError(err.Error())
		return fits
	}

	for ID := range res {
		if doc, err := Fittings.Read(ID); err == nil {
			fits = append(fits, MapFit(doc))
		} else {
			log.LogError(err.Error())
		}
	}
	return fits
}

func GetFitFromDB(ID int) (Fit, error) {
	fit := Fit{}
	Fittings.ForEachDoc(func(id int, doc []byte) bool {
		data, err := Fittings.Read(id)
		if err != nil {
			log.LogError(err.Error())
			return true
		}
		fit = MapFit(data)
		if fit.FitID == ID {
			return false
		}
		return true
	})
	return fit, nil
}

func NewFit(username string, TypeID int) (int, error) {
	t, _ := SDE.GetType(TypeID)
	id := GetCleanFitID()
	log.Info("GetCleanFitID returned", id)
	doc, err := Fittings.Insert(map[string]interface{}{
		"TypeID":         TypeID,
		"HighSlot1":      0,
		"HighSlot2":      0,
		"HighSlot3":      0,
		"HighSlot4":      0,
		"HighSlot5":      0,
		"LowSlot1":       0,
		"LowSlot2":       0,
		"LowSlot3":       0,
		"LowSlot4":       0,
		"LowSlot5":       0,
		"WeaponSlot1":    0,
		"WeaponSlot2":    0,
		"WeaponSlot3":    0,
		"EquipmentSlot1": 0,
		"EquipmentSlot2": 0,
		"EquipmentSlot3": 0,
		"EquipmentSlot4": 0,
		"Owner":          username,
		"Name":           t.GetName(),
		"FitID":          id,
	})
	d, _ := Fittings.Read(doc)
	fmt.Println(d)
	// for k, v := range d["Fit"].(map[string]interface{}) {
	// 	fmt.Printf("k: %v v: %v\n", k, v)
	// }

	return id, err
}

func GetCleanFitID() int {
	i := 1

	Fittings.ForEachDoc(func(id int, doc []byte) bool {
		data, _ := Fittings.Read(id)
		log.Info("Checking out fit", id, int(data["FitID"].(float64)))
		if int(data["FitID"].(float64)) >= i {
			i = int(data["FitID"].(float64)) + 1
		}
		return true
	})

	return i
}

func GetFitDocFromID(ID int) int {
	fit := Fit{}
	var i int
	Fittings.ForEachDoc(func(id int, doc []byte) bool {
		data, err := Fittings.Read(id)
		if err != nil {
			log.LogError(err.Error())
			return true
		}
		fit = MapFit(data)
		if fit.FitID == ID {
			i = id
			return false
		}
		return true
	})
	return i
}

func GetAllFits() []Fit {
	out := make([]Fit, 0)
	Fittings.ForEachDoc(func(id int, doc []byte) bool {
		data, err := Fittings.Read(id)
		if err != nil {
			log.LogError(err.Error())
			return true
		}
		fit := MapFit(data)
		out = append(out, fit)
		return true
	})
	return out
}
