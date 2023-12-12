package main

import "advent-of-code-2023/util"

func main() {
	var fileName = "test-input.txt"
	lines := util.ReadFile(fileName)
	result := 0

	for index, _ := range lines {
		above := ""
		bellow := ""

		if index > 0 {
			above = lines[index-1]
		}

		if len(lines) > index+1 {
			bellow = lines[index+1]
		}

		result += extractPartNumbers(lines[index], above, bellow)
	}

	println("Result: ", result)
}

func extractPartNumbers(current string, above string, below string) int {
	result := 0
	numberAsString := ""

	for _, char := range current {
		//isPartNr := false
		if util.IsNrChar(string(char)) {
			numberAsString += string(char)
		} else {
			if numberAsString != "" {
				println(numberAsString)
				numberAsString = ""
			}
		}
	}
	println(numberAsString)
	return result
}
