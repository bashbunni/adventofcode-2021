package day2

import (
	"adventofcode/utils"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

const direction = 0
const value = 1

func Pilot(file string) int {
	scanner := utils.ReadFile(file)
	fmt.Println(scanner)
	var depth []int
	var position []int
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		input := strings.Split(scanner.Text(), " ")
		fmt.Println(len(input))
		num, err := strconv.Atoi(input[value])
		if err != nil {
			log.Fatal(err)
		}
		switch input[direction] {
		case "forward":
			position = append(position, num)
		case "up":
			depth = append(depth, num)
		case "down":
			depth = append(depth, -num)
		}
	}
	fmt.Println(utils.Sum(depth))
	fmt.Println(utils.Sum(position))
	return (int)(math.Abs((float64)(utils.Sum(depth) * utils.Sum(position))))
}
