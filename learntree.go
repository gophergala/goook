package main

import (
	"fmt"
)

type LearnTree struct {
	Height          int
	NodeHeights     map[int]int
	NodeClassifiers map[int]*NaiveBayes
}

func pow(a, b int) int {
	res := a
	for i := 1; i < b; i++ {
		res *= a
	}
	return res
}

func NewLearnTree(height int) *LearnTree {

	tree := &LearnTree{
		Height:          height,
		NodeHeights:     make(map[int]int),
		NodeClassifiers: make(map[int]*NaiveBayes),
	}

	// All the classes that are in the order of `depth` are leaves, so
	// they won't contain a classifier.
	for i := 0; i < pow(10, height-1); i++ {
		fmt.Println(i)
		tree.NodeClassifiers[i] = NewNaiveBayes()
	}

	return tree
}

func (tree *LearnTree) TrainTree(path int, data map[int][]string) {
	// For each node that contains a classifier
	for path, classifier := range tree.NodeClassifiers {
		// If the node/path is 22, we shall train the model on the classes like 22*
		for classID, words := range data {
			if classID/tree.NodeHeights[path] == path {
				classifier.Learn(classID, words)
			}
		}
	}
}
