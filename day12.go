package adventOfCode23

import (
	"fmt"
	"strconv"
	"strings"
)

var springArrangementsCache = make(map[string]int)

func Day12() {
	_, springConditions := Input(12)

	totalArrangements := calculateTotalArrangements(springConditions, false)
	fmt.Printf("Total arrangements for original conditions: %d\n", totalArrangements)

	totalArrangementsUnfolded := calculateTotalArrangements(springConditions, true)
	fmt.Printf("Total arrangements for unfolded conditions: %d\n", totalArrangementsUnfolded)
}

func calculateTotalArrangements(springConditions []string, unfold bool) int {
	total := 0
	for _, condition := range springConditions {
		parts := strings.Fields(condition)
		springs := parts[0]
		conditions := parts[1]

		if unfold {
			springs = strings.Join([]string{springs, springs, springs, springs, springs}, "?")
			conditions = strings.Join([]string{conditions, conditions, conditions, conditions, conditions}, ",")
		}

		total += calculateArrangements(springs, parseConditions(conditions))
	}
	return total
}

func evaluateBaseCase(damagedGroups []int) int {
	if len(damagedGroups) == 0 {
		return 1
	}
	return 0
}

func calculateArrangements(springStates string, damagedGroups []int) int {
	cacheKey := buildCacheKey(springStates, damagedGroups)

	if result, found := springArrangementsCache[cacheKey]; found {
		return result
	}

	if len(springStates) == 0 {
		return evaluateBaseCase(damagedGroups)
	}

	firstSpring := springStates[0]
	switch firstSpring {
	case '?':
		return calculateUnknownSpring(springStates, damagedGroups)
	case '.':
		return calculateOperationalSpring(springStates, damagedGroups, cacheKey)
	case '#':
		return calculateDamagedSpring(springStates, damagedGroups, cacheKey)
	}

	return 0
}

func buildCacheKey(springStates string, damagedGroups []int) string {
	cacheKey := springStates
	for _, size := range damagedGroups {
		cacheKey += strconv.Itoa(size) + ","
	}
	return cacheKey
}

func calculateUnknownSpring(springStates string, damagedGroups []int) int {
	withOperational := calculateArrangements(replaceFirstUnknown(springStates, "."), damagedGroups)
	withDamaged := calculateArrangements(replaceFirstUnknown(springStates, "#"), damagedGroups)
	return withOperational + withDamaged
}

func calculateOperationalSpring(springStates string, damagedGroups []int, cacheKey string) int {
	result := calculateArrangements(springStates[1:], damagedGroups)
	springArrangementsCache[cacheKey] = result
	return result
}

func calculateDamagedSpring(springStates string, damagedGroups []int, cacheKey string) int {
	return handleDamagedSpring(springStates, damagedGroups, cacheKey)
}

func replaceFirstUnknown(springStates, replacement string) string {
	return strings.Replace(springStates, "?", replacement, 1)
}

func handleDamagedSpring(springStates string, damagedGroups []int, cacheKey string) int {
	switch {
	case len(damagedGroups) == 0 || len(springStates) < damagedGroups[0]:
		springArrangementsCache[cacheKey] = 0
		return 0

	case strings.Contains(springStates[:damagedGroups[0]], "."):
		springArrangementsCache[cacheKey] = 0
		return 0

	case len(damagedGroups) > 1 && (len(springStates) <= damagedGroups[0] || springStates[damagedGroups[0]] == '#'):
		springArrangementsCache[cacheKey] = 0
		return 0
	}

	var result int
	if len(damagedGroups) > 1 {
		result = calculateArrangements(springStates[damagedGroups[0]+1:], damagedGroups[1:])
	} else {
		result = calculateArrangements(springStates[damagedGroups[0]:], []int{})
	}

	springArrangementsCache[cacheKey] = result
	return result
}

func parseConditions(conditionString string) []int {
	var conditions []int
	for _, group := range strings.Split(conditionString, ",") {
		size, _ := strconv.Atoi(group)
		conditions = append(conditions, size)
	}
	return conditions
}
