package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
)

// Define structs for line numbers, symbol indices, and part numbers.
type LineNum struct {
	value int
}

type SymbolIdx struct {
	index int
}

type PartNumber struct {
	value         int
	previousIndex int
	nextIndex     int
}

// Global maps to store symbols' indices and part numbers.
var linesSymbolsIdxs = make(map[LineNum][]SymbolIdx)
var partNumbers = make(map[LineNum][]PartNumber)
var sum int

// Part1 reads a file and calculates a sum based on certain conditions.
func Part1() int {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNum := 0

	for scanner.Scan() {
		line := scanner.Text()
		lineNum++

		// Regular expressions to identify symbols and numbers.
		isSymbol := regexp.MustCompile(`[^\w\s.]`)
		isNumber := regexp.MustCompile(`\d+`)

		lineNumTyped := LineNum{lineNum}
		println("---" + strconv.Itoa(lineNum) + "---")
		symbolIdxPairs := isSymbol.FindAllStringIndex(line, -1)
		var symbolIdxs []SymbolIdx
		for _, symbolIdxsPair := range symbolIdxPairs {
			symbolIdxs = append(symbolIdxs, SymbolIdx{symbolIdxsPair[0]})
		}
		// linesSymbolsIdxs[139] = [23,34,51,65,76,79,84,91,94,104,110,133]
		linesSymbolsIdxs[lineNumTyped] = append(linesSymbolsIdxs[lineNumTyped], symbolIdxs...)

		numPairs := isNumber.FindAllString(line, -1)
		var numPairsInt []int
		for _, numPair := range numPairs {
			println(numPair)
			asInt, _ := strconv.Atoi(numPair)
			numPairsInt = append(numPairsInt, asInt)
		}

		// Getting indexes of numbers.
		numIdxPairs := isNumber.FindAllStringIndex(line, -1)

		// Pair numbers with their neighboring indices.
		for index, numIndePair := range numIdxPairs {
			previousNeighbour := 0
			if numIndePair[0] > 0 {
				previousNeighbour = numIndePair[0] - 1
			}
			// partNumbers[139] = {value: 23, previousIndex: 22, nextIndex: 24}
			partNumbers[lineNumTyped] = append(partNumbers[lineNumTyped], PartNumber{numPairsInt[index], previousNeighbour, numIndePair[1]})
		}
	}

	// Process each part number and calculate the sum.
	for lineNumber, partNumbers := range partNumbers {
		for _, partNumber := range partNumbers {
			shouldBeCounted := false

			// Retrieve symbol indices from adjacent lines.
			previousLineSymbolIndexes := linesSymbolsIdxs[LineNum{lineNumber.value - 1}]
			thisLineSymbolIndexes := linesSymbolsIdxs[LineNum{lineNumber.value}]
			nextLineSymbolIndexes := linesSymbolsIdxs[LineNum{lineNumber.value + 1}]

			partNumberMinIndex := partNumber.previousIndex
			partNumberMaxIndex := partNumber.nextIndex

			// Check if the number is adjacent to any symbol on the same or neighboring lines.
			for _, previousLineSymbolIndex := range previousLineSymbolIndexes {
				if previousLineSymbolIndex.index >= partNumberMinIndex && previousLineSymbolIndex.index <= partNumberMaxIndex {
					shouldBeCounted = true
					break
				}
			}

			for _, thisLineSymbolIndex := range thisLineSymbolIndexes {
				if thisLineSymbolIndex.index == partNumberMinIndex || thisLineSymbolIndex.index == partNumberMaxIndex {
					shouldBeCounted = true
					break
				}

			}

			for _, nextLineSymbolIndex := range nextLineSymbolIndexes {
				if nextLineSymbolIndex.index >= partNumberMinIndex && nextLineSymbolIndex.index <= partNumberMaxIndex {
					shouldBeCounted = true
					break
				}
			}

			// Add the number's value to sum if it meets the condition.
			if shouldBeCounted {
				sum += partNumber.value
			}
		}
	}

	return sum
}
