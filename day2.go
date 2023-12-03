package adventOfCode23

import (
	"fmt"
	"strconv"
	"strings"
)

func Day2() {
	limits := map[string]int{"red": 12, "green": 13, "blue": 14}

	_, input := Input(2)
	games := parseInput(input)

	part1, part2 := 0, 0
	for _, game := range games {
		if isGamePossible(game, limits) {
			part1 += game.id
		}
		part2 += calculateGamePower(game)
	}

	println("Day 2, Part 1:")
	fmt.Printf("Sum of IDs of games possible with limits (Red: %d, Green: %d, Blue: %d): %d\n", limits["red"], limits["green"], limits["blue"], part1)

	println("\nDay 2, Part 2:")
	fmt.Printf("Sum of the power of the minimum sets of cubes required for each game: %d\n", part2)
}

// Game represents a single game's data.
type Game struct {
	id     int
	rounds []map[string]int
}

// parseInput converts the raw input strings into structured Game data.
func parseInput(input []string) []Game {
	var games []Game
	for _, line := range input {
		parts := strings.Split(line, ":")
		id, _ := strconv.Atoi(strings.Fields(parts[0])[1])
		roundsData := strings.Split(parts[1], ";")

		var rounds []map[string]int
		for _, roundData := range roundsData {
			round := make(map[string]int)
			cubes := strings.Split(roundData, ",")
			for _, cube := range cubes {
				cube = strings.TrimSpace(cube)
				parts := strings.Fields(cube)
				count, _ := strconv.Atoi(parts[0])
				round[parts[1]] = count
			}
			rounds = append(rounds, round)
		}
		games = append(games, Game{id: id, rounds: rounds})
	}
	return games
}

// isGamePossible determines if a game could be played given the limits.
func isGamePossible(game Game, limits map[string]int) bool {
	for _, round := range game.rounds {
		for color, count := range round {
			if count > limits[color] {
				return false
			}
		}
	}
	return true
}

// calculateGamePower calculates the power of the minimum set of cubes needed for a game.
func calculateGamePower(game Game) int {
	minCubes := map[string]int{"red": 0, "green": 0, "blue": 0}
	for _, round := range game.rounds {
		for color, count := range round {
			if count > minCubes[color] {
				minCubes[color] = count
			}
		}
	}
	return minCubes["red"] * minCubes["green"] * minCubes["blue"]
}
