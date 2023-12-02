package main

import (
	"regexp"
	. "utils"
)

func puzzle1(input []string) (result int) {
	return solve(input, func(maxes map[string]int, game int) int {
		if maxes["red"] <= 12 && maxes["green"] <= 13 && maxes["blue"] <= 14 {
			return game
		}
		return 0
	})
}

func puzzle2(input []string) (result int) {
	return solve(input, func(maxes map[string]int, game int) int {
		return maxes["red"] * maxes["green"] * maxes["blue"]
	})
}

func solve(input []string, solver func(maxes map[string]int, game int) int) (result int) {
	for i, line := range input {
		game := regexp.MustCompile(`Game \d+: (.+)`).FindStringSubmatch(line)[1]

		draws := regexp.MustCompile(`(\d+) (\w+)`).FindAllStringSubmatch(game, -1)

		maxes := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}

		for _, draw := range draws {
			if ToInt(draw[1]) > maxes[draw[2]] {
				maxes[draw[2]] = ToInt(draw[1])
			}
		}

		result += solver(maxes, i+1)
	}

	return
}
