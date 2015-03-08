package main

import (
	"fmt"
	"github.com/THUNDERGROOVE/SDETool2/sde"
)

func main() {
	s, err := sde.Open("wl-uf-latest")
	if err != nil {
		fmt.Println(err.Error())
	}
	ss, err1 := s.GetTypeWhereNameContains("Commando")
	if err1 != nil {
		fmt.Println(err.Error())
	}
	for _, v := range ss {
		v.GetAttributes()
		if !v.IsAurum() && !v.IsFaction() {
			fmt.Printf("{value: %v, text: \"%v\"},\n", v.TypeID, v.GetName())
		}
	}
}
