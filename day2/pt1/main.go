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

	gameIdSum := 0

	for fileScanner.Scan() {
		gameIdSum += getGameIDIfPossible(fileScanner.Text())
	}

	calibrationDataFile.Close()

	fmt.Printf("Sum of possible game IDs: %d\n", gameIdSum)
}

var possibleGameCubeValues = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func getGameIDIfPossible(line string) int {
	id, isValidGame := parseGame(line)
	if isValidGame {
		return id
	}

	return 0
}

func parseGame(line string) (int, bool) {
	lineSplitOnColon := strings.Split(line, ":")
	gameID, err := strconv.Atoi(strings.Split(lineSplitOnColon[0], " ")[1])
	if err != nil {
		fmt.Println(err)
	}

	return gameID, isValidGame(lineSplitOnColon[1])
}

func isValidGame(game string) bool {
	validGame := true
	gameSplitOnSemi := strings.Split(game, ";")
	for _, pull := range gameSplitOnSemi {
		pullColors := strings.Split(pull, ",")
		for _, gameColor := range pullColors {
			gameColor = strings.TrimSpace(gameColor)
			value, color := strings.Split(gameColor, " ")[0], strings.Split(gameColor, " ")[1]
			valueInt, err := strconv.Atoi(value)
			if err != nil {
				fmt.Println(err)
			}
			if valueInt > possibleGameCubeValues[color] {
				validGame = false
			}
		}
	}
	return validGame
}
