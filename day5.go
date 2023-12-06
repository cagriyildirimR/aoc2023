package adventOfCode23

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

type Mapper struct {
	Destination int64
	Source      int64
	Rang        int64
}

func (m *Mapper) String() string {
	return fmt.Sprintf("source: %v, destination: %v with range: %v", m.Source, m.Destination, m.Rang)
}

type Seeds []int64

func Day5() {
	_, input := Input(5)

	m := make(map[string][]Mapper)
	m["seed-to-soil map:"] = []Mapper{}
	m["soil-to-fertilizer map:"] = []Mapper{}
	m["fertilizer-to-water map:"] = []Mapper{}
	m["water-to-light map:"] = []Mapper{}
	m["light-to-temperature map:"] = []Mapper{}
	m["temperature-to-humidity map:"] = []Mapper{}
	m["humidity-to-location map:"] = []Mapper{}

	var seeds Seeds
	ss := strings.Fields(strings.Split(input[0], ":")[1])
	for _, v := range ss {
		w, _ := strconv.Atoi(v)
		seeds = append(seeds, int64(w))
	}

	var key string
	for _, v := range input[1:] {

		switch v {
		case "seed-to-soil map:":
			key = "seed-to-soil map:"
			continue
		case "soil-to-fertilizer map:":
			key = "soil-to-fertilizer map:"
			continue
		case "fertilizer-to-water map:":
			key = "fertilizer-to-water map:"
			continue
		case "water-to-light map:":
			key = "water-to-light map:"
			continue
		case "light-to-temperature map:":
			key = "light-to-temperature map:"
			continue
		case "temperature-to-humidity map:":
			key = "temperature-to-humidity map:"
			continue
		case "humidity-to-location map:":
			key = "humidity-to-location map:"
			continue
		case "":
			continue
		default:
			numbers := strings.Fields(v)
			if len(numbers) < 3 {
				continue
			}
			destination, err1 := strconv.ParseInt(numbers[0], 10, 64)
			source, err2 := strconv.ParseInt(numbers[1], 10, 64)
			rang, err3 := strconv.ParseInt(numbers[2], 10, 64)

			if err1 != nil || err2 != nil || err3 != nil {
				continue
			}

			mapper := Mapper{Destination: destination, Source: source, Rang: rang}
			m[key] = append(m[key], mapper)
		}
	}

	lo := []string{"seed-to-soil map:", "soil-to-fertilizer map:", "fertilizer-to-water map:", "water-to-light map:", "light-to-temperature map:", "temperature-to-humidity map:", "humidity-to-location map:"}
	for _, v := range lo { // loop through map
		for i, s := range seeds {
			for _, w := range m[v] {
				b, l := isIn(s, w)
				if b {
					seeds[i] = l
					break
				}
			}
		}
	}

	result := seeds[0]
	for _, s := range seeds {
		//fmt.Println(s)
		if result > s {
			result = s
		}
	}

	fmt.Println(result)

	seeds = Seeds{}
	for _, v := range ss {
		w, _ := strconv.Atoi(v)
		seeds = append(seeds, int64(w))
	}

	var rangeSeeds []Range
	log.Println("Initializing rangeSeeds")
	for i := 0; i < len(seeds); i = i + 2 {
		rangeSeeds = append(rangeSeeds, Range{
			start:     seeds[i],
			end:       seeds[i] + seeds[i+1],
			processed: false,
		})
	}

	for _, v := range lo { // layers
		for i := 0; i < len(rangeSeeds); i++ { // take seed
			s := rangeSeeds[i]
			for _, w := range m[v] { // ranges inside layers
				rs := rangeMap(s, w)
				if len(rs) != 0 {
					for _, r := range rs {
						if r.processed {
							rangeSeeds[i] = r
						} else {
							rangeSeeds = append(rangeSeeds, r)
						}
					}
					break
				}
			}
		}
		for i := range rangeSeeds {
			rangeSeeds[i].processed = false
		}
	}

	log.Println("Calculating minimum end value")
	mx := Range{
		start:     math.MaxInt64,
		end:       math.MaxInt64,
		processed: false,
	}
	for _, r := range rangeSeeds {
		if r.start < mx.start {
			log.Printf("New minimum end value found: %d\n", r.end)
			mx = r
		}
	}
	fmt.Println("Minimum end value:", mx.start)

}

func isIn(seed int64, m Mapper) (bool, int64) {
	if seed >= m.Source && seed < m.Source+m.Rang {
		return true, m.Destination + (seed - m.Source)
	}
	return false, seed
}

type Range struct {
	start, end int64
	processed  bool
}

func rangeMap(sourceRange Range, mapper Mapper) []Range {
	var result []Range
	offset := mapper.Destination - mapper.Source

	if sourceRange.start <= mapper.Source && sourceRange.end > mapper.Source && sourceRange.end < mapper.Source+mapper.Rang {
		result = append(result, Range{mapper.Source + offset, sourceRange.end + offset, true})
		if sourceRange.start != mapper.Source {
			result = append(result, Range{sourceRange.start, mapper.Source, false})
		}

		return result
	}

	if sourceRange.start > mapper.Source && sourceRange.start < mapper.Source+mapper.Rang && sourceRange.end >= mapper.Source+mapper.Rang {
		result = append(result, Range{sourceRange.start + offset, mapper.Source + mapper.Rang + offset, true})

		if sourceRange.end != mapper.Source+mapper.Rang {
			result = append(result, Range{mapper.Source + mapper.Rang, sourceRange.end, false})
		}

		return result
	}

	if sourceRange.end < mapper.Source {
		return result
	}

	if sourceRange.start > mapper.Source+mapper.Rang {
		return result
	}

	if sourceRange.start <= mapper.Source && sourceRange.end >= mapper.Source+mapper.Rang {
		result = append(result, Range{mapper.Source + offset, mapper.Source + mapper.Rang + offset, true})
		if sourceRange.start != mapper.Source {
			result = append(result, Range{sourceRange.start, mapper.Source, false})
		}

		if sourceRange.end != mapper.Source+mapper.Rang {
			result = append(result, Range{mapper.Source + mapper.Rang, sourceRange.end, false})
		}
	}

	if sourceRange.start >= mapper.Source && sourceRange.end <= mapper.Source+mapper.Rang {
		result = append(result, Range{sourceRange.start + offset, sourceRange.end + offset, true})
		return result
	}

	return result
}
