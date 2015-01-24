package main

import "fmt"

type LearnTree struct {
	Depth      int
	Classifier *NaiveBayes
	Children   map[int]*LearnTree
}

func NewLearnTree(depth int) *LearnTree {
	if depth < 0 {
		return nil
	}

	var nc NaiveBayes // TODO:make
	t := &LearnTree{
		Depth:      depth,
		Classifier: &nc,
		Children:   make(map[int]*LearnTree),
	}

	for i := 0; i < 10; i++ {
		t.Children[i] = NewLearnTree(depth - 1)
	}

	return t
}

func (classifier *NaiveBayes) trainForPath(path, depth int, data map[int][]string) {

	for classID, _ := range data {
		if classID/10^depth == path {
			// classifier.Learn(cl, words)
			fmt.Printf("Treated path %d\n", path)
		}
	}
}

func (tree *LearnTree) TrainTree(path int, data map[int][]string) {
	if tree.Depth <= 0 {
		return
	}

	for i, child := range tree.Children {
		child.Classifier.trainForPath(path*10+i, tree.Depth, data)
		child.TrainTree(path*10+i, data)
	}
}
