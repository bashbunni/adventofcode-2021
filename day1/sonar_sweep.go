package day1

import (
	"adventofcode/utils"
	"log"
)

// Counts the number of times a depth measurement increases from the previous measurement
// Question 1 of adventofcode
func CountDepthIncreased(depths []int) int {
	// only go to second last element with i so we don't go out of bounds
	count := 0
	for i := 0; i < len(depths)-1; i++ {
		if depths[i] < depths[i+1] {
			count++
		}
	}
	return count
}

// Counts the number of times a depth measurement increases from the previous measurement from a file.
func CountDepthIncreasedFromFile(file string) int {
	values, err := utils.ReadDataToInts(file)
	if err != nil {
		log.Fatal(err)
	}
	return CountDepthIncreased(values)
}
