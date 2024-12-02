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

	fmt.Println(partOne(lines))
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
