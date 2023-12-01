package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

// 56397

func wordToDigit(line string) string {
	wordsToDigits := map[string]string{
		"one": "on1ne", "two": "tw2wo", "three": "thre3hree", "four": "fou4our",
		"five": "fiv5ive", "six": "si6ix", "seven": "seve7even", "eight": "eigh8ight", "nine": "nin9ine",
	}
	re := regexp.MustCompile(`(?:one|two|three|four|five|six|seven|eight|nine)`)
	return re.ReplaceAllStringFunc(line, func(w string) string {
		if digit, ok := wordsToDigits[w]; ok {
			return digit
		}
		return w
	})
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	totalSum := 0

	for scanner.Scan() {
		line := scanner.Text()

		println("line:", line)
		// Replace spelled out numbers with digits
		line = wordToDigit(line)
		println("mod line:", line)
		line = wordToDigit(line)
		line = wordToDigit(line)
		line = wordToDigit(line)
		line = wordToDigit(line)
		line = wordToDigit(line)
		line = wordToDigit(line)
		line = wordToDigit(line)
		line = wordToDigit(line)
		line = wordToDigit(line)
		line = wordToDigit(line)
		line = wordToDigit(line)
		line = wordToDigit(line)
		line = wordToDigit(line)
		line = wordToDigit(line)
		line = wordToDigit(line)
		line = wordToDigit(line)
		line = wordToDigit(line)
		line = wordToDigit(line)
		line = wordToDigit(line)
		line = wordToDigit(line)
		println("fin line:", line)

		// Find the first and last digit in the line
		digits := regexp.MustCompile(`\d`).FindAllString(line, -1)
		if len(digits) > 0 {
			println("Digits 0:", digits[0])
			println("Digits 1:", digits[len(digits)-1])

			calibrationValue, _ := strconv.Atoi(digits[0] + digits[len(digits)-1])
			// println("Calibration value:", calibrationValue)
			totalSum += calibrationValue
			// println("Total sum:", totalSum)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from file:", err)
	}

	fmt.Println("Total sum of calibration values:", totalSum)
}
