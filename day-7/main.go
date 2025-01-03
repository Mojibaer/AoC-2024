package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.ReadFile("data.txt")

	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(file), "\n")
	result := 0

	for _, line := range lines {
		equationRes := getEquationResult(line)

		numbersAsString := strings.Split(getTheNumbersAsString(line), " ")

		numbers := make([]int, len(numbersAsString))

		// Convert numbers from string to int
		for i, value := range numbersAsString {
			num, _ := strconv.Atoi(value)
			numbers[i] = num
		}

		combinations := getCombinations(numbersAsString)

		if checkResult(numbers, combinations, equationRes) {
			// fmt.Println("True")
			result += equationRes
		} else {
			// fmt.Println("False")
		}

		// fmt.Println(line)
		// fmt.Println(equationRes, ": numbers length=", len(numbers))
		// fmt.Println(combinations)
	}

	fmt.Println(result)
}

func getEquationResult(line string) int {
	result := ""

	for _, value := range line {
		if value == ':' {
			break
		}

		result += string(value)
	}

	numRes, _ := strconv.Atoi(result)
	return numRes
}

func getTheNumbersAsString(line string) string {
	number := strings.Split(line, ":")
	return strings.TrimSpace(number[1])
}

func getCombinations(numbers []string) int {
	result := 1
	for i := 1; i <= len(numbers)-1; i++ {
		result *= 2
	}

	return result
}

func checkResult(numbers []int, combinations int, equationResult int) bool {
	operaterCount := len(numbers) - 1
	// fmt.Println("-------------------------------------")
	// fmt.Println("combinations: ", combinations)
	// fmt.Println("numbers: ", numbers)

	for i := 0; i <= combinations-1; i++ {

		result := 0
		binary := fmt.Sprintf("%0*s", operaterCount, strconv.FormatInt(int64(i), 2))
		// fmt.Println("Binary: ", binary)

		if len(numbers) < 3 {
			sum := numbers[0] + numbers[1]
			product := numbers[0] * numbers[1]

			return equationResult == sum || equationResult == product
		}

		// fmt.Println("index: ", i)
		// fmt.Println("binary length: ", len(binary))
		// fmt.Println("numbers length: ", len(numbers))
		// fmt.Println(binary)
		// Handle multiplication
		for j, value := range binary {
			// fmt.Println("pop: ", result)

			if value == '1' {
				if j == 0 {
					result = numbers[j] * numbers[j+1]
				} else {
					result *= numbers[j+1]
				}
			} else {
				if j == 0 {
					result = numbers[j] + numbers[j+1]
				} else {
					result += numbers[j+1]
				}
			}
		}

		// fmt.Println("Binary: ", binary)
		// fmt.Println("result: ", result, "equationResult: ", equationResult)
		// fmt.Println("-------------------------------------")

		if result == equationResult {
			return true
		}

		result = 0
	}

	return false
}
