package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isSafe(levels []int) bool {
	if len(levels) < 2 {
		return true
	}

	if levels[0] == levels[1] {
		return false
	}

	ascending := true

	if levels[0] > levels[1] {
		ascending = false
	}

	for i := 1; i < len(levels); i++ {
		diff := levels[i] - levels[i-1]
		if !ascending {
			diff *= -1
		}

		if diff > 3 || diff < 1 {
			return false
		}
	}
	return true
}

func isPartiallySafe(levels []int) bool {
	if len(levels) < 2 {
		return true
	}

	return isPartiallySafeHelper(levels, true) || isPartiallySafeHelper(levels, false)
}
func isPartiallySafeHelper(levels []int, ascending bool) bool {

	for i := 0; i < len(levels); i++ {
		prevIdx := 0

		safe := false

		for j := 1; j < len(levels); j++ {
			if prevIdx == i {
				prevIdx = j
				continue
			}

			if j == i {
				continue
			}
			diff := levels[j] - levels[prevIdx]
			prevIdx = j

			if !ascending {
				diff *= -1
			}

			if diff > 3 || diff < 1 {
				break
			}
			if j == len(levels)-1 || (i == len(levels)-1 && j == len(levels)-2) {

				safe = true
			}
		}

		if safe {
			return true
		}
	}
	return false
}

func Day2() {
	file, err := os.Open("inputs/day2input.txt")

	defer file.Close()

	if err != nil {
		panic("Input file for day 2 could not be found")
	}

	scanner := bufio.NewScanner(file)

	safeReports := 0
	partiallySafeReports := 0

	for scanner.Scan() {
		content := scanner.Text()
		var levels []int
		for _, level := range strings.Split(content, " ") {
			iLevel, err := strconv.Atoi(level)
			if err != nil {
				panic("Invalid input, only numbers are allowed")
			}

			levels = append(levels, iLevel)
		}

		if isSafe(levels) {
			safeReports += 1
		}

		if isPartiallySafe(levels) {
			partiallySafeReports += 1
		}
	}

	fmt.Println("Safe reports:", safeReports)
	fmt.Println("Partially safe reports", partiallySafeReports)

}
