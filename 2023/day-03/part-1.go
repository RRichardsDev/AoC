package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
)

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

var linesSymbolsIdxs = make(map[LineNum][]SymbolIdx)
var partNumbers = make(map[LineNum][]PartNumber)
var sum int

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

		isSymbol := regexp.MustCompile(`[^\w\s.]`)
		isNumber := regexp.MustCompile(`\d+`)

		lineNumTyped := LineNum{lineNum}
		symbolIdxPairs := isSymbol.FindAllStringIndex(line, -1)
		var symbolIdxs []SymbolIdx
		for _, symbolIdxsPair := range symbolIdxPairs {
			symbolIdxs = append(symbolIdxs, SymbolIdx{symbolIdxsPair[0]})
		}
		linesSymbolsIdxs[lineNumTyped] = append(linesSymbolsIdxs[lineNumTyped], symbolIdxs...)

		numPairs := isNumber.FindAllString(line, -1)
		var numPairsInt []int
		for _, numPair := range numPairs {
			asInt, _ := strconv.Atoi(numPair)
			numPairsInt = append(numPairsInt, asInt)
		}

		// indexes of numbers
		numIdxPairs := isNumber.FindAllStringIndex(line, -1)

		// zip to PartNumber - numbers with extended to previous and next indexes
		for index, numIndePair := range numIdxPairs {
			previousNeighbour := 0
			if numIndePair[0] > 0 {
				previousNeighbour = numIndePair[0] - 1
			}
			partNumbers[lineNumTyped] = append(partNumbers[lineNumTyped], PartNumber{numPairsInt[index], previousNeighbour, numIndePair[1]})
		}
	}

	for lineNumber, partNumbers := range partNumbers {
		for _, partNumber := range partNumbers {
			shouldBeCounted := false

			previousLineSymbolIndexes := linesSymbolsIdxs[LineNum{lineNumber.value - 1}]
			thisLineSymbolIndexes := linesSymbolsIdxs[LineNum{lineNumber.value}]
			nextLineSymbolIndexes := linesSymbolsIdxs[LineNum{lineNumber.value + 1}]

			partNumberMinIndex := partNumber.previousIndex
			partNumberMaxIndex := partNumber.nextIndex

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

			if shouldBeCounted {
				sum += partNumber.value
			}
		}
	}

	return sum
}
