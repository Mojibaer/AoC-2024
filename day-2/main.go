package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("data.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines [][]int

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		parts := strings.Fields(scanner.Text())
		line := make([]int, 0, len(parts))

		for _, value := range parts {
			num, err := strconv.Atoi(string(value))

			if err != nil {
				log.Fatal("Char to Int conversion failed. Error: ", err)
			}

			line = append(line, num)
		}

		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Scanner thrown an Error: ", err)
	}

	//fmt.Println(partOne(lines))
	fmt.Println("Part Two Result: ", partTwo(lines))
}

func partOne(lines [][]int) int {
	safeCounter := 0

	for _, line := range lines {
		isSafe := true
		increase := true
		decrease := true

		for i := 0; i < len(line)-1; i++ {
			x := line[i] - line[i+1]
			if (x > 0) && ((x) < 4) && decrease {
				increase = false
			} else if (x < 0) && ((x) > -4) && increase {
				decrease = false
			} else {
				isSafe = false
				break
			}
		}

		if isSafe {
			safeCounter++
		}
	}

	return safeCounter
}

func partTwo(lines [][]int) int {
	safeCounter := 0

	for _, line := range lines {
		increase := false
		decrease := false

		// Check first if the line is valid without any bad-levels
		if increaseLine(line) || decreaseLine(line) {
			safeCounter++
			continue
		}

		for i := 0; i < len(line)-1; i++ {
			x := line[i] - line[i+1]

			// Set Increase flag and check if ordering switches.
			if (x > -4) && (x < 0) {
				if decrease {
					if isNewLineValid(line, i) {
						safeCounter++
					}

					break
				}

				increase = true
			}

			// Set Decrease flag and check if ordering switches.
			if (x > 0) && (x < 4) {
				if increase {
					if isNewLineValid(line, i) {
						safeCounter++
					}

					break
				}

				decrease = true
			}

			// Case if x is out of range or zero
			if (x < -3) || (x > 3) || (x == 0) {
				if isNewLineValid(line, i) {
					safeCounter++
				}

				break
			}
		}
	}

	return safeCounter
}

func increaseLine(line []int) bool {

	for i := 0; i < len(line)-1; i++ {
		x := line[i+1] - line[i]

		if !(x > 0 && x < 4) {
			return false
		}
	}

	return true
}

func decreaseLine(line []int) bool {
	for i := 0; i < len(line)-1; i++ {
		x := line[i] - line[i+1]

		if !(x > 0 && x < 4) {
			return false
		}
	}

	return true
}

func getNewLineWithoutCurrentIndex(line []int, index int) []int {
	newLine := make([]int, 0, len(line)-1)

	for i, value := range line {
		if i == index {
			continue
		}

		newLine = append(newLine, value)
	}

	return newLine
}

// Check the line after a bad-level
// by removing the current, after and previous index.
func isNewLineValid(line []int, index int) bool {
	newLine := getNewLineWithoutCurrentIndex(line, index)

	if increaseLine(newLine) || decreaseLine(newLine) {
		return true
	}

	if (index + 1) <= len(line) {
		newLine = getNewLineWithoutCurrentIndex(line, index+1)

		if increaseLine(newLine) || decreaseLine(newLine) {
			return true
		}
	}

	if (index - 1) <= len(line) {
		newLine = getNewLineWithoutCurrentIndex(line, index-1)

		if increaseLine(newLine) || decreaseLine(newLine) {
			return true
		}
	}

	return false
}
