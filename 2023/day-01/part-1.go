package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func Part1() int {
	// Open the file
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Create a Scanner to read the file
	scanner := bufio.NewScanner(file)
	outPutArr := []int{} // Corrected declaration

	// Loop through each line
	for scanner.Scan() {
		line := scanner.Text()
		lineArry := []rune(line)

		var startNo, endNo int = -1, -1

		// Find the first number
		for _, char := range lineArry {
			if char >= '0' && char <= '9' {
				startNo = int(char - '0')
				break
			}
		}

		// Find the last number
		for i := len(lineArry) - 1; i >= 0; i-- {
			char := lineArry[i]
			if char >= '0' && char <= '9' {
				endNo = int(char - '0')
				break
			}
		}

		if startNo != -1 && endNo != -1 {
			// Convert numbers to string and concatenate
			lineVal := strconv.Itoa(startNo) + strconv.Itoa(endNo)

			// Convert concatenated string to int
			val, err := strconv.Atoi(lineVal)
			if err == nil {
				outPutArr = append(outPutArr, val) // Corrected append usage
			}
		}
	}

	total := 0
	for _, number := range outPutArr {
		total += number
	}

	// Check for errors during Scan
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return total
}
