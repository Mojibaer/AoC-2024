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

	for i := 0; i < len(data)-3; i++ {

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

		} else {
			continue
		}
	}

	fmt.Println(result)
}

func multiply(num1 int, num2 int) int {
	return num1 * num2
}
