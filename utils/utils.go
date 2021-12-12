package utils

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func ReadFile(file string) (scanner *bufio.Scanner) {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	return bufio.NewScanner(f)
}

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

func ConvertDataTo2DArray(file string) [][]string {
	var report [][]string
	scanner := ReadFile(file)
	for scanner.Scan() {
		row := strings.Split(scanner.Text(), "")
		report = append(report, row)
	}
	return report
}
