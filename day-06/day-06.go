package main

import (
	"regexp"
	"strings"
	. "utils"

	"github.com/samber/lo"
)

func puzzle1(lines []string) (result int) {
	times := extractNumbers(lines[0])
	distances := extractNumbers(lines[1])

	games := []Game{}
	for i := range times {
		games = append(games, Game{times[i], distances[i]})
	}

	return lo.Reduce(games, func(acc int, game Game, _ int) int {
		return acc * getNumOptions(game)
	}, 1)
}

func puzzle2(lines []string) (result int) {
	time := ToInt(strings.Join(extractNumberStrings(lines[0]), ""))
	distance := ToInt(strings.Join(extractNumberStrings(lines[1]), ""))

	return getNumOptions(Game{time, distance})
}

type Game struct {
	time, distance int
}

func getNumOptions(game Game) (numOptions int) {
	for i := 0; i < game.time; i++ {
		distance := (game.time - i) * i
		if distance > game.distance {
			numOptions++
		}
	}
	return
}

func extractNumbers(s string) (numbers []int) {
	return MapToInts(extractNumberStrings(s))
}

func extractNumberStrings(s string) (numberStrings []string) {
	return lo.Map(regexp.MustCompile(`(\d+)`).FindAllStringSubmatch(s, -1), func(m []string, _ int) string { return m[0] })
}
