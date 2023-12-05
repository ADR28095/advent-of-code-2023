package main

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	var fileName string = "input.txt"
	var result int = 0

	lines := readFile(fileName)

	for _, line := range lines {
		temp := parseToNr(getCalibrationValue(line))
		result += temp
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

	for index, char := range line {
		c := string(char)
		if isNrChar(c) {
			if firstNr == "" {
				firstNr = c
			}
			lastNr = c
		} else {
			parsedChar := checkIfNrAsText(line[index:])
			if isNrChar(parsedChar) {
				if firstNr == "" {
					firstNr = parsedChar
				}
				lastNr = parsedChar
			}
		}
	}
	return firstNr + lastNr
}

func checkIfNrAsText(text string) string {
	increment := 5
	if len(text) < increment {
		increment = len(text)
		return parseTextToNumber(text)[0:1]
	}

	targetString := substring(text, 0, increment)
	result := strings.Replace(text, targetString, parseTextToNumber(targetString), 1)
	return result[0:1]
}

func parseToNr(char string) int {
	i, err := strconv.Atoi(char)
	check(err)
	return i
}

func parseTextToNumber(text string) string {
	var result string
	if len(text) >= 3 {
		result = parse(text[0:3])
		if len(result) == 3 && len(text) >= 4 {
			result = parse(text[0:4])
			if len(result) == 4 && len(text) >= 5 {
				result = parse(text[0:5])
			}
		}
	} else {
		return text
	}

	return result
}

func parse(text string) string {
	var numbersAsText = [9]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	for index, number := range numbersAsText {
		text = regexp.MustCompile(number).ReplaceAllString(text, strconv.Itoa(index+1))
	}
	return text
}

func substring(str string, start int, end int) string {
	return strings.TrimSpace(str[start:end])
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
