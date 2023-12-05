package main

import (
	"strings"
	. "utils"

	"github.com/samber/lo"
)

type Card struct {
	winningNumbers []int
	myNumbers      []int
}

func puzzle1(input string) (result int) {
	cards := parseCards(strings.Split(input, "\n"))

	for _, card := range cards {
		r := 0
		for _, myNumber := range card.myNumbers {
			if lo.Contains(card.winningNumbers, myNumber) {
				if r == 0 {
					r = 1
					continue
				}
				r *= 2
			}
		}
		result += r
	}

	return
}

func puzzle2(input string) (result int) {
	cards := parseCards(strings.Split(input, "\n"))

	return solvePuzzle2(cards, 0, len(cards))
}

func solvePuzzle2(cards []Card, start, length int) (result int) {
	for i, card := range cards[start : start+length] {
		result++
		matches := 0
		for _, myNumber := range card.myNumbers {
			if lo.Contains(card.winningNumbers, myNumber) {
				matches++
			}
		}
		if matches != 0 {
			result += solvePuzzle2(cards, start+i+1, matches)
		}
	}
	return
}

func parseCards(lines []string) (cards []Card) {
	for _, line := range lines {
		c := Card{}
		numbers := strings.Split(line, ":")[1]

		split := strings.Split(numbers, "|")
		winningNumbers, myNumbers := split[0], split[1]

		n := ""
		for _, char := range strings.Split(winningNumbers, "") {
			if char == " " {
				if n != "" {
					c.winningNumbers = append(c.winningNumbers, ToInt(n))
					n = ""
				}
				continue
			}
			n += char
		}

		n = ""
		for _, char := range strings.Split(myNumbers, "") {
			if char == " " {
				if n != "" {
					c.myNumbers = append(c.myNumbers, ToInt(n))
					n = ""
				}
				continue
			}
			n += char
		}
		if n != "" {
			c.myNumbers = append(c.myNumbers, ToInt(n))
			n = ""
		}

		cards = append(cards, c)
	}
	return
}
