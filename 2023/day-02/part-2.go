package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

// 12 red cubes, 13 green cubes, and 14 blue cubes

func Part2() int {
	sum := 0

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	setPowers := make([]int, 0)

	for scanner.Scan() {
		line := scanner.Text()
		// game := strings.Split(line, ":")[0]
		// id := strings.Split(game, " ")[1]
		cubes := strings.Split(line, ":")[1]
		handfuls := strings.Split(cubes, ";")
		gamePowers := make(map[int]map[string]int)

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
				gamePowers[num] = keyValuePairs

			}
			if err != nil {
				log.Fatal(err)
				break
			}
		}
		gamePower := 0
		highestRed := 0
		highestGreen := 0
		highestBlue := 0

		for _, handful := range gamePowers {
			if handful["red"] > highestRed {
				highestRed = handful["red"]
			}
			if handful["green"] > highestGreen {
				highestGreen = handful["green"]
			}
			if handful["blue"] > highestBlue {
				highestBlue = handful["blue"]
			}
		}
		println("highestRed", highestRed)
		println("highestGreen", highestGreen)
		println("highestBlue", highestBlue)
		gamePower = highestRed * highestGreen * highestBlue
		println("---")
		println("gamePower", gamePower)
		println("---")

		setPowers = append(setPowers, gamePower)

	}
	for _, power := range setPowers {
		// println("power", power)
		sum += power
	}
	return sum
}
