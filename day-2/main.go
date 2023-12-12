package main

import (
	"advent-of-code-2023/util"
	"strings"
)

// 32767 34310

func main() {
	var fileName = "input.txt"
	lines := util.ReadFile(fileName)
	var result int = 0
	var totalPower = 0

	for _, line := range lines {
		gameLine := strings.Split(line, ":")
		gameId := strings.Split(gameLine[0], " ")[1]
		isPossible, power := isGamePossible(gameLine[1])
		totalPower += power
		if isPossible {
			result += util.ParseToNr(gameId)
		}
	}

	println("Result: ", result)
	println("Total power: ", totalPower)
}

func isGamePossible(gameData string) (bool, int) {
	maxRed := 1
	maxBlue := 1
	maxGreen := 1
	isPossible := true

	roundResult := strings.Split(gameData, ";")
	for _, round := range roundResult {
		roundPossibilityState, values := isRoundPossible(round)
		if maxRed < values[0] {
			maxRed = values[0]
		}
		if maxBlue < values[1] {
			maxBlue = values[1]
		}
		if maxGreen < values[2] {
			maxGreen = values[2]
		}
		if !roundPossibilityState {
			isPossible = false
		}
	}
	return isPossible, calcPower(maxRed, maxBlue, maxGreen)
}

func calcPower(red int, blue int, green int) int {
	return red * blue * green
}

func isRoundPossible(round string) (bool, [3]int) {
	isPossible := false
	selections := strings.Split(round, ",")
	values := [3]int{0, 0, 0}
	for _, selection := range selections {
		selectionPossibilityState, color, amount := isSelectionPossible(selection)

		if color == "red" {
			values[0] += amount
		} else if color == "blue" {
			values[1] += amount
		} else if color == "green" {
			values[2] += amount
		}

		if !selectionPossibilityState {
			isPossible = false
		}
	}
	return isPossible, values
}

func isSelectionPossible(selection string) (bool, string, int) {
	maxRed := 12
	maxGreen := 13
	maxBlue := 14
	data := strings.Split(strings.TrimSpace(selection), " ")

	amount := util.ParseToNr(data[0])
	redCheck := data[1] == "red" && amount > maxRed
	greenCheck := data[1] == "green" && amount > maxGreen
	blueCheck := data[1] == "blue" && amount > maxBlue
	if redCheck || greenCheck || blueCheck {
		return false, data[1], amount
	}
	return true, data[1], amount
}
