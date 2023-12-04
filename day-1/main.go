package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	var fileName string = "input.txt"
	var result int = 0

	lines := readFile(fileName)

	for _, line := range lines {
		result += parseToNr(getCalibrationValue(line))
	}

	println("Result: ", result)
}

func readFile(fileName string) []string {
	var lines []string
	data, err := os.Open(fileName)
	check(err)
	fileScanner := bufio.NewScanner(data)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}
	return lines
}

func getCalibrationValue(line string) string {
	var firstNr, lastNr string

	for _, char := range line {
		if isNrChar(string(char)) {
			if firstNr == "" {
				firstNr = string(char)
			}
			lastNr = string(char)
		}
	}
	return firstNr + lastNr
}

func parseToNr(char string) int {
	i, err := strconv.Atoi(char)
	check(err)
	return i
}

func isNrChar(char string) bool {
	var numbers = [10]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	for _, number := range numbers {
		if char == number {
			return true
		}
	}
	return false
}

func check(error error) {
	if error != nil {
		panic(error)
	}
}
