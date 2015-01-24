package main

import (
	"github.com/sjwhitworth/golearn/base"
)

type ModelTree struct {
	DDCLevel int
	Model    base.Classifier
	Children []*ModelTree
}

func NewModelTree(levels, arity int, classifier base.Classifier) *ModelTree {
	if levels == 0 {
		return nil
	}

	t := ModelTree{levels, classifier, make([]*ModelTree, arity)}

	for i := 0; i < arity; i++ {
		t.Children[i] = NewModelTree(levels-1, arity, classifier)
	}
	return &t
}
