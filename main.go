package main

import (
	"fmt"
	"log"
)

func main() {
	ddc, err := LoadDDC("data/classes_clean")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(ddc)

	test := Class{111, "Hello"}
	fmt.Println(test)
	fmt.Println(test.Depth())
}
