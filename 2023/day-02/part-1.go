package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

var MaxRed, MaxGreen, MaxBlue = 12, 13, 14

func Part1() int {
	sum := 0

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	validGames := make([]int, 0)

	for scanner.Scan() {
		line := scanner.Text()
		game := strings.Split(line, ":")[0]
		id := strings.Split(game, " ")[1]
		cubes := strings.Split(line, ":")[1]
		handfuls := strings.Split(cubes, ";")
		validHandfuls := make([]bool, 0)

		for _, handful := range handfuls {
			colors := strings.Split(handful, ",")
			keyValuePairs := make(map[string]int)

			for _, color := range colors {
				trimmed := strings.TrimLeft(color, " ")
				i := strings.Split(trimmed, " ")[0]
				c := strings.Split(trimmed, " ")[1]
				num, err := strconv.Atoi(i)
				if err != nil {
					log.Fatal(err)
					break
				}
				keyValuePairs[c] = num
			}
			if err != nil {
				log.Fatal(err)
				break
			}
			if keyValuePairs["red"] <= MaxRed &&
				keyValuePairs["green"] <= MaxGreen &&
				keyValuePairs["blue"] <= MaxBlue {
				validHandfuls = append(validHandfuls, true)
			}
		}
		if len(validHandfuls) == len(handfuls) {
			numID, err := strconv.Atoi(id)
			if err != nil {
				log.Fatal(err)
				break
			}
			// println("valid game: " + id)
			validGames = append(validGames, numID)
		}
	}
	for _, id := range validGames {
		sum += id
	}
	return sum
}
