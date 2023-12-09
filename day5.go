package adventOfCode23

import (
	"errors"
	"fmt"
	"log"
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

type Range struct {
	start, end int64
	processed  bool
}

func Day5() {
	_, input := Input(5)

	maps := []string{
		"seed-to-soil map:", "soil-to-fertilizer map:", "fertilizer-to-water map:",
		"water-to-light map:", "light-to-temperature map:", "temperature-to-humidity map:",
		"humidity-to-location map:",
	}

	m := make(map[string][]Mapper)
	for _, key := range maps {
		m[key] = []Mapper{}
	}

	seeds, err := parseSeeds(strings.Split(input[0], ":")[1])
	if err != nil {
		log.Fatalf("Error parsing seeds: %v", err)
	}

	key := ""
	for _, line := range input[1:] {
		if contains(maps, line) {
			key = line
		} else {
			mapper, err := parseMapper(line)
			if err == nil {
				m[key] = append(m[key], mapper)
			}
		}
	}

	processSeedsThroughLayers(&seeds, m, maps)

	rangeSeeds := createRangeSeeds(seeds)
	processRangeSeedsThroughLayers(&rangeSeeds, m, maps)

	minEndValue := findMinEndValue(rangeSeeds)
	fmt.Printf("Result: %d\n", minimum(seeds))
	fmt.Printf("Minimum end value: %d\n", minEndValue.start)
}

func contains(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

func parseSeeds(input string) (Seeds, error) {
	var seeds Seeds
	ss := strings.Fields(input)
	for _, v := range ss {
		w, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return nil, err
		}
		seeds = append(seeds, w)
	}
	return seeds, nil
}

func parseMapper(input string) (Mapper, error) {
	numbers := strings.Fields(input)
	if len(numbers) < 3 {
		return Mapper{}, errors.New("invalid mapper format")
	}

	destination, err1 := strconv.ParseInt(numbers[0], 10, 64)
	source, err2 := strconv.ParseInt(numbers[1], 10, 64)
	rang, err3 := strconv.ParseInt(numbers[2], 10, 64)
	if err1 != nil || err2 != nil || err3 != nil {
		return Mapper{}, fmt.Errorf("error parsing mapper: %v, %v, %v", err1, err2, err3)
	}

	return Mapper{Destination: destination, Source: source, Rang: rang}, nil
}

func processSeedsThroughLayers(seeds *Seeds, m map[string][]Mapper, maps []string) {
	for _, v := range maps {
		for i, s := range *seeds {
			for _, w := range m[v] {
				if b, l := isIn(s, w); b {
					(*seeds)[i] = l
					break
				}
			}
		}
	}
}

func minimum(seeds Seeds) int64 {
	m := seeds[0]
	for _, s := range seeds[1:] {
		if s < m {
			m = s
		}
	}
	return m
}

func createRangeSeeds(seeds Seeds) []Range {
	var rangeSeeds []Range
	for i := 0; i < len(seeds); i += 2 {
		rangeSeeds = append(rangeSeeds, Range{
			start:     seeds[i],
			end:       seeds[i] + seeds[i+1],
			processed: false,
		})
	}
	return rangeSeeds
}

func processRangeSeedsThroughLayers(rangeSeeds *[]Range, m map[string][]Mapper, maps []string) {
	for _, v := range maps {
		for i := 0; i < len(*rangeSeeds); i++ {
			s := (*rangeSeeds)[i]
			for _, w := range m[v] {
				rs := rangeMap(s, w)
				if len(rs) != 0 {
					for _, r := range rs {
						if r.processed {
							(*rangeSeeds)[i] = r
						} else {
							*rangeSeeds = append(*rangeSeeds, r)
						}
					}
					break
				}
			}
		}
		for i := range *rangeSeeds {
			(*rangeSeeds)[i].processed = false
		}
	}
}

func findMinEndValue(rangeSeeds []Range) Range {
	minEndValue := rangeSeeds[0]
	for _, r := range rangeSeeds {
		if r.start < minEndValue.start {
			minEndValue = r
		}
	}
	return minEndValue
}

func isIn(seed int64, m Mapper) (bool, int64) {
	if seed >= m.Source && seed < m.Source+m.Rang {
		return true, m.Destination + (seed - m.Source)
	}
	return false, seed
}

func rangeMap(sourceRange Range, mapper Mapper) []Range {
	var result []Range
	offset := mapper.Destination - mapper.Source

	if isOverlapInFirstScenario(sourceRange, mapper) {
		result = append(result, Range{mapper.Source + offset, sourceRange.end + offset, true})
		if sourceRange.start != mapper.Source {
			result = append(result, Range{sourceRange.start, mapper.Source, false})
		}
		return result
	}

	if isOverlapInSecondScenario(sourceRange, mapper) {
		result = append(result, Range{sourceRange.start + offset, mapper.Source + mapper.Rang + offset, true})
		if sourceRange.end != mapper.Source+mapper.Rang {
			result = append(result, Range{mapper.Source + mapper.Rang, sourceRange.end, false})
		}
		return result
	}

	if isCompleteOverlap(sourceRange, mapper) {
		result = append(result, Range{mapper.Source + offset, mapper.Source + mapper.Rang + offset, true})
		if sourceRange.start != mapper.Source {
			result = append(result, Range{sourceRange.start, mapper.Source, false})
		}
		if sourceRange.end != mapper.Source+mapper.Rang {
			result = append(result, Range{mapper.Source + mapper.Rang, sourceRange.end, false})
		}
		return result
	}

	if isContainedWithin(sourceRange, mapper) {
		result = append(result, Range{sourceRange.start + offset, sourceRange.end + offset, true})
		return result
	}

	return result
}

func isOverlapInFirstScenario(sourceRange Range, mapper Mapper) bool {
	return sourceRange.start <= mapper.Source && sourceRange.end > mapper.Source && sourceRange.end < mapper.Source+mapper.Rang
}

func isOverlapInSecondScenario(sourceRange Range, mapper Mapper) bool {
	return sourceRange.start > mapper.Source && sourceRange.start < mapper.Source+mapper.Rang && sourceRange.end >= mapper.Source+mapper.Rang
}

func isCompleteOverlap(sourceRange Range, mapper Mapper) bool {
	return sourceRange.start <= mapper.Source && sourceRange.end >= mapper.Source+mapper.Rang
}

func isContainedWithin(sourceRange Range, mapper Mapper) bool {
	return sourceRange.start >= mapper.Source && sourceRange.end <= mapper.Source+mapper.Rang
}
