package adventOfCode23

import (
	"strconv"
	"strings"
	"unicode"
)

func Day3() {
	_, engineSchematic := Input(3)

	var engineSchematicByte [][]byte

	for _, v := range engineSchematic {
		engineSchematicByte = append(engineSchematicByte, []byte(v))
	}

	specialCharacters := allSpecialCharacters(engineSchematic)

	schematicHeight := len(engineSchematic)
	schematicWidth := len(engineSchematic[0])

	currentNumber := ""
	totalSum := int64(0)

	gears := make(map[*byte][]int)

	// Part 1
	for row := 0; row < schematicHeight; row++ {
		for col := 0; col < schematicWidth; col++ {
			char := rune(engineSchematic[row][col])
			if unicode.IsDigit(char) {
				currentNumber = currentNumber + string(engineSchematic[row][col])
			}

			isLastCol := col == schematicWidth-1
			if (!unicode.IsDigit(char) || isLastCol) && len(currentNumber) > 0 {
				v, _ := strconv.Atoi(currentNumber)
				if isNumberValid(currentNumber, row, col, engineSchematic, specialCharacters) {
					//fmt.Printf("valid number: %v\n", v)
					totalSum += int64(v)
				}
				l, b := isRatioNumber(currentNumber, row, col, engineSchematicByte)
				if l {
					gears[b] = append(gears[b], v)
				}
				currentNumber = ""
			}
		}
	}

	println(totalSum)

	// Part 2

	totalGearRatio := int64(0)
	for _, v := range gears {
		if len(v) == 2 {
			totalGearRatio += int64(v[0]) * int64(v[1])
		}
	}
	println(totalGearRatio)
}

func isNumberValid(number string, rowNum int, colNum int, schematic []string, specialChars string) bool {
	start := colNum - len(number)
	end := colNum - 1

	for row := rowNum - 1; row <= rowNum+1; row++ {
		if row < 0 || row >= len(schematic) {
			continue
		}
		for col := start - 1; col <= end+1; col++ {
			if col < 0 || col >= len(schematic[row]) {
				continue
			}
			if isSpecialCharacter(schematic[row][col], specialChars) {
				return true
			}
		}
	}
	return false
}

func isRatioNumber(number string, rowNum int, colNum int, schematic [][]byte) (bool, *byte) {
	start := colNum - len(number)
	end := colNum - 1

	for row := rowNum - 1; row <= rowNum+1; row++ {
		if row < 0 || row >= len(schematic) {
			continue
		}
		for col := start - 1; col <= end+1; col++ {
			if col < 0 || col >= len(schematic[row]) {
				continue
			}
			if isSpecialCharacter(schematic[row][col], "*") {
				return true, &schematic[row][col]
			}
		}
	}
	return false, nil
}

func isSpecialCharacter(character byte, specialCharsList string) bool {
	return strings.Contains(specialCharsList, string(character))
}

func allSpecialCharacters(schematic []string) string {
	uniqueChars := ""
	for _, line := range schematic {
		for _, char := range line {
			if !strings.Contains(uniqueChars, string(char)) && !unicode.IsDigit(char) && char != '.' {
				uniqueChars = uniqueChars + string(char)
			}
		}
	}
	return uniqueChars
}
