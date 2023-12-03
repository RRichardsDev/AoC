package main

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"strconv"
	"unicode"
)

type Coord struct {
	row int
	col int
}

type Part struct {
	number  int
	engines []Coord
}

func getCoordsToCheck(rowIndex, from, to int, matrix [][]rune) []Coord {
	maxRow, maxCol := len(matrix)-1, len(matrix[0])-1
	var initialList []Coord
	// above and botton row
	for col := from - 1; col <= to+1; col++ {
		initialList = append(
			initialList,
			Coord{row: rowIndex - 1, col: col},
			Coord{row: rowIndex + 1, col: col},
		)
	}
	initialList = append(initialList, Coord{row: rowIndex, col: from - 1}, Coord{row: rowIndex, col: to + 1})
	var cleanList []Coord
	for _, coord := range initialList {
		row, col := coord.row, coord.col
		if row >= 0 && row <= maxRow && col >= 0 && col <= maxCol {
			cleanList = append(cleanList, coord)
		}
	}
	return cleanList
}
func parseInput() [][]rune {
	// lines := common.ReadFile("day3/input_small")
	lines := ReadFile("input.txt")
	var matrix [][]rune
	for _, line := range lines {
		var row []rune
		for _, char := range line {
			row = append(row, char)
		}
		matrix = append(matrix, row)
	}
	return matrix
}

func getEnginesAround(rowIndex, from, to int, matrix [][]rune) []Coord {
	var engines []Coord
	coordToCheck := getCoordsToCheck(rowIndex, from, to, matrix)
	for _, coord := range coordToCheck {
		char := matrix[coord.row][coord.col]
		if char == '*' {
			engines = append(engines, coord)
		}
	}
	return engines
}
func ErrPanic(err error) {
	if err != nil {
		panic(err)
	}
}

func ReadFile(path string) []string {
	content, err := ioutil.ReadFile(path)
	ErrPanic(err)
	scanner := bufio.NewScanner(bytes.NewReader(content))
	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	return lines
}

func getEngineParts(parts []Part, engine Coord) []Part {
	var engineParts []Part
	for _, part := range parts {
		for _, e := range part.engines {
			if e.row == engine.row && e.col == engine.col {
				engineParts = append(engineParts, part)
			}
		}
	}
	return engineParts
}

func calculateGearsRatio(parts [][]Part, engines [][]Coord, matrix [][]rune) int {
	var sum int
	for i, row := range engines {
		for _, engine := range row {
			var engineParts []Part
			if i > 0 {
				engineParts = append(engineParts, getEngineParts(parts[i-1], engine)...)
			}
			engineParts = append(engineParts, getEngineParts(parts[i], engine)...)
			if i < (len(engines) - 1) {
				engineParts = append(engineParts, getEngineParts(parts[i+1], engine)...)
			}
			if len(engineParts) == 2 {
				sum += (engineParts[0].number * engineParts[1].number)
			}
		}
	}

	return sum
}
func getPartsAndEngines(matrix [][]rune) ([][]Part, [][]Coord) {
	parts := make([][]Part, len(matrix))
	engines := make([][]Coord, len(matrix))
	var start, end int
	var number string
	var parsingNumber bool
	for i, row := range matrix {
		start, end = 0, 0
		number, parsingNumber = "", false
		for j, char := range row {
			if char == '*' {
				engines[i] = append(engines[i], Coord{row: i, col: j})
			}
			if unicode.IsDigit(char) {
				number += string(char)
			}
			if unicode.IsDigit(char) && !parsingNumber {
				parsingNumber = true
				start = j
			}
			if !unicode.IsDigit(char) && parsingNumber {
				end = j - 1
				// check
				enginesAround := getEnginesAround(i, start, end, matrix)
				n, err := strconv.Atoi(number)
				ErrPanic(err)
				part := Part{number: n, engines: enginesAround}
				parts[i] = append(parts[i], part)
				number, start, end, parsingNumber = "", 0, 0, false
			}
		}
		if parsingNumber {
			enginesAround := getEnginesAround(i, start, len(row)-1, matrix)
			n, err := strconv.Atoi(number)
			ErrPanic(err)
			part := Part{number: n, engines: enginesAround}
			parts[i] = append(parts[i], part)
		}
	}
	return parts, engines
}
func Part2() int {
	matrix := parseInput()
	parts, engines := getPartsAndEngines(matrix)
	return calculateGearsRatio(parts, engines, matrix)
}
