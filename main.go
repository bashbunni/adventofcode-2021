package main

import (
	"adventofcode/day1"
	"fmt"
)

func main() {
	fmt.Printf("result from example given: %d\n", day1.CountDepthIncreasedFromFile("./day1/sample_data.txt"))
	fmt.Printf("result from my puzzle data: %d\n", day1.CountDepthIncreasedFromFile("./day1/data.txt"))
	fmt.Printf("%d\n", day1.ThreeMeasurement())
}
