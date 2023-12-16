package adventOfCode23

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type IndexedInt struct {
	Value int
	Index int
}

type IndexedInts []IndexedInt

func (x IndexedInts) Len() int           { return len(x) }
func (x IndexedInts) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }
func (x IndexedInts) Less(i, j int) bool { return x[i].Index < x[j].Index }

func processOperation(op string, boxes []map[string]IndexedInt, index int) {
	if strings.Contains(op, "=") {
		processInsertion(op, boxes, index)
	} else if strings.Contains(op, "-") {
		processRemoval(op, boxes)
	}
}

func processInsertion(op string, boxes []map[string]IndexedInt, index int) {
	parts := strings.Split(op, "=")
	label := parts[0]
	boxIndex := HASH(label)
	focalLength, _ := strconv.Atoi(parts[1])

	value, exists := boxes[boxIndex][label]
	if !exists {
		value.Index = index
	}
	value.Value = focalLength

	boxes[boxIndex][label] = value
}

func processRemoval(op string, boxes []map[string]IndexedInt) {
	label := strings.Trim(op, "-")
	boxIndex := HASH(label)

	delete(boxes[boxIndex], label)
}

func calculateFocusingPower(boxes []map[string]IndexedInt) int64 {
	var totalPower int64
	for i, box := range boxes {
		lenses := make(IndexedInts, 0, len(box))
		for _, lens := range box {
			lenses = append(lenses, lens)
		}

		sort.Sort(lenses)
		for j, lens := range lenses {
			totalPower += int64((i + 1) * (j + 1) * lens.Value)
		}
	}
	return totalPower
}

func HASH(ss string) int {
	var result int
	for i := range ss {
		result += int(ss[i])
		result *= 17
		result %= 256
	}
	return result
}

func Day15() {
	_, input := Input(15)
	operations := strings.Split(input[0], ",")
	boxes := make([]map[string]IndexedInt, 256)
	var hashSum int64

	for i := range boxes {
		boxes[i] = make(map[string]IndexedInt)
	}

	for i, op := range operations {
		hashSum += int64(HASH(op))
		processOperation(op, boxes, i)
	}

	fmt.Println("Sum of HASH values:", hashSum)
	fmt.Println("Total Focusing Power:", calculateFocusingPower(boxes))
}
