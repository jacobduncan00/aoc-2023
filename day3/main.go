package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	engineSchematic := readEngineSchematic("./input.txt")

	sumEngineSchematics := getPartSum(engineSchematic)
	sumEngineGearRatios := getGearRatioSum(engineSchematic)

	fmt.Printf("Sum of engine schematics: %d\n", sumEngineSchematics)
	fmt.Printf("Sum of gear ratios: %d\n", sumEngineGearRatios)
}

func readEngineSchematic(fileName string) []string {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	schematic := []string{}
	for fileScanner.Scan() {
		schematic = append(schematic, fileScanner.Text())
	}

	if err := fileScanner.Err(); err != nil {
		fmt.Println(err)
	}

	return schematic
}

func getPartSum(engineSchematic []string) int {
	partNums := []int{}
	for i, line := range engineSchematic {
		adjacent := false
		num := ""
		for j, char := range line {
			if char >= '0' && char <= '9' {
				num += string(char)

				if !adjacent {
					adjacent = checkAdjacency(engineSchematic, i, j)
				}
			}

			if (char < '0' || char > '9') || j == len(line)-1 {
				if adjacent {
					numInt, err := strconv.Atoi(num)
					if err != nil {
						fmt.Println(err)
					}
					partNums = append(partNums, numInt)
				}
				num = ""
				adjacent = false
			}
		}
	}

	sum := 0
	for _, num := range partNums {
		sum += num
	}

	return sum
}

func checkAdjacency(schematic []string, row, col int) bool {
	startRow := row - 1
	endRow := row + 1
	startCol := col - 1
	endCol := col + 1

	if row == 0 {
		startRow = 0
	}

	if row == len(schematic)-1 {
		endRow = len(schematic) - 1
	}

	if col == 0 {
		startCol = 0
	}

	if col == len(schematic[row])-1 {
		endCol = len(schematic[row]) - 1
	}

	for i := startRow; i <= endRow; i++ {
		for j := startCol; j <= endCol; j++ {
			intAsRune := rune(schematic[i][j])
			if schematic[i][j] != '.' && (intAsRune < '0' || intAsRune > '9') {
				return true
			}
		}
	}

	return false
}

func adjacentNums(line string, gearIdx int) (int, []int) {
	var result int
	var adjacentMatches []int

	reNum := regexp.MustCompile(`\d+`)
	numMatches := reNum.FindAllStringIndex(line, -1)

	for _, match := range numMatches {
		if gearIdx >= match[0]-1 && gearIdx <= match[1] {
			result++
			num, _ := strconv.Atoi(line[match[0]:match[1]])
			adjacentMatches = append(adjacentMatches, num)
		}
	}
	return result, adjacentMatches
}

func getGearRatioSum(schematic []string) int {
	gearRatioSum := 0

	reGear := regexp.MustCompile(`\*`)

	for i := range schematic {
		gearMatch := reGear.FindAllStringIndex(schematic[i], -1)

		for _, match := range gearMatch {
			numMatches := 0
			matchingNums := []int{}

			if i > 0 {
				count, nums := adjacentNums(schematic[i-1], match[0])
				numMatches += count
				matchingNums = append(matchingNums, nums...)
			}

			count, nums := adjacentNums(schematic[i], match[0])
			numMatches += count
			matchingNums = append(matchingNums, nums...)

			if i < len(schematic)-1 {
				count, nums := adjacentNums(schematic[i+1], match[0])
				numMatches += count
				matchingNums = append(matchingNums, nums...)
			}

			if numMatches == 2 {
				gearRatioSum += matchingNums[0] * matchingNums[1]
			}
		}
	}

	return gearRatioSum
}
