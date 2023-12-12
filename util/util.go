package util

import (
	"bufio"
	"os"
	"strconv"
)

func ReadFile(fileName string) []string {
	var lines []string
	data, err := os.Open(fileName)
	Check(err)
	fileScanner := bufio.NewScanner(data)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}
	return lines
}

func ParseToNr(text string) int {
	i, err := strconv.Atoi(text)
	Check(err)
	return i
}

func Check(error error) {
	if error != nil {
		panic(error)
	}
}
