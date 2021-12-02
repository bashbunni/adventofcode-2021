package utils

import (
	"os"
	"bufio"
	"log"
	"strconv"
)

func ReadDataToInts(file string) ([]int, error) {
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
