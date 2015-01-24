package main

type NaiveBayes struct {
	classes []string
	counts  map[int]uint
	models  map[int]map[string]uint
}

func NewNaiveBayes() *NaiveBayes {
	return &NaiveBayes{
		counts: make(map[int]uint),
		models: make(map[int]map[string]uint),
	}
}

func (nb *NaiveBayes) GetClassLabel(id int) string {
	if id < len(nb.classes) {
		return nb.classes[id]
	}

	return ""
}

func (nb *NaiveBayes) Learn(class string, words []string) {
	classID := -1
	for id, label := range nb.classes {
		if class == label {
			classID = id
			break
		}
	}
	if classID == -1 {
		classID = len(nb.classes)
		nb.classes = append(nb.classes, class)
	}

	classModel, exists := nb.models[classID]
	if !exists {
		classModel = make(map[string]uint)
		nb.models[classID] = classModel
	}

	for _, word := range words {
		classModel[word]++
		nb.counts[classID]++
	}
}

func (nb *NaiveBayes) Classify(words []string) map[int]float64 {
	probs := make(map[int]float64)
	var norm float64
	for id, model := range nb.models {
		var p float64
		for _, word := range words {
			p += float64(model[word])
		}

		p /= float64(nb.counts[id])

		probs[id] = p
		norm += p
	}

	for id, p := range probs {
		probs[id] = p / norm
	}

	return probs
}
