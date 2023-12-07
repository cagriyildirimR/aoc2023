package adventOfCode23

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

var types map[string]int
var strengths map[rune]int
var strengthsJoker map[rune]int

func init() {
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
	strengthsJoker = map[rune]int{
		'J': 0,
		'2': 1,
		'3': 2,
		'4': 3,
		'5': 4,
		'6': 5,
		'7': 6,
		'8': 7,
		'9': 8,
		'T': 9,
		'Q': 11,
		'K': 12,
		'A': 13,
	}
}

type Card struct {
	Hand     string
	Strength int64
}

func (c Card) String() string {
	return fmt.Sprintf("Hand: %q, Strength: %v", c.Hand, c.Strength)
}

type CamelCards []Card

func (c CamelCards) Len() int {
	return len(c)
}

func (c CamelCards) Less(i, j int) bool {
	f := c[i].Hand
	s := c[j].Hand
	si, sj := runeFrequencyMapJoker(f), runeFrequencyMapJoker(s)

	if types[si] == types[sj] {
		for k, _ := range f {
			if f[k] != s[k] {
				return strengthsJoker[rune(f[k])] < strengthsJoker[rune(s[k])]
			}
		}
	}
	return types[si] < types[sj]
}

func runeFrequencyMap(s string) string {
	freqMap := make(map[rune]int)
	for _, v := range s {
		freqMap[v]++
	}

	freq := make([]int, 0, len(freqMap))
	for _, v := range freqMap {
		freq = append(freq, v)
	}

	sort.Ints(freq)

	var result string
	for _, v := range freq {
		result += strconv.Itoa(v)
	}
	return result
}

func runeFrequencyMapJoker(s string) string {
	freqMap := make(map[rune]int)
	for _, v := range s {
		freqMap[v]++
	}

	if freqMap['J'] == 5 {
		return "5"
	}

	var js int
	if freqMap['J'] > 0 {
		js = freqMap['J']
		delete(freqMap, 'J')
	}

	var freq []int
	for _, v := range freqMap {
		freq = append(freq, v)
	}

	sort.Ints(freq)
	freq[len(freqMap)-1] += js

	var result string
	for _, v := range freq {
		result += strconv.Itoa(v)
	}
	return result
}

func (c CamelCards) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func Day7() {
	_, input := Input(7)

	var cs CamelCards
	for _, v := range input {
		w := strings.Fields(v)
		y, _ := strconv.ParseInt(w[1], 10, 64)
		cs = append(cs, Card{
			Hand:     w[0],
			Strength: y,
		})
	}

	//cs = CamelCards{
	//	Card{
	//		Hand:     "32T3K",
	//		Strength: 765,
	//	},
	//	Card{
	//		Hand:     "T55J5",
	//		Strength: 684,
	//	},
	//	Card{
	//		Hand:     "KK677",
	//		Strength: 28,
	//	},
	//	Card{
	//		Hand:     "KTJJT",
	//		Strength: 220,
	//	},
	//	Card{
	//		Hand:     "QQQJA",
	//		Strength: 483,
	//	},
	//}

	sort.Sort(cs)
	var result int64
	for i, v := range cs {
		fmt.Println(v)
		result += int64(i+1) * v.Strength
	}

	fmt.Println(result)
}
