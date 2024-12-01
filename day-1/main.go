package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("data.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lList []int
	var rList []int

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		// Seprate lines by space
		parts := strings.Fields(scanner.Text())

		if len(parts) >= 2 {
			// Convert Strings to Number
			firstPart, errFp := strconv.Atoi(parts[0])
			secondPart, errSp := strconv.Atoi(parts[1])

			if errFp != nil || errSp != nil {
				fmt.Println("Error produced by converting string to int")
			}

			lList = append(lList, firstPart)
			rList = append(rList, secondPart)
		}
	}

	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Part one: ", partOne(lList, rList))
	fmt.Println("Part two: ", partTwo(lList, rList))

}

func partOne(lList []int, rList []int) int {
	// Sort the Parts from Min to Max
	sort.Ints(lList)
	sort.Ints(rList)

	var distance int

	for i := range lList {
		tmp := lList[i] - rList[i]

		// If tmp is a negative number turn it to positive with bit manipulation
		if tmp < 0 {
			tmp = ^tmp + 1 // Two's complement
		}

		distance += tmp
	}

	return distance
}

func partTwo(lList []int, rList []int) int {
	var similarityScore int

	for _, valueL := range lList {
		tmp := 0

		for _, valueR := range rList {
			if valueL == valueR {
				tmp += 1
			}
		}

		similarityScore += valueL * tmp
	}

	return similarityScore
}
