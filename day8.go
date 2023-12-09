package adventOfCode23

import (
	"fmt"
	"strings"
)

// Node represents a node in the network.
type Node struct {
	Head, Left, Right string
}

// Day8 executes the main logic for Day 8 challenge.
func Day8() {
	_, input := Input(8)
	instructions := input[0]
	networkMap := buildNetworkMap(input[2:])

	// Part One
	stepsToZZZ, _ := navigateFromNode("AAA", instructions, networkMap)
	fmt.Printf("Steps required to reach ZZZ (Part One): %d\n", stepsToZZZ)

	// Part Two
	startNodes := findStartingNodes(networkMap)
	results := navigateNetwork(instructions, networkMap, startNodes)
	fmt.Printf("Combined steps to reach nodes ending with Z (Part Two): %d\n", lcmOfList(results))
}

// buildNetworkMap constructs a map of nodes from the input slice of strings.
func buildNetworkMap(input []string) map[string]Node {
	networkMap := make(map[string]Node)
	for _, line := range input {
		parts := strings.Split(line, " = ")
		lr := strings.Split(strings.Trim(parts[1], "()"), ", ")
		networkMap[parts[0]] = Node{Head: parts[0], Left: lr[0], Right: lr[1]}
	}
	return networkMap
}

// findStartingNodes identifies all nodes ending with 'A' to start the navigation.
func findStartingNodes(networkMap map[string]Node) []string {
	var startNodes []string
	for _, node := range networkMap {
		if node.Head[len(node.Head)-1] == 'A' {
			startNodes = append(startNodes, node.Head)
		}
	}
	return startNodes
}

// navigateNetwork processes each starting node and calculates the number of steps to reach the end.
func navigateNetwork(ins string, networkMap map[string]Node, startNodes []string) []int64 {
	var results []int64
	for _, startNode := range startNodes {
		steps, _ := navigateFromNode(startNode, ins, networkMap)
		results = append(results, steps)
	}
	return results
}

// navigateFromNode navigates the network from a given node according to the instructions.
func navigateFromNode(node string, ins string, networkMap map[string]Node) (int64, bool) {
	var steps int64
	for {
		for _, direction := range ins {
			currentNode := networkMap[node]
			if direction == 'L' {
				node = currentNode.Left
			} else if direction == 'R' {
				node = currentNode.Right
			}
			steps++
			if node[len(node)-1] == 'Z' {
				return steps, true
			}
		}
	}
}

// gcd calculates the greatest common divisor of two integers.
func gcd(a, b int64) int64 {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// lcm calculates the least common multiple of two integers.
func lcm(a, b int64) int64 {
	return a * b / gcd(a, b)
}

// lcmOfList calculates the least common multiple of a list of integers.
func lcmOfList(arr []int64) int64 {
	lcmValue := arr[0]
	for _, num := range arr[1:] {
		lcmValue = lcm(lcmValue, num)
	}
	return lcmValue
}
