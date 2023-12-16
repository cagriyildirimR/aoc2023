package adventOfCode23

import (
	"fmt"
	"sort"
)

type TileType int

const (
	EmptySpace TileType = iota
	MirrorSlash
	MirrorBackslash
	SplitterHorizontal
	SplitterVertical
)

type Tile struct {
	North, South, East, West bool // can go to direction
	Type                     TileType
	Visited                  bool
	Position                 Position
}

func (t Tile) String() string {
	switch t.Type {
	case EmptySpace:
		return "."
	case MirrorSlash:
		return "/"
	case MirrorBackslash:
		return "X"
	case SplitterHorizontal:
		return "-"
	case SplitterVertical:
		return "|"
	default:
		return "?"
	}
}

type Position struct {
	X, Y int
}

type Beam struct {
	Past, Current Position
}

func parseDay16() [][]Tile {
	_, input := Input(16)
	//input = []string{
	//	".|...\\....",
	//	"|.-.\\.....",
	//	".....|-...",
	//	"........|.",
	//	"..........",
	//	".........\\",
	//	"..../.\\\\..",
	//	".-.-/..|..",
	//	".|....-|.\\",
	//	"..//.|....",
	//}

	t := make([][]Tile, len(input))

	for i, v := range input {
		t[i] = make([]Tile, len(v))
		for j, w := range []byte(v) {
			var tmp int
			switch w {
			case '.':
				tmp = 0
			case '/':
				tmp = 1
			case '\\':
				tmp = 2
			case '-':
				tmp = 3
			case '|':
				tmp = 4
			}
			t[i][j] = Tile{
				North:   true,
				South:   true,
				East:    true,
				West:    true,
				Type:    TileType(tmp),
				Visited: false,
				Position: Position{
					X: j,
					Y: i,
				},
			}
		}
	}
	return t
}

func direction(beam Beam) Position {
	return Position{
		X: beam.Current.X - beam.Past.X,
		Y: beam.Current.Y - beam.Past.Y,
	}
}

func next(current Position, direction Position) Position {
	return Position{
		X: current.X + direction.X,
		Y: current.Y + direction.Y,
	}
}

func mirrorSlashDirection(d Position) Position {
	return Position{
		X: -d.Y,
		Y: -d.X,
	}
}

func mirrorBackslashDirection(d Position) Position {
	return Position{
		X: d.Y,
		Y: d.X,
	}
}

func move(beam Beam, grid [][]Tile) []Beam {
	var result []Beam
	if beam.Current.Y < 0 || beam.Current.Y >= len(grid) || beam.Current.X < 0 || beam.Current.X >= len(grid[0]) {
		return result
	}

	grid[beam.Current.Y][beam.Current.X].Visited = true
	ct := grid[beam.Current.Y][beam.Current.X].Type
	d := direction(beam)

	checkDir := func(d Position) bool {
		if d.Y == 0 && d.X > 0 {
			return grid[beam.Current.Y][beam.Current.X].East
		}
		if d.Y == 0 && d.X < 0 {
			return grid[beam.Current.Y][beam.Current.X].West
		}
		if d.Y > 0 && d.X == 0 {
			return grid[beam.Current.Y][beam.Current.X].South
		}
		return grid[beam.Current.Y][beam.Current.X].North
	}

	changeDir := func(d Position) {
		if d.Y == 0 && d.X > 0 {
			grid[beam.Current.Y][beam.Current.X].East = false
		}
		if d.Y == 0 && d.X < 0 {
			grid[beam.Current.Y][beam.Current.X].West = false
		}
		if d.Y > 0 && d.X == 0 {
			grid[beam.Current.Y][beam.Current.X].South = false
		}
		if d.Y < 0 && d.X == 0 {
			grid[beam.Current.Y][beam.Current.X].North = false
		}
	}
	switch ct {
	case EmptySpace:
		if checkDir(d) {
			result = append(result, Beam{
				Past:    beam.Current,
				Current: next(beam.Current, d),
			})
			changeDir(d)
		}
	case MirrorSlash:
		d = mirrorSlashDirection(d)
		if checkDir(d) {
			result = append(result, Beam{
				Past:    beam.Current,
				Current: next(beam.Current, d),
			})
			changeDir(d)
		}
	case MirrorBackslash:
		d = mirrorBackslashDirection(d)
		if checkDir(d) {
			result = append(result, Beam{
				Past:    beam.Current,
				Current: next(beam.Current, d),
			})
			changeDir(d)
		}
	case SplitterHorizontal:
		if d.Y == 0 {
			if checkDir(d) {
				result = append(result, Beam{
					Past:    beam.Current,
					Current: next(beam.Current, d),
				})
				changeDir(d)
			}
		} else {
			tmp := d
			d = mirrorSlashDirection(d)
			if checkDir(d) {
				result = append(result, Beam{
					Past:    beam.Current,
					Current: next(beam.Current, d),
				})
				changeDir(d)
			}
			d = mirrorBackslashDirection(tmp)
			if checkDir(d) {
				result = append(result, Beam{
					Past:    beam.Current,
					Current: next(beam.Current, d),
				})
				changeDir(d)
			}
		}
	case SplitterVertical:
		if d.X == 0 {
			if checkDir(d) {
				result = append(result, Beam{
					Past:    beam.Current,
					Current: next(beam.Current, d),
				})
				changeDir(d)
			}
		} else {
			tmp := d
			if checkDir(d) {
				d = mirrorSlashDirection(d)
				result = append(result, Beam{
					Past:    beam.Current,
					Current: next(beam.Current, d),
				})
				changeDir(d)

			}
			d = mirrorBackslashDirection(tmp)
			if checkDir(d) {
				result = append(result, Beam{
					Past:    beam.Current,
					Current: next(beam.Current, d),
				})
				changeDir(d)

			}
		}
	}
	return result
}

func Day16() {
	g := parseDay16()
	// part 1
	bs := []Beam{
		{
			Past:    Position{-1, 0},
			Current: Position{0, 0},
		},
	}

	for i := 0; i < len(bs); i++ {
		bs = append(bs, move(bs[i], g)...)
	}

	var result int
	for i := range g {
		for j := range g[i] {
			if g[i][j].Visited {
				result++
			}
		}
	}

	fmt.Println(result)

	// part 2

	var results []int

	for i := 0; i < len(g[0]); i++ {
		grid := parseDay16()
		bs := []Beam{
			{
				Past:    Position{i, -1},
				Current: Position{i, 0},
			},
		}
		for i := 0; i < len(bs); i++ {
			bs = append(bs, move(bs[i], grid)...)
		}

		var result int
		for i := range grid {
			for j := range grid[i] {
				if grid[i][j].Visited {
					result++
				}
			}
		}
		results = append(results, result)
	}

	for i := 0; i < len(g[0]); i++ {
		grid := parseDay16()
		bs := []Beam{
			{
				Past:    Position{i, len(grid)},
				Current: Position{i, len(grid) - 1},
			},
		}
		for i := 0; i < len(bs); i++ {
			bs = append(bs, move(bs[i], grid)...)
		}

		var result int
		for i := range grid {
			for j := range grid[i] {
				if grid[i][j].Visited {
					result++
				}
			}
		}
		results = append(results, result)
	}

	for i := 0; i < len(g); i++ {
		grid := parseDay16()
		bs := []Beam{
			{
				Past:    Position{-1, i},
				Current: Position{0, i},
			},
		}
		for i := 0; i < len(bs); i++ {
			bs = append(bs, move(bs[i], grid)...)
		}

		var result int
		for i := range grid {
			for j := range grid[i] {
				if grid[i][j].Visited {
					result++
				}
			}
		}
		results = append(results, result)
	}

	for i := 0; i < len(g); i++ {
		grid := parseDay16()
		bs := []Beam{
			{
				Past:    Position{len(grid[0]), i},
				Current: Position{len(grid[0]) - 1, i},
			},
		}
		for i := 0; i < len(bs); i++ {
			bs = append(bs, move(bs[i], grid)...)
		}

		var result int
		for i := range grid {
			for j := range grid[i] {
				if grid[i][j].Visited {
					result++
				}
			}
		}
		results = append(results, result)
	}

	sort.Ints(results)
	fmt.Println(results[len(results)-1])
}
