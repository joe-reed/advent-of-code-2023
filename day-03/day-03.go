package main

import (
	"strconv"
	"strings"
	. "utils"
)

func puzzle1(input string) (result int) {
	lines := strings.Split(input, "\n")
	schematic := make([][]string, len(lines))
	for y, line := range lines {
		schematic[y] = strings.Split(line, "")
	}

	for y, line := range schematic {
		for x := 0; x < len(line); {
			char := line[x]
			if char == "." {
				x++
				continue
			}

			_, err := strconv.Atoi(char)

			if err != nil {
				x++
				continue
			}

			start := x
			end, number := getNumber(line, start)

			included := false

			for i := start - 1; i <= end; i++ {
				if i < 0 {
					continue
				}
				if i > len(line)-1 {
					continue
				}
				if y > 0 && schematic[y-1][i] != "." {
					included = true
					break
				}
				if y < len(schematic)-1 && schematic[y+1][i] != "." {
					included = true
					break
				}
			}

			if start > 0 && schematic[y][start-1] != "." {
				included = true
			}
			if end < len(line)-1 && schematic[y][end] != "." {
				included = true
			}

			if included {
				result += number
			}

			x += end - start
		}
	}

	return
}

type Coord struct {
	x, y int
}

func puzzle2(input string) (result int) {
	lines := strings.Split(input, "\n")
	schematic := make([][]string, len(lines))
	for y, line := range lines {
		schematic[y] = strings.Split(line, "")
	}

	adjCount := map[Coord]int{}
	gearRatios := map[Coord]int{}

	for y, line := range schematic {
		for x := 0; x < len(line); {
			char := line[x]
			if char == "." {
				x++
				continue
			}

			_, err := strconv.Atoi(char)

			if err != nil {
				x++
				continue
			}

			start := x
			end, number := getNumber(line, start)

			for i := start - 1; i <= end; i++ {
				if i < 0 {
					continue
				}
				if i > len(line)-1 {
					continue
				}
				if y > 0 && schematic[y-1][i] != "." {
					adjCount[Coord{i, y - 1}]++
					if gearRatios[Coord{i, y - 1}] == 0 {
						gearRatios[Coord{i, y - 1}] = 1
					}
					gearRatios[Coord{i, y - 1}] *= number
				}
				if y < len(schematic)-1 && schematic[y+1][i] != "." {
					adjCount[Coord{i, y + 1}]++
					if gearRatios[Coord{i, y + 1}] == 0 {
						gearRatios[Coord{i, y + 1}] = 1
					}
					gearRatios[Coord{i, y + 1}] *= number
				}
			}

			if start > 0 && schematic[y][start-1] != "." {
				adjCount[Coord{start - 1, y}]++
				if gearRatios[Coord{start - 1, y}] == 0 {
					gearRatios[Coord{start - 1, y}] = 1
				}
				gearRatios[Coord{start - 1, y}] *= number
			}
			if end < len(line)-1 && schematic[y][end] != "." {
				adjCount[Coord{end, y}]++
				if gearRatios[Coord{end, y}] == 0 {
					gearRatios[Coord{end, y}] = 1
				}
				gearRatios[Coord{end, y}] *= number
			}

			x += end - start
		}
	}

	for coord, count := range adjCount {
		if count == 2 {
			result += gearRatios[coord]
		}
	}

	return
}

func getNumber(line []string, start int) (end int, result int) {
	numStr := line[start]
	i := start + 1
	for i < len(line) {
		_, err := strconv.Atoi(line[i])

		if err != nil {
			end = i
			break
		}

		numStr += line[i]
		i++
	}

	return i, ToInt(numStr)
}
