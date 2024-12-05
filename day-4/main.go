package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.ReadFile("data.txt")

	if err != nil {
		log.Fatal(err)
	}

	var result = 0
	var masResult = 0

	lines := strings.Split(string(file), "\n")

	// Iterate through lines, lines are string of each line
	for i := 0; i < len(lines); i++ {
		line := lines[i]

		// Calculating xmas's for one line forward & backward
		result += checkLineForward(line) + checkLineBackward(line)

		// Handle the Diagonal and Vertical cases
		for j := 0; j < len(line); j++ {
			leftIndicesCheck := j-3 >= 0
			rightIndicesCheck := j+3 < len(line)
			topIndicesCheck := (i - 3) >= 0
			downIndicesCheck := (i + 3) < len(lines)

			if line[j] == 'A' {

				if (i-1) >= 0 && i+1 < len(lines) && (j-1) >= 0 && j+1 < len(line) {
					if (lines[i-1][j-1] == 'M' && lines[i+1][j+1] == 'S') || (lines[i-1][j-1] == 'S' && lines[i+1][j+1] == 'M') {
						if (lines[i-1][j+1] == 'M' && lines[i+1][j-1] == 'S') || (lines[i-1][j+1] == 'S' && lines[i+1][j-1] == 'M') {
							masResult++
						}
					}
				}
			}

			if line[j] == 'X' {
				xmas := []byte("----")

				if topIndicesCheck {
					xmas[0] = line[j]
					xmas[1] = lines[i-1][j]
					xmas[2] = lines[i-2][j]
					xmas[3] = lines[i-3][j]

					if string(xmas) == "XMAS" {
						result++
					}

					xmas = []byte("----")
				}

				if downIndicesCheck {
					xmas[0] = line[j]
					xmas[1] = lines[i+1][j]
					xmas[2] = lines[i+2][j]
					xmas[3] = lines[i+3][j]

					if string(xmas) == "XMAS" {
						result++
					}

					xmas = []byte("----")
				}

				if leftIndicesCheck && topIndicesCheck {
					xmas[0] = line[j]
					xmas[1] = lines[i-1][j-1]
					xmas[2] = lines[i-2][j-2]
					xmas[3] = lines[i-3][j-3]

					if string(xmas) == "XMAS" {
						result++
					}

					xmas = []byte("----")
				}

				if leftIndicesCheck && downIndicesCheck {
					xmas[0] = line[j]
					xmas[1] = lines[i+1][j-1]
					xmas[2] = lines[i+2][j-2]
					xmas[3] = lines[i+3][j-3]

					if string(xmas) == "XMAS" {
						result++
					}

					xmas = []byte("----")

				}

				if rightIndicesCheck && topIndicesCheck {
					xmas[0] = line[j]
					xmas[1] = lines[i-1][j+1]
					xmas[2] = lines[i-2][j+2]
					xmas[3] = lines[i-3][j+3]

					if string(xmas) == "XMAS" {
						result++
					}

					xmas = []byte("----")

				}

				if rightIndicesCheck && downIndicesCheck {
					xmas[0] = line[j]
					xmas[1] = lines[i+1][j+1]
					xmas[2] = lines[i+2][j+2]
					xmas[3] = lines[i+3][j+3]

					if string(xmas) == "XMAS" {
						result++
					}

					xmas = []byte("----")

				}
			}

		}
	}

	fmt.Println("Result: ", result)
	fmt.Println("masResult: ", masResult)

}

func checkLineForward(line string) int {
	xmas := ""
	result := 0
	counter := 0

	for i := 0; i < len(line)-1; i++ {
		xmas = ""
		counter = 0
		if line[i] == 'X' && i+4 <= len(line) {
			for counter < 4 {
				xmas += string(line[i+counter])
				counter++
			}

			if xmas == "XMAS" {
				result++
			}
		}
	}

	return result
}

func checkLineBackward(line string) int {
	xmas := ""
	result := 0
	counter := 0

	for i := 0; i < len(line)-1; i++ {
		xmas = ""
		counter = 0
		if line[i] == 'S' && i+4 <= len(line) {
			for counter < 4 {
				xmas += string(line[i+counter])
				counter++
			}

			if xmas == "SAMX" {
				result++
			}
		}
	}

	return result
}
