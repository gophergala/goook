package main

type NaiveBayes map[string]map[string]uint

func (nb NaiveBayes) Learn(class string, words []string) {
	for _, word := range words {
		wordModel, exists := nb[word]
		if !exists {
			wordModel = make(map[string]uint)
			nb[word] = wordModel
		}

		wordModel[class]++
	}
}

func (nb NaiveBayes) Classify(words []string) map[string]float64 {
	total := uint(0)
	counts := make(map[string]uint)

	for _, word := range words {
		wordModel, exists := nb[word]
		if !exists {
			continue
		}

		for class, count := range wordModel {
			counts[class] += count
			total += count
		}
	}

	probabilities := make(map[string]float64)
	for class, count := range counts {
		probabilities[class] = float64(count) / float64(total)
	}

	return probabilities
}
