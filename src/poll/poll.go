package poll

import "fmt"

// ConfidenceInterval - Struct which holds interval values
type ConfidenceInterval struct {
	a float64
	b float64
}

var p PStruct

// PStruct - Struct which holds poll's values
type PStruct struct {
	populationSize int
	sampleSize int
	votingIntentions float64
	variance float64
	confidenceInterval95 ConfidenceInterval
	confidenceInterval99 ConfidenceInterval
}

func displayResults() {
	fmt.Printf("Population size:                   %d\n", p.populationSize)
	fmt.Printf("Sample size:                       %d\n", p.sampleSize)
	fmt.Printf("Voting intentions:                 %.2f%%\n", p.votingIntentions)
	fmt.Printf("Variance:                          %.6f\n", p.variance)
	fmt.Printf("95%% confidence interval:           [%.2f%%; %.2f%%]\n", p.confidenceInterval95.a, p.confidenceInterval95.b)
	fmt.Printf("99%% confidence interval:           [%.2f%%; %.2f%%]\n", p.confidenceInterval99.a, p.confidenceInterval99.b)
}

// Poll - main
func Poll() {
	computeConfidentInterval(95)
	computeConfidentInterval(99)

	displayResults()
}