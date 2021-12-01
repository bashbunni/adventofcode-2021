package day1

import (
	"bufio"
	"log"
	"os"
	"strconv"
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

func CountDepthIncreasedFromFile(file string) int {
	values, err := readData(file)
	if err != nil {
		log.Fatal(err)
	}
	// only go to second last element with i so we don't go out of bounds
	count := 0
	for i := 0; i < len(values)-1; i++ {
		if values[i] < values[i+1] {
			count++
		}
	}
	return count
}

func readData(file string) ([]int, error) {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	values := []int{}
	for scanner.Scan() {
		val, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return values, err
		}
		values = append(values, val)
	}
	return values, scanner.Err()
}
