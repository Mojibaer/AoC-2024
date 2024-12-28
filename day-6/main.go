package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.ReadFile("test-data.txt")

	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(file), "\n")

	guardsPos := findGuard(lines)
	locatedSpots := make(map[string]byte)
	direction := lines[guardsPos[0]][guardsPos[1]]

	locatedPositions := getGuardsRoute(lines, guardsPos[1], guardsPos[0], locatedSpots, direction)

	fmt.Println(locatedPositions)

}

func findGuard(lines []string) []int {
	guardPos := []int{}
	for i := 0; i <= len(lines)-1; i++ {
		line := lines[i]

		for j := 0; j <= len(line)-1; j++ {
			if line[j] == '^' || line[j] == 'v' || line[j] == '<' || line[j] == '>' {
				// Register the lines index, as Y (vertical)
				guardPos = append(guardPos, i)
				// Register the line index, as X (horizontal)
				guardPos = append(guardPos, j)

				return guardPos
			}
		}
	}

	return guardPos
}

// Get recursive the located spots.
func getGuardsRoute(lines []string, x int, y int, locatedSpots map[string]byte, direction byte) map[string]byte {
	line := lines[y]
	locatedSpot := strconv.Itoa(y) + strconv.Itoa(x)
	locatedSpots[locatedSpot] = 'X'

	// Check if guard reached the top side of the map.
	if y == 0 && direction == '^' {
		return locatedSpots
	}

	// Check if guard reached the right side of the map.
	if len(line)-1 == x && direction == '>' {
		return locatedSpots
	}

	// Check if guard reached the bottom of the map.
	if len(lines)-1 == y && direction == 'v' {
		return locatedSpots
	}

	// Check if guard reached the left side of the map.
	if x == 0 && direction == '<' {
		return locatedSpots
	}

	// Handle the top move.
	if direction == '^' && lines[y-1][x] != '#' {
		y--
	} else if direction == '^' && lines[y-1][x] == '#' {
		direction = '>'
	}

	// Handle the right move.
	if direction == '>' && line[x+1] != '#' {
		x++
	} else if direction == '>' && lines[y][x+1] == '#' {
		direction = 'v'
	}

	// Handle the bottom move.
	if direction == 'v' && lines[y+1][x] != '#' {
		y++
	} else if direction == 'v' && lines[y+1][x] == '#' {
		direction = '<'
	}

	// Handle the left move.
	if direction == '<' && lines[y][x-1] != '#' {
		x--
	} else if direction == '<' && lines[y][x-1] == '#' {
		direction = '^'
	}

	return getGuardsRoute(lines, x, y, locatedSpots, direction)
}
