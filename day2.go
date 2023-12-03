package adventOfCode23

import (
	"strconv"
	"strings"
)

func DayTwoPart1() {

	limits := make(map[string]int)
	limits["red"] = 12
	limits["green"] = 13
	limits["blue"] = 14

	_, input := Input(2)

	result := 0

	for _, v := range input {
		result += checkPossibleGames(v, limits)
	}
	println(result)
}

func DayTwoPart2() {
	_, input := Input(2)

	result := 0

	for _, v := range input {
		result += calculatePower(v)
	}
	println(result)
}

func checkPossibleGames(ss string, limits map[string]int) int {

	tmp := strings.Split(ss, ":")
	id, _ := strconv.Atoi(strings.Split(tmp[0], " ")[1])
	games := strings.Split(tmp[1], ";")

	for _, g := range games {
		cubes := strings.Split(g, ",")
		for _, c := range cubes {
			c := strings.Trim(c, " ")
			v := strings.Split(c, " ")

			i, _ := strconv.Atoi(v[0])
			if limits[v[1]] < i {
				return 0
			}
		}
	}
	return id
}

func calculatePower(ss string) int {

	m := make(map[string]int)
	m["red"] = 0
	m["green"] = 0
	m["blue"] = 0

	tmp := strings.Split(ss, ":")
	games := strings.Split(tmp[1], ";")

	for _, g := range games {
		cubes := strings.Split(g, ",")
		for _, c := range cubes {
			c := strings.Trim(c, " ")
			v := strings.Split(c, " ")

			i, _ := strconv.Atoi(v[0])
			if m[v[1]] < i {
				m[v[1]] = i
			}
		}
	}
	return m["red"] * m["green"] * m["blue"]
}
