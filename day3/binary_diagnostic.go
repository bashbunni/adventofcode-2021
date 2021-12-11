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
	report := convertDataToArray(file)
	MLCommonBits := getMostAndLeastCommonBits(report)

	return calculatePower(binaryToDecimal(MLCommonBits["common"]), binaryToDecimal(MLCommonBits["uncommon"]))
}

func convertDataToArray(file string) [][]string {
	var report [][]string
	scanner := utils.ReadFile(file)
	for scanner.Scan() {
		row := strings.Split(scanner.Text(), "")
		report = append(report, row)
	}
	return report
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

func CalculateLifeSupportRating() {

}
