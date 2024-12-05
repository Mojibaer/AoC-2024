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
	fmt.Println("function two", partTwo(lines))
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
		tolerateCounter := 0

		if increaseLine(line) || decreaseLine(line) {
			safeCounter++
			continue
		}

		for i := 0; i < len(line)-1; i++ {

			x := line[i] - line[i+1]

			// Case if x is out of range
			if (x < -3) || (x > 3) {
				tolerateCounter++

				if tolerateCounter > 1 {
					break
				}

				newLine := getNewLineWithoutCurrentIndex(line, i+1)

				if tolerateCounter < 2 && (increaseLine(newLine) || decreaseLine(newLine)) {
					safeCounter++
				}
			}

			// Case if x is zero
			if x == 0 {
				tolerateCounter++

				newLine := getNewLineWithoutCurrentIndex(line, i)

				if tolerateCounter < 2 && (increaseLine(newLine) || decreaseLine(newLine)) {
					safeCounter++
				}

				break
			}

			// Case if sorting switches from increase to decrease
			if (x > 0) && ((x) < 4) && increase {
				tolerateCounter++

				newLine := getNewLineWithoutCurrentIndex(line, i)

				if tolerateCounter < 2 && (increaseLine(newLine) || decreaseLine(newLine)) {
					safeCounter++
				}

				break
			}

			// Case if sorting switches from decrease to increase
			if (x < 0) && ((x) > -4) && decrease {
				tolerateCounter++

				newLine := getNewLineWithoutCurrentIndex(line, i)

				if tolerateCounter < 2 && (increaseLine(newLine) || decreaseLine(newLine)) {
					safeCounter++
				}

				break
			}

			// TolerateCounter check on increase
			if (x < 0) && ((x) > -4) {
				if decrease {
					tolerateCounter++

					if tolerateCounter > 1 {
						break
					}
				}

				increase = true
				decrease = false
			}

			// TolerateCounter check on decrease
			if (x > 0) && ((x) < 4) {
				if increase {
					tolerateCounter++

					if tolerateCounter > 1 {
						break
					}
				}

				increase = false
				decrease = true
			}

			if i == len(line)-1 && tolerateCounter < 2 {
				safeCounter++
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
