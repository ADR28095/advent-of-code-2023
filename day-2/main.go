package main

import (
	"advent-of-code-2023/util"
	"strings"
)

func main() {
	var fileName = "test-input.txt"
	lines := util.ReadFile(fileName)
	var result int = 0

	for _, line := range lines {
		parseGame(line)
	}

	println("Result: ", result)
}

func parseGame(gameLine string) {
	result := strings.Split(gameLine, ":")
	gameId := strings.Split(result[0], " ")[1]
	println(gameId)
	roundResult := strings.Split(result[1], ";")
	for _, round := range roundResult {
		if !isRoundPossible(round) {
			return
		}
	}
}

func isRoundPossible(round string) bool {
	selections := strings.Split(round, ",")
	for _, selection := range selections {
		if !isSelectionPossible(selection) {
			return false
		}
	}
	return true
}

func isSelectionPossible(selection string) bool {
	maxRed := 12
	maxGreen := 13
	maxBlue := 14
	data := strings.Split(selection, " ")
	amount := util.ParseToNr(data[0])
	if data[1] == "red" && amount > maxRed {
		return false
	} else if data[1] == "green" && amount > maxGreen {
		return false
	} else if data[1] == "blue" && amount > maxBlue {
		return false
	}
	return true
}
