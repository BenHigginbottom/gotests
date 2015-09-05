package main

import	. "bayesian"


const (
	Short Class = "Good"
	Long Class = "Bad"
)

classifier := NewClassifier(Good, Bad)
goodStates := []string{"E","L"}
badStates := []string{"M","H"}
classifier.Learn(goodStates, Good)
classifier.Learn(badStates, Bad)

probs, likely, _ := classifier.ProbScores(
						[]string{"H","H","H","M","E","M","L","L","L","E","L","E"}
					)