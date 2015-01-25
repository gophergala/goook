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

	test := Class{111, "Hello"}
	fmt.Println(test)
	fmt.Println(test.Depth())

	nb := NewNaiveBayes()
	nb.Learn(1337, []string{"hello", "world", "fun"})
	nb.Learn(666, []string{"viagra", "fat", "prince"})

	fmt.Println(nb.Classify([]string{"hello", "world"}))
	fmt.Println(nb.Classify([]string{"hello", "viagra"}))
	fmt.Println(nb.Classify([]string{"viagra", "prince"}))
	fmt.Println(nb.Classify([]string{"hello", "viagra", "prince"}))

	tree := NewTreeBayes()

	fmt.Println("Learning")
	for id := range ddc.Classes {
		tree.Learn(id, []string{fmt.Sprintf("%d", id)})
	}

	fmt.Println("Classifying")
	class, prob := tree.Classify([]string{"100", "120", "123"}, .5)
	fmt.Println("Class =", class, " / P = ", prob)
}
