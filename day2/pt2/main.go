package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	calibrationDataFile, err := os.Open("../input.txt")
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(calibrationDataFile)

	fileScanner.Split(bufio.ScanLines)

	gamePowerSum := 0

	for fileScanner.Scan() {
		gamePowerSum += getPowerOfGame(parseGame(fileScanner.Text()))
	}

	calibrationDataFile.Close()

	fmt.Printf("Sum of power of games: %d\n", gamePowerSum)
}

func parseGame(line string) string {
	return strings.Split(line, ":")[1]
}

func getPowerOfGame(line string) int {
	redVal, greenVal, blueVal := getMinimumCubesNeeded(line)
	return redVal * greenVal * blueVal
}

func getMinimumCubesNeeded(game string) (int, int, int) {
	minColorValNeeded := map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}

	gameSplitOnSemi := strings.Split(game, ";")
	for _, pull := range gameSplitOnSemi {
		pullColors := strings.Split(pull, ",")
		for _, gameColor := range pullColors {
			gameColor = strings.TrimSpace(gameColor)
			fmt.Println(gameColor)
			value, color := strings.Split(gameColor, " ")[0], strings.Split(gameColor, " ")[1]
			valueInt, err := strconv.Atoi(value)
			if err != nil {
				fmt.Println(err)
			}
			if minColorValNeeded[color] < valueInt {
				minColorValNeeded[color] = valueInt
			}
		}
	}
	return minColorValNeeded["red"], minColorValNeeded["green"], minColorValNeeded["blue"]
}
