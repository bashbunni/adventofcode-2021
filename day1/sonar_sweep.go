package day1

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
