package adventOfCode23

import "fmt"

func Day14() {
	_, input := Input(14)

	//input := []string{
	//	"O....#....",
	//	"O.OO#....#",
	//	".....##...",
	//	"OO.#O....O",
	//	".O.....O#.",
	//	"O.#..O.#.#",
	//	"..O..#O..O",
	//	".......O..",
	//	"#....###..",
	//	"#OO..#....",
	//}

	var result int64
	for j := 0; j < len(input[0]); j++ {
		anc := len(input)
		for i := 0; i < len(input); i++ {
			if input[i][j] == '#' {
				anc = len(input) - i - 1
			}
			if input[i][j] == 'O' {
				result += int64(anc)
				anc--
			}
		}
	}

	fmt.Println(result)

	part2(input)
}

type Direction struct {
	X, Y int
}

func part2(input []string) {
	rows, cols := len(input), len(input[0])
	grid := make([][]rune, rows)
	for i := range grid {
		grid[i] = make([]rune, cols)
		for j := range grid[i] {
			grid[i][j] = rune(input[i][j])
		}
	}

	var cycle []int64

	for t := 0; t < 1000000000; t++ {
		// North
		for j := 0; j < cols; j++ {
			anc := 0
			for i := 0; i < rows; i++ {
				if grid[i][j] == '#' {
					anc = i + 1
				}
				if grid[i][j] == 'O' {
					grid[i][j] = '.'
					if anc < rows {
						grid[anc][j] = 'O'
						anc++
					}
				}
			}
		}

		// West
		for i := 0; i < rows; i++ {
			anc := 0
			for j := 0; j < cols; j++ {
				if grid[i][j] == '#' {
					anc = j + 1
				}
				if grid[i][j] == 'O' {
					grid[i][j] = '.'
					if anc < cols {
						grid[i][anc] = 'O'
						anc++
					}
				}
			}
		}

		// South
		for j := 0; j < cols; j++ {
			anc := rows - 1
			for i := rows - 1; i >= 0; i-- {
				if grid[i][j] == '#' {
					anc = i - 1
				}
				if grid[i][j] == 'O' {
					grid[i][j] = '.'
					if anc >= 0 {
						grid[anc][j] = 'O'
						anc--
					}
				}
			}
		}

		// East
		for i := 0; i < rows; i++ {
			anc := cols - 1
			for j := cols - 1; j >= 0; j-- {
				if grid[i][j] == '#' {
					anc = j - 1
				}
				if grid[i][j] == 'O' {
					grid[i][j] = '.'
					if anc >= 0 {
						grid[i][anc] = 'O'
						anc--
					}
				}
			}
		}
		cycle = append(cycle, calculateTotalLoad(grid))
		if len(cycle) > 20 {
			vs, b := findCycle(cycle)
			if b {
				fmt.Println(vs[(1000000000-t-1)%len(vs)])
				return
			} else {
				cycle = []int64{}
			}
		}
	}
}

func findCycle(nums []int64) ([]int64, bool) {
	if len(nums) == 0 {
		return nil, false
	}

	seen := make(map[int64]int)
	cycle := []int64{}
	hasCycle := false

	for i, num := range nums {
		if index, found := seen[num]; found {
			hasCycle = true
			cycle = nums[index:i]
			break
		}
		seen[num] = i
	}

	return cycle, hasCycle
}

func calculateTotalLoad(grid [][]rune) int64 {
	rows := len(grid)
	var totalLoad int64

	for i, row := range grid {
		for _, cell := range row {
			if cell == 'O' {
				totalLoad += int64(rows - i)
			}
		}
	}

	return totalLoad
}
