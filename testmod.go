package main

import (
	"fmt"

	"github.com/iKOPKACtraxa/testmod/v2"
)

func main() {
	// fmt.Println(testmod.Hi("roberto"))
	g, err := testmod.Hi("Roberto", "pt")
	if err != nil {
		panic(err)
	}
	fmt.Println(g)
}
