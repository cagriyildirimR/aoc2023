package adventOfCode23

import (
	"fmt"
	"strconv"
	"strings"
)

type Sequence []int64

func Day9() {
	_, input := Input(9)

	var OASIS []Sequence

	for _, v := range input {
		tmp := strings.Fields(v)
		var t Sequence
		for _, w := range tmp {
			parseInt, err := strconv.ParseInt(w, 10, 64)
			if err != nil {
				return
			}
			t = append(t, parseInt)
		}
		OASIS = append(OASIS, t)
	}

	var start, end int64
	for i := range OASIS {
		a, b := extrapolate(OASIS[i], false)
		start += a
		end += b
	}
	fmt.Println(start, end)

	//fmt.Println(extrapolate(Sequence{1, 3, 6, 10, 15, 21}))
	//fmt.Println(extrapolate(Sequence{10, 13, 16, 21, 30, 45}))
}

func extrapolate(s Sequence, p bool) (int64, int64) {
	var v Sequence
	var endVals Sequence
	var startVals Sequence
	var history []Sequence

	for !isZero(s) {
		history = append(history, append([]int64(nil), s...))
		endVals = append(endVals, s[len(s)-1])
		startVals = append(startVals, s[0])
		for i := 0; i < len(s)-1; i++ {
			v = append(v, s[i+1]-s[i])
		}
		s = v
		v = Sequence{}
	}

	var endX int64
	var startX int64
	extrapolatedValuesEnd := make([]int64, len(history))
	extrapolatedValuesStart := make([]int64, len(history))

	for i := range history {
		endX = 0
		startX = 0
		for j := len(endVals) - 1; j >= len(history)-1-i; j-- {
			endX += endVals[j]
			startX = startVals[j] - startX
		}
		extrapolatedValuesEnd[len(history)-1-i] = endX
		extrapolatedValuesStart[len(history)-1-i] = startX
	}

	// Print each sequence in history with its extrapolated value
	if p {
		for i, seq := range history {
			fmt.Printf("%v <- %v -> %v\n", extrapolatedValuesStart[i], seq, extrapolatedValuesEnd[i])
		}
	}

	return extrapolatedValuesStart[0], extrapolatedValuesEnd[0]
}

func isZero(s Sequence) bool {
	for _, v := range s {
		if v != 0 {
			return false
		}
	}
	return true
}
