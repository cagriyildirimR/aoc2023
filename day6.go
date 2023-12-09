package adventOfCode23

import (
	"fmt"
	"strconv"
	"strings"
)

type Race struct {
	Duration, Distance int64
}

func (r Race) String() string {
	return fmt.Sprintf("Race lasts %v milliseconds and record distance is %v millimeters", r.Duration, r.Distance)
}

func Day6() {
	races, extendedRace := parseInputDay6()
	totalWays := int64(1)

	for _, race := range races {
		totalWays *= calculateWinningWays(race)
	}
	fmt.Println("Total winning ways for all races:", totalWays)
	fmt.Println("Winning ways for the extended race:", calculateWinningWays(extendedRace))
}

func parseInputDay6() ([]Race, Race) {
	_, input := Input(6)
	times := parseLine(input[0])
	distances := parseLine(input[1])
	extendedTime := parseLine(strings.ReplaceAll(input[0], " ", ""))[0]
	extendedDistance := parseLine(strings.ReplaceAll(input[1], " ", ""))[0]

	extendedRace := Race{
		Duration: extendedTime,
		Distance: extendedDistance,
	}

	var races []Race
	for i, time := range times {
		races = append(races, Race{
			Duration: time,
			Distance: distances[i],
		})
	}
	return races, extendedRace
}

func parseLine(line string) []int64 {
	fields := strings.Fields(strings.Split(line, ":")[1])
	var numbers []int64
	for _, f := range fields {
		number, _ := strconv.ParseInt(f, 10, 64)
		numbers = append(numbers, number)
	}
	return numbers
}

func calculateWinningWays(race Race) int64 {
	left := findLeftMostPossible(race, 0, race.Duration)
	right := findRightMostPossible(race, 0, race.Duration)
	return right - left + 1
}

func findLeftMostPossible(race Race, low, high int64) int64 {
	for low < high {
		mid := low + (high-low)/2
		if isValid(race, mid) {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return low
}

func findRightMostPossible(race Race, low, high int64) int64 {
	for low < high {
		mid := low + (high-low)/2
		if isValid(race, mid+1) {
			low = mid + 1
		} else {
			high = mid
		}
	}
	return low
}

func isValid(race Race, time int64) bool {
	return time*(race.Duration-time) >= race.Distance
}
