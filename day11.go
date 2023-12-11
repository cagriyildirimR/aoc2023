package adventOfCode23

import (
	"fmt"
	"math"
)

type Galaxy struct {
	X, Y int64
}

var expantionRate int64

func (g *Galaxy) norm2(g2 Galaxy) float64 {
	x := g.X - g2.X
	y := g.Y - g2.Y
	return math.Sqrt(float64((x * x) + (y * y)))
}

func (g *Galaxy) norm1(g2 Galaxy) int {
	x := g.X - g2.X
	y := g.Y - g2.Y
	return int(math.Abs(float64(x)) + math.Abs(float64(y)))
}

func (g *Galaxy) calculateExpand(eh, ev []int64) {
	x := int64(0)
	for _, v := range ev {
		if v < g.X {
			x++
		}
	}
	y := int64(0)
	for _, v := range eh {
		if v <= g.Y {
			y++
		}
	}
	g.X += x * expantionRate
	g.Y += y * expantionRate
}

type Galaxies []Galaxy

func Day11() {
	_, input := Input(11)
	expantionRate = 999999
	//input := []string{
	//	"...#......",
	//	".......#..",
	//	"#.........",
	//	"..........",
	//	"......#...",
	//	".#........",
	//	".........#",
	//	"..........",
	//	".......#..",
	//	"#...#.....",
	//}

	var eH []int64
	var eV []int64

	for i := range input {
		if allDotsH(input[i]) {
			eH = append(eH, int64(i))
		}
	}

	for j := range input[0] {
		if allDotsV(input, j) {
			eV = append(eV, int64(j))
		}
	}

	var gs Galaxies

	for i := range input {
		for j, v := range input[i] {
			if v == '#' {
				gs = append(gs, Galaxy{
					X: int64(j),
					Y: int64(i),
				})
			}
		}
	}
	for i := range gs {
		gs[i].calculateExpand(eH, eV)
	}

	var shortestDistance int64
	for i := 0; i < len(gs)-1; i++ {
		for j := i + 1; j < len(gs); j++ {
			shortestDistance += int64(gs[i].norm1(gs[j]))
		}
	}

	fmt.Println(shortestDistance)
}

func allDotsH(inp string) bool {
	for _, v := range inp {
		if v != '.' {
			return false
		}
	}
	return true
}

func allDotsV(input []string, j int) bool {
	for i := 0; i < len(input); i++ {
		if input[i][j] != '.' {
			return false
		}
	}
	return true
}
