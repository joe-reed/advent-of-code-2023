package main

import (
	"regexp"
	"strings"

	. "utils"
)

func puzzle1(file string) int {
	strings := strings.Split(file, "\n")
	re := regexp.MustCompile("[0-9]")

	result := 0
	for _, str := range strings {
		matches := re.FindAllString(str, -1)
		result += ConvertToInt(matches[0] + matches[len(matches)-1])
	}
	return result
}

func puzzle2(file string) int {
	strings := strings.Split(file, "\n")
	re := regexp.MustCompile("one|two|three|four|five|six|seven|eight|nine|[0-9]")
	reverseRe := regexp.MustCompile("eno|owt|eerht|ruof|evif|xis|neves|thgie|enin|[0-9]")

	result := 0
	for _, str := range strings {
		first := re.FindString(str)
		last := reverseString(reverseRe.FindString(reverseString(str)))
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
		result += ConvertToInt(strMap[first] + strMap[last])
	}
	return result
}

func reverseString(input string) string {
	runes := []rune(input)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}
