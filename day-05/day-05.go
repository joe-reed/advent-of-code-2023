package main

import (
	"regexp"
	"strings"
	. "utils"

	"github.com/samber/lo"
)

func puzzle1(input string) (result int) {
	seeds, maps := parseInput(input)

	return solve(seeds, maps)
}

// this takes nearly 8 minutes to run on my laptop; forgive me
func puzzle2(input string) (result int) {
	seedRanges, maps := parseInput(input)

	seeds := []int{}

	for i := 0; i < len(seedRanges); i += 2 {
		for j := seedRanges[i]; j < seedRanges[i]+seedRanges[i+1]; j++ {
			seeds = append(seeds, j)
		}
	}

	return solve(seeds, maps)
}

func solve(seeds []int, maps []Map) (result int) {
	for i := range seeds {
		for _, m := range maps {
			for _, part := range m.parts {
				if seeds[i] >= part.sourceStart && seeds[i] < part.sourceStart+part.length {
					seeds[i] = seeds[i] - part.sourceStart + part.destinationStart
					break
				}
			}
		}
	}
	return lo.Min(seeds)
}

func parseInput(input string) (seeds []int, maps []Map) {
	split := strings.Split(input, "\n\n")

	seeds = MapToInts(lo.Map(regexp.MustCompile(`\d+`).FindAllStringSubmatch(split[0], -1), func(m []string, _ int) string {
		return m[0]
	}))

	for _, mapString := range split[1:] {
		m := Map{}
		for i, line := range strings.Split(mapString, "\n") {
			if i == 0 {
				continue
			}
			numbers := MapToInts(lo.Map(regexp.MustCompile(`(\d+)`).FindAllStringSubmatch(line, -1), func(m []string, _ int) string {
				return m[0]
			}))
			part := MapPart{numbers[1], numbers[0], numbers[2]}
			m.parts = append(m.parts, part)
		}
		maps = append(maps, m)
	}
	return seeds, maps
}

type SeedRange struct {
	start, length int
}

type MapPart struct {
	sourceStart, destinationStart, length int
}

type Map struct {
	parts []MapPart
}
