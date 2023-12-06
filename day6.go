package adventOfCode23

import (
	"fmt"
	"strconv"
	"strings"
)

type Race struct {
	Duration int64
	Distance int64
}

func (r *Race) String() string {
	return fmt.Sprintf("Race lasts %v milliseconds and record distance is %v millimeters", r.Duration, r.Distance)
}

func Day6() {
	_, input := Input(6)

	time := strings.Fields(strings.Split(input[0], ":")[1])
	distance := strings.Fields(strings.Split(input[1], ":")[1])

	var races []Race
	for i := range time {
		d, _ := strconv.ParseInt(time[i], 10, 64)
		dd, _ := strconv.ParseInt(distance[i], 10, 64)
		races = append(races, Race{
			Duration: d,
			Distance: dd,
		})
	}

	result1 := int64(1)
	for _, v := range races {
		fmt.Println(v)
		result1 *= possible(v)
	}

	fmt.Println(result1)

	fmt.Println(possible(Race{
		Duration: 54946592,
		Distance: 302147610291404,
	}))

	x := Race{
		Duration: 54946592,
		Distance: 302147610291404,
	}

	l := LeftPossible(x, 0, x.Duration)
	r := RightPossible(x, 0, x.Duration)

	fmt.Printf("Left possible %v\n", l)
	fmt.Printf("Right possible %v\n", r)
	fmt.Println(r - l + 1)

}

func possible(r Race) int64 {
	var result int64
	for i := int64(0); i < r.Duration; i++ {
		if i*(r.Duration-i) >= r.Distance {
			result++
		}
	}
	return result
}

func LeftPossible(r Race, lo int64, hi int64) int64 {
	mid := lo + (hi-lo)/2

	if check(r, mid) && check(r, mid-1) {
		hi = mid
		return LeftPossible(r, lo, hi)
	}
	if !check(r, mid) && !check(r, mid+1) {
		lo = mid
		return LeftPossible(r, lo, hi)
	}
	if check(r, mid) && !check(r, mid-1) {
		return mid
	}

	return 0
}

func RightPossible(r Race, lo, hi int64) int64 {
	if hi-lo <= 1 {
		if check(r, hi) {
			return hi
		}
		return lo
	}

	mid := lo + (hi-lo)/2

	if check(r, mid) && check(r, mid-1) {
		return RightPossible(r, mid, hi)
	}
	if !check(r, mid) && !check(r, mid+1) {
		return RightPossible(r, lo, mid)
	}
	if check(r, mid) && !check(r, mid+1) {
		return mid
	}

	return 0
}

func check(r Race, l int64) bool {
	return l*(r.Duration-l) >= r.Distance
}
