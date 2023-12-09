package adventOfCode23

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"unicode"
)

const DayOne = 1

func Day1() {
	input := getInput()

	var sumOriginalValues int64
	var sumAdjustedValues int64

	for _, calibrationLine := range input {
		err, digitFromStart := getDigitFromStart(calibrationLine)
		if err != nil {
			log.Printf("encountered error while getting digits %v", err)
		}

		err, digitFromEnd := getDigitFromEnd(calibrationLine)
		if err != nil {
			log.Printf("encountered error while getting digits %v", err)
		}

		combinedDigits, err := strconv.Atoi(digitFromStart + digitFromEnd)
		if err != nil {
			log.Printf("cannot convert to digit, error: %v", err)
		}
		sumOriginalValues += int64(combinedDigits)

		adjustedLine := replaceSpelledNumbersWithDigits(calibrationLine)

		err, digitFromStartAdjusted := getDigitFromStart(adjustedLine)
		if err != nil {
			log.Printf("encountered error while getting digits %v", err)
		}

		err, digitFromEndAdjusted := getDigitFromEnd(adjustedLine)
		if err != nil {
			log.Printf("encountered error while getting digits %v", err)
		}

		combinedAdjustedDigits, err := strconv.Atoi(digitFromStartAdjusted + digitFromEndAdjusted)
		if err != nil {
			log.Printf("cannot convert to digit, error: %v", err)
		}
		sumAdjustedValues += int64(combinedAdjustedDigits)
	}

	fmt.Printf("Part 1 Solution: Sum of Original Calibration Values is %v\n", sumOriginalValues)
	fmt.Printf("Part 2 Solution: Sum of Adjusted Calibration Values is %v\n", sumAdjustedValues)
}

// getInput retrieves the calibration data for Day One from AdventOfCode.
// Returns a slice of strings, each representing a line of calibration data.
// If an error occurs during data retrieval, the error is printed and an empty slice is returned.
func getInput() []string {
	err, input := Input(DayOne)
	if err != nil {
		fmt.Printf("Calibration data read error: %v", err)
	}
	return input
}

// getDigitFromStart scans a string from left to right and returns the first digit it encounters.
// If no digit is found, an error is returned.
//
// Input: ss - the string to be scanned for a digit.
//
// Output: The first digit found as a string, or an error if no digit is present.
func getDigitFromStart(ss string) (error, string) {
	for _, s := range ss {
		if unicode.IsDigit(s) {
			return nil, string(s)
		}
	}
	return errors.New("error getting digit from start"), ""
}

// getDigitFromEnd scans a string from right to left and returns the first digit it encounters.
// If no digit is found, an error is returned.
//
// Input: ss - the string to be scanned for a digit.
//
// return: The first digit found as a string, or an error if no digit is present.
func getDigitFromEnd(ss string) (error, string) {
	for i := len(ss) - 1; i >= 0; i-- {
		v := rune(ss[i])
		if unicode.IsDigit(v) {
			return nil, string(v)
		}
	}
	return errors.New("error getting digit from end"), ""
}

// replaceSpelledNumbersWithDigits processes a string and replaces all spelled-out numbers (one, two, etc.)
// with their corresponding digit characters (1, 2, etc.), ensuring digits in the output are numerically correct.
// This aids in extracting consecutive digits from strings where numbers are spelled out.
//
// Example: "eightwo" is transformed to "ei8ht2o".
//
// Input: s - the string with spelled-out numbers.
//
// Output: The string with spelled-out numbers replaced by their corresponding digits.
func replaceSpelledNumbersWithDigits(s string) string {
	s = strings.ReplaceAll(s, "one", "o1e")
	s = strings.ReplaceAll(s, "two", "t2o")
	s = strings.ReplaceAll(s, "three", "t3ree")
	s = strings.ReplaceAll(s, "four", "f4ur")
	s = strings.ReplaceAll(s, "five", "f5ve")
	s = strings.ReplaceAll(s, "six", "s6x")
	s = strings.ReplaceAll(s, "seven", "se7en")
	s = strings.ReplaceAll(s, "eight", "ei8ht")
	s = strings.ReplaceAll(s, "nine", "ni9e")
	return s
}
