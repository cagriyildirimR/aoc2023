package adventOfCode23

import (
	"fmt"
	"strings"
)

func Day8() {

	_, input := Input(8)

	ins := input[0]
	var network []Node

	for i := 2; i < len(input); i++ {
		tmp := strings.Split(input[i], " = ")
		n := Node{Head: tmp[0]}
		lr := strings.Split(strings.Trim(tmp[1], "()"), ", ")
		n.Left = lr[0]
		n.Right = lr[1]
		network = append(network, n)
	}

	var networkMap map[string]Node = make(map[string]Node)
	for _, v := range network {
		networkMap[v.Head] = v
	}

	//result, loop := findEnd(ins, networkMap)

	var startNodes []string
	for _, v := range network {
		if v.Head[2] == 'A' {
			startNodes = append(startNodes, v.Head)
		}
	}

	fmt.Printf("Starting nodes are %q\n", startNodes)

	var results []int64
	for _, s := range startNodes {
		y, _ := findEnd2(s, ins, networkMap)
		results = append(results, y)
	}

	fmt.Println(lcmOfList(results))

}

// gcd calculates the Greatest Common Divisor using the Euclidean algorithm
func gcd(a, b int64) int64 {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// lcm calculates the Least Common Multiple of two integers
func lcm(a, b int64) int64 {
	return a * b / gcd(a, b)
}

// lcmOfList calculates the LCM of a slice of integers
func lcmOfList(arr []int64) int64 {
	lcmValue := arr[0]
	for _, num := range arr[1:] {
		lcmValue = lcm(lcmValue, num)
	}
	return lcmValue
}

func findEnd(ins string, networkMap map[string]Node) (int64, bool) {
	var result int64
	pos := "AAA"
	loop := true
	for loop {
		for _, v := range ins {
			if v == 'L' {
				pos = networkMap[pos].Left
			}
			if v == 'R' {
				pos = networkMap[pos].Right
			}
			result++
			if pos == "ZZZ" {
				loop = false
				break
			}
		}
	}
	fmt.Println(result)
	return result, loop
}

func findEnd2(pos string, ins string, networkMap map[string]Node) (int64, bool) {
	var result int64
	loop := true
	for loop {
		for _, v := range ins {
			if v == 'L' {
				pos = networkMap[pos].Left
			}
			if v == 'R' {
				pos = networkMap[pos].Right
			}
			result++
			if pos[2] == 'Z' {
				loop = false
				break
			}
		}
	}
	fmt.Printf("pos: %q has end %v\n", pos, result)
	return result, loop
}

type Node struct {
	Head, Left, Right string
}

func (n Node) String() string {
	return fmt.Sprintf("Head: %v, Left: %v, Right: %v", n.Head, n.Left, n.Right)
}

func allZ(ss []string, result int64) bool {
	for i, s := range ss {
		if s[2] != 'Z' {
			return false
		}

		if s[2] == 'Z' {
			fmt.Printf("index %v reached it's end in %v iterations\n", i, result)
		}
	}
	return true
}
