package adventOfCode23

import (
	"fmt"
	"strconv"
	"unicode"
)

func Day1() {
	err, calibrationInput := Input(1)
	if err != nil {
		fmt.Printf("Calibration data read error: %v", err)
	}

	var sumOriginalValues int64
	var sumAdjustedValues int64

	for _, calibrationLine := range calibrationInput {
		firstDigitOriginal := getFirstDigit(calibrationLine)
		secondDigitOriginal := getSecondDigit(calibrationLine)

		combinedDigits, _ := strconv.Atoi(firstDigitOriginal + secondDigitOriginal)
		sumOriginalValues += int64(combinedDigits)

		adjustedLine := replaceSpelledNumbersWithDigits(calibrationLine)

		firstDigitAdjusted := getFirstDigit(adjustedLine)
		secondDigitAdjusted := getSecondDigit(adjustedLine)

		combinedAdjustedDigits, _ := strconv.Atoi(firstDigitAdjusted + secondDigitAdjusted)
		sumAdjustedValues += int64(combinedAdjustedDigits)
	}

	fmt.Printf("Part 1 Solution: Sum of Original Calibration Values is %v\n", sumOriginalValues)
	fmt.Printf("Part 2 Solution: Sum of Adjusted Calibration Values is %v\n", sumAdjustedValues)
}

func getFirstDigit(ss string) string {
	for _, v := range ss {
		if unicode.IsDigit(v) {
			return string(v)
		}
	}
	return "error"
}

func getSecondDigit(ss string) string {
	for i := len(ss) - 1; i >= 0; i-- {
		v := rune(ss[i])
		if unicode.IsDigit(v) {
			return string(v)
		}
	}
	return "error"
}

// replaceSpelledNumbersWithDigits takes a string `ss` and replaces occurrences of spelled-out numbers (one, two, three, ..., nine)
// with their corresponding single-digit numerals (1, 2, 3, ..., 9). The function scans the string from left to right,
// and upon encountering a spelled-out number, it replaces only the first character of the spelled-out number with the
// corresponding digit. The rest of the characters in the spelled-out number are left unchanged.
//
// For example:
// - "oneight" becomes "1eight" (replaces 'one' with '1').
// - "twofive" becomes "2five" (replaces 'two' with '2').
// - "threenine" becomes "3nine" (replaces 'three' with '3').
//
// The function is designed to work in a specific context where only the first and last digits in the resultant string
// are considered significant for further processing. This behavior is crucial for the function's application in certain
// scenarios, such as solving specific coding challenges or puzzles.
//
// Parameters:
//
//	ss (string): The input string to process.
//
// Returns:
//
//	string: A new string with spelled-out numbers replaced by their corresponding single-digit numerals.
func replaceSpelledNumbersWithDigits(ss string) string {
	var tmp = []byte(ss)
	for i, s := range []byte(ss) {
		switch s {
		case 'o':
			if i+2 < len(ss) {
				if ss[i+1] == byte('n') && ss[i+2] == byte('e') {
					tmp[i] = byte('1')
				}
			}
		case 't':
			if i+2 < len(ss) {
				if ss[i+1] == byte('w') && ss[i+2] == byte('o') {
					tmp[i] = byte('2')
				}
			}
			if i+4 < len(ss) {
				if ss[i+1] == byte('h') && ss[i+2] == byte('r') && ss[i+3] == byte('e') && ss[i+4] == byte('e') {
					tmp[i] = byte('3')
				}
			}
		case 'f':
			if i+3 < len(ss) {
				if ss[i+1] == byte('o') && ss[i+2] == byte('u') && ss[i+3] == byte('r') {
					tmp[i] = byte('4')
				}
				if ss[i+1] == byte('i') && ss[i+2] == byte('v') && ss[i+3] == byte('e') {
					tmp[i] = byte('5')
				}
			}
		case 's':
			if i+2 < len(ss) {
				if ss[i+1] == byte('i') && ss[i+2] == byte('x') {
					tmp[i] = byte('6')
				}
			}
			if i+4 < len(ss) {
				if ss[i+1] == byte('e') && ss[i+2] == byte('v') && ss[i+3] == byte('e') && ss[i+4] == byte('n') {
					tmp[i] = byte('7')
				}
			}
		case 'e':
			if i+4 < len(ss) {
				if ss[i+1] == byte('i') && ss[i+2] == byte('g') && ss[i+3] == byte('h') && ss[i+4] == byte('t') {
					tmp[i] = byte('8')
				}
			}
		case 'n':
			if i+3 < len(ss) {
				if ss[i+1] == byte('i') && ss[i+2] == byte('n') && ss[i+3] == byte('e') {
					tmp[i] = byte('9')
				}
			}
		}
	}
	return string(tmp)
}

// these functions don't work, but they are useful to know in similar situations
//func replaceLettersX(s string) string {
//	s = strings.ReplaceAll(s, "one", "1ne")
//	s = strings.ReplaceAll(s, "two", "2wo")
//	s = strings.ReplaceAll(s, "three", "3hree")
//	s = strings.ReplaceAll(s, "four", "4our")
//	s = strings.ReplaceAll(s, "five", "5ive")
//	s = strings.ReplaceAll(s, "six", "6ix")
//	s = strings.ReplaceAll(s, "seven", "7even")
//	s = strings.ReplaceAll(s, "eight", "8ight")
//	s = strings.ReplaceAll(s, "nine", "9ine")
//	return s
//}

//func replaceLetters(s string) string {
//	replacer := strings.NewReplacer(
//		"one", "1ne",
//		"two", "2wo",
//		"three", "3hree",
//		"four", "4our",
//		"five", "5ive",
//		"six", "6ix",
//		"seven", "7even",
//		"eight", "8eight",
//		"nine", "9ine",
//	)
//	return replacer.Replace(s)
//}

//func replaceLetters(s string) string {
//	numberWords := map[string]string{
//		"one":   "1",
//		"two":   "2",
//		"three": "3",
//		"four":  "4",
//		"five":  "5",
//		"six":   "6",
//		"seven": "7",
//		"eight": "8",
//		"nine":  "9",
//	}
//
//	var result strings.Builder
//	for len(s) > 0 {
//		matchFound := false
//		for word, digit := range numberWords {
//			if strings.HasPrefix(s, word) {
//				result.WriteString(digit)
//				s = s[len(word):]
//				matchFound = true
//				break
//			}
//		}
//		if !matchFound {
//			result.WriteByte(s[0])
//			s = s[1:]
//		}
//	}
//
//	return result.String()
//}
