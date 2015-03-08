package main

import (
	"fmt"
	"github.com/THUNDERGROOVE/SDETool2/sde"
	"github.com/THUNDERGROOVE/SDETool2/types/tags"
	"os"
)

func main() {
	s, err := sde.Open("wl-uf-latest")
	if err != nil {
		fmt.Println(err.Error())
	}
	os.Remove("modules.js")
	if f, err := os.OpenFile("modules.js", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0777); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("File openend.  Collecting dropsuit weapon upgrades.")
		fmt.Fprintf(f, "var WeaponUpgrades = [\n")
		for _, v := range s.GetTypesWithTag(tags.Tag_module_weaponupgrade) {
			if v.HasTag(tags.Tag_moduleclass_infantry) && !v.IsFaction() && !v.IsAurum() {
				fmt.Fprintf(f, "\t{value: %v, text: \"%v\"},\n", v.TypeID, v.GetName())
			}
		}
		fmt.Fprintf(f, "];\n")

		fmt.Println("Collecting plate upgrades.")
		fmt.Fprintf(f, "\nvar ArmorPlates = [\n")
		for _, v := range s.GetTypesWithTag(tags.Tag_module_armor) {
			if v.HasTag(tags.Tag_moduleclass_infantry) && !v.IsFaction() && !v.IsAurum() {
				fmt.Fprintf(f, "\t{value: %v, text: \"%v\"},\n", v.TypeID, v.GetName())
			}
		}
		fmt.Fprintf(f, "];\n")

		fmt.Println("Collecting shield upgrades .")
		fmt.Fprintf(f, "\nvar ShieldExtenders = [\n")
		for _, v := range s.GetTypesWithTag(tags.Tag_module_shield) {
			if v.HasTag(tags.Tag_moduleclass_infantry) && !v.IsFaction() && !v.IsAurum() {
				fmt.Fprintf(f, "\t{value: %v, text: \"%v\"},\n", v.TypeID, v.GetName())
			}
		}
		fmt.Fprintf(f, "];\n")

		fmt.Println("Collecting biotic upgrades .")
		fmt.Fprintf(f, "\nvar BioticUpgrades = [\n")
		for _, v := range s.GetTypesWithTag(tags.Tag_module_biotic) {
			if v.HasTag(tags.Tag_moduleclass_infantry) && !v.IsFaction() && !v.IsAurum() {
				fmt.Fprintf(f, "\t{value: %v, text: \"%v\"},\n", v.TypeID, v.GetName())
			}
		}
		fmt.Fprintf(f, "];\n")

		fmt.Println("Collecting electronics upgrades .")
		fmt.Fprintf(f, "\nvar ElectronicUpgrades = [\n")
		for _, v := range s.GetTypesWithTag(tags.Tag_module_electronics) {
			if v.HasTag(tags.Tag_moduleclass_infantry) && !v.IsFaction() && !v.IsAurum() {
				fmt.Fprintf(f, "\t{value: %v, text: \"%v\"},\n", v.TypeID, v.GetName())
			}
		}
		fmt.Fprintf(f, "];\n")

		fmt.Println("Collecting enginering upgrades .")
		fmt.Fprintf(f, "\nvar EnginerringUpgrades = [\n")
		for _, v := range s.GetTypesWithTag(tags.Tag_module_engineering) {
			if v.HasTag(tags.Tag_moduleclass_infantry) && !v.IsFaction() && !v.IsAurum() {
				fmt.Fprintf(f, "\t{value: %v, text: \"%v\"},\n", v.TypeID, v.GetName())
			}
		}
		fmt.Fprintf(f, "];\n")

	}
}
