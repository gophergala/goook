package main

type NaiveBayes interface {
	Learn(class string, words []string)
	Classify(words []string) map[string]float64
}
