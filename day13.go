package adventOfCode23

import "fmt"

type Pattern []string

func (p Pattern) String() string {
	var result string
	for _, line := range p {
		result += line + "\n"
	}
	return result
}

func (p Pattern) Summary(s int) int64 {
	hLine := findReflection(p, s, false)
	vLine := findReflection(p, s, true)

	return int64(hLine)*100 + int64(vLine)
}

func findReflection(p Pattern, maxDifferences int, isRow bool) int {
	x := len(p[0])
	y := len(p)
	if !isRow {
		x = len(p)
		y = len(p[0])
	}

	for midpoint := 0; midpoint < x-1; midpoint++ {
		totalDifferences := 0

		for index := 0; index < y; index++ {
			totalDifferences += countDifferences(p, index, midpoint, x, isRow)
			if totalDifferences > maxDifferences {
				break
			}
		}

		if totalDifferences == maxDifferences {
			return midpoint + 1
		}
	}

	return 0
}

func countDifferences(p Pattern, rowIndexOrColumnIndex, midpoint, widthOrHeight int, isRow bool) int {
	differences := 0
	for offset := 0; ; offset++ {
		startIndex := midpoint - offset
		endIndex := midpoint + offset + 1

		if startIndex < 0 || endIndex >= widthOrHeight {
			break
		}

		if isRow {
			if p[rowIndexOrColumnIndex][startIndex] != p[rowIndexOrColumnIndex][endIndex] {
				differences++
			}
		} else {
			if p[startIndex][rowIndexOrColumnIndex] != p[endIndex][rowIndexOrColumnIndex] {
				differences++
			}
		}
	}
	return differences
}

func Day13() {
	_, input := Input(13)

	var patterns []Pattern
	var currentPattern Pattern
	for _, line := range input {
		if line == "" {
			patterns = append(patterns, currentPattern)
			currentPattern = []string{}
		} else {
			currentPattern = append(currentPattern, line)
		}
	}
	patterns = append(patterns, currentPattern)

	var result, resultWithSmudgeFixed int64
	for _, pattern := range patterns {
		result += pattern.Summary(0)
		resultWithSmudgeFixed += pattern.Summary(1)
	}
	fmt.Printf("Part 1 Result: %d\n", result)
	fmt.Printf("Part 2 Result: %d\n", resultWithSmudgeFixed)
}
