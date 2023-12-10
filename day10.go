package adventOfCode23

import (
	"fmt"
	"strings"
)

type Coordinate struct {
	X, Y  int
	Char  byte
	Count int
}

func Day10() {
	_, input := Input(10)

	vs := make([][]byte, len(input))
	for i := range vs {
		vs[i] = make([]byte, len(input[i]))
		for j := range vs[i] {
			vs[i][j] = input[i][j]
		}
	}

	visited := make([][]bool, len(input))
	for i := range visited {
		visited[i] = make([]bool, len(input[i]))
	}

	S, N := getStart(input)
	visited[S.Y][S.X] = true

	p1 := N[1]
	f := true
	for f {
		f = p1.Next(visited, input, vs)
	}

	fmt.Println((p1.Count + 1) / 2) // taking the farthest value

	// part 2
	evs := make([][]byte, len(input)*3)
	for i := range evs {
		evs[i] = make([]byte, len(input[0])*3)
	}

	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			e := expand(input[i][j])
			evs[i*3][j*3] = e[0][0]
			evs[i*3][j*3+1] = e[0][1]
			evs[i*3][j*3+2] = e[0][2]
			evs[i*3+1][j*3] = e[1][0]
			evs[i*3+1][j*3+1] = e[1][1]
			evs[i*3+1][j*3+2] = e[1][2]
			evs[i*3+2][j*3] = e[2][0]
			evs[i*3+2][j*3+1] = e[2][1]
			evs[i*3+2][j*3+2] = e[2][2]
		}
	}

	expandedVisited := make([][]bool, len(evs))
	for i := range expandedVisited {
		expandedVisited[i] = make([]bool, len(evs[i]))
	}

	var expandedInput []string
	for _, v := range evs {
		expandedInput = append(expandedInput, string(v))
	}

	ES, EN := getStart(expandedInput)
	expandedVisited[ES.Y][ES.X] = true
	evs[ES.Y][ES.X] = '*'

	ep1 := EN[1]
	ef := true
	for ef {
		ef = ep1.Next(expandedVisited, expandedInput, evs)
	}

	//fmt.Println((ep1.Count + 1) / 6) // checking the value

	for i := 0; i < len(evs); i++ {
		floodFill(evs, len(evs)-1, i)
		floodFill(evs, 0, i)
	}
	for j := 0; j < len(evs[0]); j++ {
		floodFill(evs, j, 0)
		floodFill(evs, j, len(evs[0])-1)
	}

	reducedRows := len(evs) / 3
	reducedCols := len(evs[0]) / 3

	reduced := make([][]byte, reducedRows)
	for i := range reduced {
		reduced[i] = make([]byte, reducedCols)
	}

	for i := 0; i < len(evs)-2; i += 3 {
		for j := 0; j < len(evs[i])-2; j += 3 {
			block := getBlock(evs, i, j)
			reduced[i/3][j/3] = checkShape(block)
		}
	}

	var r int
	for _, row := range reduced {
		for _, val := range row {
			if val != '*' {
				r++
			}
		}
	}
	fmt.Println(r)
}

func getBlock(evs [][]byte, i, j int) [][]byte {
	block := make([][]byte, 3)
	for di := 0; di < 3; di++ {
		block[di] = make([]byte, 3)
		for dj := 0; dj < 3; dj++ {
			block[di][dj] = evs[i+di][j+dj]
		}
	}
	return block
}

func checkShape(block [][]byte) byte {
	for i := range block {
		for j := range block[i] {
			if block[i][j] == '*' {
				return '*'
			}
		}
	}
	return 'I'
}

// Next checks current position and moves towards unvisited path
// returns true if lands on unvisited pipe
// else returns false signifying the end of unvisited pipes
func (p *Coordinate) Next(visited [][]bool, input []string, vs [][]byte) bool {
	visited[p.Y][p.X] = true
	vs[p.Y][p.X] = '*'
	moved := false

	switch p.Char {
	case 'J':
		// J connects north and west
		if p.Y > 0 && !visited[p.Y-1][p.X] {
			p.Count++
			p.Y--
			moved = true
		} else if p.X > 0 && !visited[p.Y][p.X-1] {
			p.Count++
			p.X--
			moved = true
		}

	case '|':
		// | connects north and south
		if p.Y > 0 && !visited[p.Y-1][p.X] {
			p.Count++
			p.Y--
			moved = true
		} else if p.Y < len(input)-1 && !visited[p.Y+1][p.X] {
			p.Count++
			p.Y++
			moved = true
		}

	case '-':
		// - connects east and west
		if p.X > 0 && !visited[p.Y][p.X-1] {
			p.Count++
			p.X--
			moved = true
		} else if p.X < len(input[0])-1 && !visited[p.Y][p.X+1] {
			p.Count++
			p.X++
			moved = true
		}

	case 'L':
		// L connects north and east
		if p.Y > 0 && !visited[p.Y-1][p.X] {
			p.Count++
			p.Y--
			moved = true
		} else if p.X < len(input[0])-1 && !visited[p.Y][p.X+1] {
			p.Count++
			p.X++
			moved = true
		}

	case '7':
		// 7 connects south and west
		if p.Y < len(input)-1 && !visited[p.Y+1][p.X] {
			p.Count++
			p.Y++
			moved = true
		} else if p.X > 0 && !visited[p.Y][p.X-1] {
			p.Count++
			p.X--
			moved = true
		}

	case 'F':
		// F connects south and east
		if p.Y < len(input)-1 && !visited[p.Y+1][p.X] {
			p.Count++
			p.Y++
			moved = true
		} else if p.X < len(input[0])-1 && !visited[p.Y][p.X+1] {
			p.Count++
			p.X++
			moved = true
		}
	}

	if moved {
		// Update the character at the new position
		p.Char = input[p.Y][p.X]
	}

	return moved
}

// getStart finds and returns S coordinates with possible connected neighbours
func getStart(input []string) (Coordinate, []Coordinate) {
	var neighbors []Coordinate

	// find start
	for i := range input {
		for j := range input[i] {
			if input[i][j] == 'S' {
				start := Coordinate{
					X: j,
					Y: i,
				}
				// Check North
				if i > 0 && strings.Contains("|7F", string(input[i-1][j])) {
					neighbors = append(neighbors, Coordinate{X: j, Y: i - 1, Count: 1, Char: input[i-1][j]})
				}
				// Check South
				if i < len(input)-1 && strings.Contains("|LJ", string(input[i+1][j])) {
					neighbors = append(neighbors, Coordinate{X: j, Y: i + 1, Count: 1, Char: input[i+1][j]})
				}
				// Check West
				if j > 0 && strings.Contains("-FL", string(input[i][j-1])) {
					neighbors = append(neighbors, Coordinate{X: j - 1, Y: i, Count: 1, Char: input[i][j-1]})
				}
				// Check East
				if j < len(input[i])-1 && strings.Contains("-7J", string(input[i][j+1])) {
					neighbors = append(neighbors, Coordinate{X: j + 1, Y: i, Count: 1, Char: input[i][j+1]})
				}
				return start, neighbors
			}
		}
	}
	return Coordinate{}, nil
}

func floodFill(vs [][]byte, i, j int) {
	if i < 0 || i >= len(vs) || j < 0 || j >= len(vs[0]) {
		return
	}
	if vs[i][j] == '*' {
		return
	}

	vs[i][j] = '*'

	floodFill(vs, i+1, j) // Down
	floodFill(vs, i-1, j) // Up
	floodFill(vs, i, j+1) // Right
	floodFill(vs, i, j-1) // Left
}

func expand(x byte) [][]byte {
	switch x {
	case '|':
		return [][]byte{
			{'.', '|', '.'},
			{'.', '|', '.'},
			{'.', '|', '.'},
		}
	case '-':
		return [][]byte{
			{'.', '.', '.'},
			{'-', '-', '-'},
			{'.', '.', '.'},
		}
	case 'L':
		return [][]byte{
			{'.', '|', '.'},
			{'.', 'L', '-'},
			{'.', '.', '.'},
		}
	case 'J':
		return [][]byte{
			{'.', '|', '.'},
			{'-', 'J', '.'},
			{'.', '.', '.'},
		}
	case '7':
		return [][]byte{
			{'.', '.', '.'},
			{'-', '7', '.'},
			{'.', '|', '.'},
		}
	case 'F':
		return [][]byte{
			{'.', '.', '.'},
			{'.', 'F', '-'},
			{'.', '|', '.'},
		}
	case 'S':
		return [][]byte{
			{'.', '|', '.'},
			{'.', 'S', '.'},
			{'.', '|', '.'},
		}
	default: // Ground
		return [][]byte{
			{'.', '.', '.'},
			{'.', '.', '.'},
			{'.', '.', '.'},
		}
	}
}
