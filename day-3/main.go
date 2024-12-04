package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

func main() {

	file, err := os.ReadFile("data.txt")

	if err != nil {
		log.Fatal(err)
	}

	data := string(file)

	var num1 = ""
	var num2 = ""
	var result = 0

	// It cost's me a little time to understand,
	// that the first mul without do and don't is do, LoL
	var isDo = true

	for i := 0; i < len(data)-3; i++ {

		if i+3 < len(data) && checkDo(i, data) {
			isDo = true
		}

		if i+6 < len(data) && checkDont(i, data) {
			isDo = false
		}

		if isDo {
			// Only check for the numbers if "mul(" pattern exist
			if data[i] == 'm' && data[i+1] == 'u' && data[i+2] == 'l' && data[i+3] == '(' {
				i = i + 3
				j := i + 1

				num1 = ""
				num2 = ""
				// Get the first Number
				if unicode.IsDigit(rune(data[j])) {
					for unicode.IsDigit(rune(data[j])) && j < len(data) {
						num1 += string(data[j])
						j++
					}
				}
				fmt.Println(num1)
				// Check the index after first number by comma
				// and get the second number
				if data[j] == ',' && num1 != "" && num2 == "" && j < len(data) {
					j++
					if unicode.IsDigit(rune(data[j])) {
						for unicode.IsDigit(rune(data[j])) && j < len(data) {
							num2 += string(data[j])
							j++
						}
					}
				}
				fmt.Println(num2)
				// Check if the numbers are closed with ')' sum the product to result
				if data[j] == ')' && num1 != "" && num2 != "" {
					p1, err1 := strconv.Atoi(num1)

					if err1 != nil {
						fmt.Println("Strconv error on first number conversion", err1)
					}

					p2, err2 := strconv.Atoi(num2)

					if err2 != nil {
						fmt.Println("Strconv error on second number conversion", err2)
					}

					result += multiply(p1, p2)
				}
			}
		}
	}

	fmt.Println(result)
}

func multiply(num1 int, num2 int) int {
	return num1 * num2
}

func checkDont(i int, data string) bool {
	return data[i] == 'd' && data[i+1] == 'o' && data[i+2] == 'n' &&
		data[i+3] == '\'' && data[i+4] == 't' && data[i+5] == '(' && data[i+6] == ')'
}

func checkDo(i int, data string) bool {
	return data[i] == 'd' && data[i+1] == 'o' && data[i+2] == '(' && data[i+3] == ')'
}
