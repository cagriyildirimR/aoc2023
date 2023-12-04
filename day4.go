package adventOfCode23

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func Day4() {
	_, inp := Input(4)
	var w [218][10]int
	var s [218][25]int
	for i, v := range inp {
		v = strings.Split(v, ":")[1]
		tmp := strings.Split(v, "|")

		for j, f := range strings.Fields(tmp[0]) {
			va, _ := strconv.Atoi(f)
			w[i][j] = va
		}

		for j, f := range strings.Fields(tmp[1]) {
			va, _ := strconv.Atoi(f)
			s[i][j] = va
		}
	}

	var result1 float64
	for i, v := range w {
		fmt.Printf("Game %v\n", i)
		tmp := checkCommonValues(v, s[i])
		if tmp > 0 {
			result1 += math.Pow(2, tmp-1)
		}
	}

	println(int64(result1))
}

func Day4Part2() {
	_, inp := Input(4)
	var w [218][10]int
	var s [218][25]int
	var r [218]int
	cards := 0

	for i := range r {
		r[i]++
	}

	for i, v := range inp {
		v = strings.Split(v, ":")[1]
		tmp := strings.Split(v, "|")

		for j, f := range strings.Fields(tmp[0]) {
			va, _ := strconv.Atoi(f)
			w[i][j] = va
		}

		for j, f := range strings.Fields(tmp[1]) {
			va, _ := strconv.Atoi(f)
			s[i][j] = va
		}
	}

	for i, v := range w {
		for r[i] > 0 {
			cards++
			r[i]--
			tmp := checkCommonValues(v, s[i])
			ind := i + 1
			for tmp > 0 {
				if ind < 218 {
					r[ind]++
					ind++
				}
				tmp--
			}
		}
	}
	println(cards)
}

func checkCommonValues(a [10]int, b [25]int) float64 {
	var result int
	for _, v := range a {
		for _, w := range b {
			if v == w {
				result++
				break
			}
		}
	}
	return float64(result)
}

func set(a [25]int) [25]int {
	neg := -1

	for i, v := range a {
		for j, w := range a {
			if i != j && v == w {
				a[i] = neg
				neg--
			}
		}
	}
	return a
}
