package day1

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func ThreeMeasurement() int {
	// A: [1, 2, 3]
	// go through data line by line, then add the value on that line to the letter's value
	data, err := os.ReadFile("./day1/data.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(data), "\n")
	nums := stringToInt(lines)
	// grab all the numbers from the file
	// get the sum of three sequential numbers
	// get the sum of i+1 three sequential numbers
	// check if sum1 < sum2 ? increase++
	increased := 0
	for i := 0; i < len(nums)-3; i++ {
		sum1 := nums[i] + nums[i+1] + nums[i+2]
		sum2 := nums[i+1] + nums[i+2] + nums[i+3]
		if sum1 < sum2 {
			increased++
		}
	}
	return increased
}

func stringToInt(strings []string) []int {
	var nums []int
	for _, value := range strings {
		num, _ := strconv.Atoi(value)
		nums = append(nums, num)
	}
	return nums
}

func contains(nums []int, target int) bool {
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			return true
		}
	}
	return false
}
