package main

type LearnTree struct {
	Depth      int
	Classifier *NaiveBayes
	Children   []*LearnTree
}
