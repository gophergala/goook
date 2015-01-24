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

	_ = ddc

	test := Class{111, "Hello"}
	fmt.Println(test)
	fmt.Println(test.Depth())

	nb := make(NaiveBayes)
	nb.Learn("ham", []string{"hello", "world", "fun"})
	nb.Learn("spam", []string{"viagra", "fat", "prince"})

	fmt.Println(nb.Classify([]string{"hello", "world"}))
	fmt.Println(nb.Classify([]string{"hello", "viagra"}))
	fmt.Println(nb.Classify([]string{"viagra", "prince"}))

	tree := NewLearnTree(3)
	fmt.Println(*tree)
}
