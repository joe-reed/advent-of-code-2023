package main

import (
	// "fmt"
	"fmt"
	"testing"
	. "utils"

	"github.com/stretchr/testify/assert"
)

func TestPuzzle1(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{
			"./test-input-1.txt",
			142,
		},
		{
			"./input.txt",
			54968,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, puzzle1(FileToStrings(test.input)))
	}
}

func TestSolvePuzzle1(t *testing.T) {
	fmt.Println("Puzzle 1:", puzzle1(FileToStrings("./input.txt")))
}

func TestPuzzle2(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{
			"./test-input-2.txt",
			281,
		},
		{
			"./input.txt",
			54094,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, puzzle2(FileToStrings(test.input)))
	}
}

func TestSolvePuzzle2(t *testing.T) {
	fmt.Println("Puzzle 2:", puzzle2(FileToStrings("./input.txt")))
}
