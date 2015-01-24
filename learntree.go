package main

type LearnTree struct {
	Depth      int
	Classifier *NaiveBayes
	Children   map[int]*LearnTree
}

func NewLearnTree(depth int) *LearnTree {
	if depth == 0 {
		return nil
	}

	var nc NaiveBayes
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
