package adventOfCode23

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type ScratchCard struct {
	WinningNumbers []int
	ScratchNumbers []int
}

func Day4() {
	cards := parseInputDay4()
	Day4Part1(cards)
	Day4Part2(cards)
}

func parseInputDay4() []ScratchCard {
	_, input := Input(4)
	cards := make([]ScratchCard, len(input))

	for i, line := range input {
		line = strings.Split(line, ":")[1]
		parts := strings.Split(line, "|")
		cards[i].WinningNumbers = parseNumbers(strings.Fields(parts[0]))
		cards[i].ScratchNumbers = parseNumbers(strings.Fields(parts[1]))
	}

	return cards
}

func parseNumbers(fields []string) []int {
	nums := make([]int, len(fields))
	for i, field := range fields {
		num, err := strconv.Atoi(field)
		if err != nil {
			fmt.Printf("Error parsing number: %v\n", err)
			continue
		}
		nums[i] = num
	}
	return nums
}

func Day4Part1(cards []ScratchCard) {
	var result float64
	for _, card := range cards {
		matches := checkCommonValues(card)
		if matches > 0 {
			result += math.Pow(2, float64(matches-1))
		}
	}
	fmt.Printf("Total Points from Scratchcards (Part 1): %d\n", int64(result))
}

func Day4Part2(cards []ScratchCard) {
	results := make([]int, len(cards))
	for i := range results {
		results[i]++
	}

	totalCards := 0

	for i, card := range cards {
		for results[i] > 0 {
			totalCards++
			results[i]--
			matches := checkCommonValues(card)
			for j := 0; j < matches && (i+j+1) < len(cards); j++ {
				results[i+j+1]++
			}
		}
	}
	fmt.Printf("Total Number of Scratchcards After Processing (Part 2): %d\n", totalCards)
}

func checkCommonValues(card ScratchCard) int {
	matchCount := 0
	for _, w := range card.WinningNumbers {
		for _, s := range card.ScratchNumbers {
			if w == s {
				matchCount++
				break
			}
		}
	}
	return matchCount
}
