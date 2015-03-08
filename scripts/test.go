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
	t, _ := s.Search()
}
