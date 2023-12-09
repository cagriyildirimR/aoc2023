package adventOfCode23

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

var (
	types = map[string]int{
		"11111": 1,
		"1112":  2,
		"122":   3,
		"113":   4,
		"23":    5,
		"14":    6,
		"5":     7,
	}
	strengths = map[rune]int{
		'2': 1,
		'3': 2,
		'4': 3,
		'5': 4,
		'6': 5,
		'7': 6,
		'8': 7,
		'9': 8,
		'T': 9,
		'J': 10,
		'Q': 11,
		'K': 12,
		'A': 13,
	}
	JokerEnabled bool
)

type Card struct {
	Hand     string
	Strength int64
}

type CamelCards []Card

func (c Card) String() string {
	return fmt.Sprintf("Hand: %q, Strength: %v", c.Hand, c.Strength)
}

func (c CamelCards) Len() int {
	return len(c)
}

func (c CamelCards) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func (c CamelCards) Less(i, j int) bool {
	f := c[i].Hand
	s := c[j].Hand
	si, sj := runeFrequencyMap(f), runeFrequencyMap(s)

	if types[si] == types[sj] {
		for k, _ := range f {
			if f[k] != s[k] {
				return strengths[rune(f[k])] < strengths[rune(s[k])]
			}
		}
	}
	return types[si] < types[sj]
}

func runeFrequencyMap(s string) string {
	if JokerEnabled {
		strengths['J'] = 0
	}

	freqMap := make(map[rune]int)
	for _, v := range s {
		freqMap[v]++
	}

	if JokerEnabled && freqMap['J'] == 5 {
		return "5"
	}

	var js int
	if JokerEnabled && freqMap['J'] > 0 {
		js = freqMap['J']
		delete(freqMap, 'J')
	}

	freq := make([]int, 0, len(freqMap))
	for _, v := range freqMap {
		freq = append(freq, v)
	}

	sort.Ints(freq)
	if JokerEnabled {
		freq[len(freqMap)-1] += js
	}

	var result string
	for _, v := range freq {
		result += strconv.Itoa(v)
	}
	return result
}

func calculateResult(cs CamelCards) int64 {
	sort.Sort(cs)
	var result int64
	for i, v := range cs {
		result += int64(i+1) * v.Strength
	}
	return result
}

func Day7() {
	_, input := Input(7)

	var cards CamelCards
	for _, line := range input {
		parts := strings.Fields(line)
		strength, _ := strconv.ParseInt(parts[1], 10, 64)
		cards = append(cards, Card{Hand: parts[0], Strength: strength})
	}

	fmt.Println("Without Joker:", calculateResult(cards))

	JokerEnabled = true
	fmt.Println("With Joker:", calculateResult(cards))
}
