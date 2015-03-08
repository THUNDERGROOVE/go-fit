package db

import (
	"github.com/THUNDERGROOVE/SDETool2/sde"
)

type User struct {
	Username  string
	PassHash  string
	UserLevel string
}

type Fit struct {
	TypeID         int
	HighSlot1      int
	HighSlot2      int
	HighSlot3      int
	HighSlot4      int
	HighSlot5      int
	LowSlot1       int
	LowSlot2       int
	LowSlot3       int
	LowSlot4       int
	LowSlot5       int
	WeaponSlot1    int
	WeaponSlot2    int
	WeaponSlot3    int
	EquipmentSlot1 int
	EquipmentSlot2 int
	EquipmentSlot3 int
	EquipmentSlot4 int
	Owner          string
	FitID          int
	Name           string
	SDEType        sde.SDEType
}

func MapFit(data map[string]interface{}) Fit {
	var f Fit
	f.TypeID = int(data["TypeID"].(float64))
	f.HighSlot1 = int(data["HighSlot1"].(float64))
	f.HighSlot2 = int(data["HighSlot2"].(float64))
	f.HighSlot3 = int(data["HighSlot3"].(float64))
	f.HighSlot4 = int(data["HighSlot4"].(float64))
	f.HighSlot5 = int(data["HighSlot5"].(float64))
	f.LowSlot1 = int(data["LowSlot1"].(float64))
	f.LowSlot2 = int(data["LowSlot2"].(float64))
	f.LowSlot3 = int(data["LowSlot3"].(float64))
	f.LowSlot4 = int(data["LowSlot4"].(float64))
	f.LowSlot5 = int(data["LowSlot5"].(float64))
	f.WeaponSlot1 = int(data["WeaponSlot1"].(float64))
	f.WeaponSlot2 = int(data["WeaponSlot2"].(float64))
	f.WeaponSlot3 = int(data["WeaponSlot3"].(float64))
	f.EquipmentSlot1 = int(data["EquipmentSlot1"].(float64))
	f.EquipmentSlot2 = int(data["EquipmentSlot2"].(float64))
	f.EquipmentSlot3 = int(data["EquipmentSlot3"].(float64))
	f.EquipmentSlot4 = int(data["EquipmentSlot4"].(float64))
	f.Owner = data["Owner"].(string)
	f.Name = data["Name"].(string)
	f.FitID = int(data["FitID"].(float64))
	f.SDEType, _ = SDE.GetType(f.TypeID)

	return f
}

func UnMapFit(f Fit) map[string]interface{} {
	data := make(map[string]interface{}, 0)
	data["TypeID"] = f.TypeID
	data["HighSlot1"] = f.HighSlot1
	data["HighSlot2"] = f.HighSlot2
	data["HighSlot3"] = f.HighSlot3
	data["HighSlot4"] = f.HighSlot4
	data["HighSlot5"] = f.HighSlot5
	data["LowSlot1"] = f.LowSlot1
	data["LowSlot2"] = f.LowSlot2
	data["LowSlot3"] = f.LowSlot3
	data["LowSlot4"] = f.LowSlot4
	data["LowSlot5"] = f.LowSlot5
	data["WeaponSlot1"] = f.WeaponSlot1
	data["WeaponSlot2"] = f.WeaponSlot2
	data["WeaponSlot3"] = f.WeaponSlot3
	data["EquipmentSlot1"] = f.EquipmentSlot1
	data["EquipmentSlot2"] = f.EquipmentSlot2
	data["EquipmentSlot3"] = f.EquipmentSlot3
	data["EquipmentSlot4"] = f.EquipmentSlot4
	data["Owner"] = f.Owner
	data["FitID"] = f.FitID
	data["Name"] = f.Name
	return data
}

func MapUser(data map[string]interface{}) User {
	var u User
	u.Username = data["username"].(string)
	u.PassHash = data["password"].(string)
	u.UserLevel = data["accountType"].(string)
	return u
}
