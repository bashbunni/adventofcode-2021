package day2

import (
	"adventofcode/utils"
	"log"
	"strconv"
	"strings"
)

const direction = 0
const value = 1

func Pilot(file string) int {
	scanner := utils.ReadFile(file)
	var depth int
	var position int
	for scanner.Scan() {
		input := strings.Split(scanner.Text(), " ")
		num, err := strconv.Atoi(input[value])
		if err != nil {
			log.Fatal(err)
		}
		switch input[direction] {
		case "forward":
			position += num
		case "up":
			depth -= num
		case "down":
			depth += num
		}
	}
	return depth * position
}
