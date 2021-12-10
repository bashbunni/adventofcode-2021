package main

import (
	"adventofcode/day2"
	"fmt"
)

func main() {
	/*
		day1 data
		fmt.Printf("result from example given: %d\n", day1.CountDepthIncreasedFromFile("./day1/sample_data.txt"))
		fmt.Printf("result from my puzzle data: %d\n", day1.CountDepthIncreasedFromFile("./day1/data.txt"))
		fmt.Printf("%d\n", day1.ThreeMeasurement())
	*/
	fmt.Println(day2.PilotWithAim("./day2/sample_data.txt"))
	fmt.Println(day2.PilotWithAim("./day2/data.txt"))
}
