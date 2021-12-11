package day3

import "adventofcode/utils"

// Returns the power consumption of the submarine
func ProcessReport(file string) int {
	var gammaRate string     // binary
	var epsilonRate string   // binary
	var powerConsumption int // gamma * epsilon

	var report [][]string
	scanner := utils.ReadFile(file)
	// build dataset
	for scanner.Scan() {
		row := strings.Split(scanner.Text(), "")
		report = append(report, row)
	}
	for row := 0; row < len(report); row++ {
		// sum values 1...N
		// 
	}

	// gammaRate: most common bit in the corresponding position
}

// looking at the columns
// create a 2D array
