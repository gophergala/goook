package main

import "fmt"

type TreeBayes []*NaiveBayes

func NewTreeBayes() TreeBayes {
	return make([]*NaiveBayes, 101)
}

func (tb TreeBayes) Learn(class int, words []string) {
	curr, prev := class/10+1, class+1
	for {
		cls := tb[curr]
		if cls == nil {
			cls = NewNaiveBayes()
			tb[curr] = cls
		}

		cls.Learn(prev, words)
		if curr <= 10 {
			curr, prev = 0, curr
		} else {
			curr, prev = curr/10+1, curr
		}
		if prev == 0 {
			break
		}
	}
}

func (tb TreeBayes) Classify(words []string, threshold float64) (int, float64) {
	p := float64(1.0)
	pos := 0
	for pos < len(tb) {
		fmt.Println(pos, " / ", p)
		cls := tb[pos]
		if cls == nil {
			break
		}

		ps := cls.Classify(words)
		if ps[0].P < threshold {
			break
		}

		p *= ps[0].P
		pos = ps[0].Class
	}

	return pos - 1, p
}
