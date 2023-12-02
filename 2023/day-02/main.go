package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	println("Part 1:", Part1())
	println("Part 2:", Part2())

	var idSum, powerSum int

	// Open the txt file
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Read line by line
	for scanner.Scan() {

		line := scanner.Text()

		// Split the line using whitespace, lines = games
		valid, id, power := checkGame(strings.Fields(line))

		if valid {
			idSum += id
		}
		powerSum += power

	}

	fmt.Println("Sum:", idSum)
	fmt.Println("Sum:", powerSum)
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
}
