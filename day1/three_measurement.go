package day1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func ThreeMeasurement() {
	// A: [1, 2, 3]
	// go through data line by line, then add the value on that line to the letter's value
	f, err := os.Open("./day1/sample_data2.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	values := make(map[string]int)
	for scanner.Scan() {
		val := strings.Fields(scanner.Text())
		var currValue int
		for i := 0; i < len(val); i++ {
			num, _ := strconv.Atoi(val[i])
			if num == 0 {
				// it's a letter
				if _, ok := values[val[i]]; ok {
					values[val[i]] += currValue
				} else {
					values[val[i]] = currValue
				}
			} else {
				currValue = num
			}
		}
	}
	fmt.Print("A" < "B")
}
