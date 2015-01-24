package main

type LearnTree struct {
	Depth      int
	Classifier *NaiveBayes
	Children   []*LearnTree
}

func NewLearnTree(depth, arity int) *LearnTree {
	if depth == 0 {
		return nil
	}

	var nc NaiveBayes
	t := &LearnTree{
		Depth:      depth,
		Classifier: &nc,
		Children:   make([]*LearnTree, arity),
	}
	for i := 0; i < arity; i++ {
		t.Children[i] = NewLearnTree(depth-1, arity)
	}
	return t
}
