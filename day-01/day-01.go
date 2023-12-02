package main

import (
	"regexp"

	. "utils"
)

func puzzle1(input []string) int {
	re := regexp.MustCompile("[0-9]")

	result := 0
	for _, line := range input {
		matches := re.FindAllString(line, -1)
		result += ToInt(matches[0] + matches[len(matches)-1])
	}
	return result
}

func puzzle2(input []string) int {
	re := regexp.MustCompile("one|two|three|four|five|six|seven|eight|nine|[0-9]")
	reverseRe := regexp.MustCompile("eno|owt|eerht|ruof|evif|xis|neves|thgie|enin|[0-9]")

	result := 0
	for _, line := range input {
		first := re.FindString(line)
		last := reverse(reverseRe.FindString(reverse(line)))
		strMap := map[string]string{
			"1":     "1",
			"2":     "2",
			"3":     "3",
			"4":     "4",
			"5":     "5",
			"6":     "6",
			"7":     "7",
			"8":     "8",
			"9":     "9",
			"one":   "1",
			"two":   "2",
			"three": "3",
			"four":  "4",
			"five":  "5",
			"six":   "6",
			"seven": "7",
			"eight": "8",
			"nine":  "9",
		}
		result += ToInt(strMap[first] + strMap[last])
	}
	return result
}

func reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}
