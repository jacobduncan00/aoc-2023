package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	calibrationDataFile, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(calibrationDataFile)

	fileScanner.Split(bufio.ScanLines)

	calibrationValueSumPt1 := 0
	calibrationValueSumPt2 := 0

	for fileScanner.Scan() {
		calibrationValueSumPt1 += getCalibrationValue(fileScanner.Text(), true)
		calibrationValueSumPt2 += getCalibrationValue(fileScanner.Text(), false)
	}

	calibrationDataFile.Close()

	fmt.Printf("PT1 Calibration Value: %d\n", calibrationValueSumPt1)
	fmt.Printf("PT2 Calibration Value: %d\n", calibrationValueSumPt2)
}

var stringToIntMap = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func getCalibrationValue(line string, digitsOnly bool) int {
	firstNum := 0
	lastNum := 0

	for idx, char := range line {
		if char >= '0' && char <= '9' {
			if firstNum == 0 {
				firstNum = int(char - '0')
			}
			lastNum = int(char - '0')
			continue
		}
		if !digitsOnly {
			for word, num := range stringToIntMap {
				if strings.HasPrefix(line[idx:], word) {
					if firstNum == 0 {
						firstNum = num
					}
					lastNum = num
				}
			}
		}
	}

	return firstNum*10 + lastNum
}
