package day3

import (
	"adventofcode/utils"
	"log"
	"strconv"
	"strings"
)

/*
 * PART 1
 */

// ProcessReport returns the power consumption of the submarine as int64
func ProcessReport(file string) int64 {
	report := utils.ConvertDataTo2DArray(file)
	MLCommonBits := getMostAndLeastCommonBits(report)
	return calculatePower(binaryToDecimal(MLCommonBits["common"]), binaryToDecimal(MLCommonBits["uncommon"]))
}



func calculatePower(gamma, epsilon int64) int64 {
	return gamma * epsilon
}

func binaryToDecimal(binary string) int64 {
	result, err := strconv.ParseInt(binary, 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	return result
}

func getMostAndLeastCommonBits(binary [][]string) map[string]string {
	output := make(map[string]string)
	output["common"] = ""
	output["uncommon"] = ""
	for i := 0; i < len(binary[i]); i++ {
		sum := 0
		for j := 0; j < len(binary); j++ {
			if binary[j][i] == "1" {
				sum++
			}
		}
		if sum > (len(binary) / 2) {
			output["common"] += "1"
			output["uncommon"] += "0"
		} else {

			output["common"] += "0"
			output["uncommon"] += "1"
		}
	}
	return output
}

/*
 * PART 2
 */

func CalculateLifeSupportRating(file string) int64 {
	report := utils.ConvertDataTo2DArray(file)
	oxygen := getOxygenGeneratorReading(report, 0)
	co2 := getCO2ScrubberRating(report, 0)
	return binaryToDecimal(co2) * binaryToDecimal(oxygen)
}

// go over one column, change values you're working with, go to second column, change values you're working with... until you've got one remaining value or you reach the last column

// getOxygenGeneratorReading gets the values with have the most common bits
func getOxygenGeneratorReading(values [][]string, col int) string {
	if len(values) == 1 {
		return strings.Join(values[0], "")
	}
	var ones [][]string
	var zeros [][]string
	for row := 0; row < len(values); row++ {
		if values[row][col] == "1" {
			ones = append(ones, values[row])
		} else {
			zeros = append(zeros, values[row])
		}
	}
	if len(ones) >= len(zeros) {
		return getOxygenGeneratorReading(ones, col+1)
	} else {
		return getOxygenGeneratorReading(zeros, col+1)
	}
}

// getCO2ScrubberRating gets the values with have the least common bits
func getCO2ScrubberRating(values [][]string, col int) string {
	if len(values) == 1 {
		return strings.Join(values[0], "")
	}
	var ones [][]string
	var zeros [][]string
	for row := 0; row < len(values); row++ {
		if values[row][col] == "1" {
			ones = append(ones, values[row])
		} else {
			zeros = append(zeros, values[row])
		}
	}
	if len(ones) < len(zeros) {
		return getCO2ScrubberRating(ones, col+1)
	} else {
		return getCO2ScrubberRating(zeros, col+1)
	}
}
