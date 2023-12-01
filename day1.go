package adventOfCode23

import (
	"fmt"
	"strconv"
	"unicode"
)

func DayOnePart1() {
	err, input := Input(1)
	if err != nil {
		fmt.Printf("Input error: %v", err)
	}

	var result int64

	for _, v := range input {
		f := getFirstDigit(v)
		s := getSecondDigit(v)

		tmp, _ := strconv.Atoi(f + s)
		result += int64(tmp)
	}
	println(result)
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

func DayOnePart2() {
	err, input := Input(1)
	if err != nil {
		fmt.Printf("Input error: %v", err)
	}

	var result int64

	for _, v := range input {
		v = replaceLetters(v)
		f := getFirstDigit(v)
		s := getSecondDigit(v)

		tmp, _ := strconv.Atoi(f + s)
		result += int64(tmp)
	}
	println(result)
}

func replaceLetters(ss string) string {
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
