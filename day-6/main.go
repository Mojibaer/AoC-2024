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

	guardsPos := findGuard(lines)
	locatedSpots := make(map[string]byte)
	direction := lines[guardsPos[0]][guardsPos[1]]

	locatedPositions := getGuardsRoute(lines, guardsPos[1], guardsPos[0], locatedSpots, direction)

	fmt.Println(len(locatedPositions))

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
	locatedSpot := strconv.Itoa(y) + "," + strconv.Itoa(x)
	locatedSpots[locatedSpot] = 'X'

	// Check if guard reached the top side of the map.
	if y == -1 && direction == '^' {
		locatedSpots[locatedSpot] = 'X'
		return locatedSpots
	}

	// Check if guard reached the right side of the map.
	if len(line)-1 == x && direction == '>' {
		locatedSpots[locatedSpot] = 'X'
		return locatedSpots
	}

	// Check if guard reached the bottom of the map.
	if len(lines)-1 == y && direction == 'v' {
		locatedSpots[locatedSpot] = 'X'
		return locatedSpots
	}

	// Check if guard reached the left side of the map.
	if x == -1 && direction == '<' {
		locatedSpots[locatedSpot] = 'X'
		return locatedSpots
	}

	// Handle the top move.
	if direction == '^' && lines[y-1][x] != '#' {
		locatedSpots[locatedSpot] = 'X'
		y--
	} else if direction == '^' && lines[y-1][x] == '#' {
		locatedSpots[locatedSpot] = 'X'
		direction = '>'
	}

	// Handle the right move.
	if direction == '>' && lines[y][x+1] != '#' {
		locatedSpots[locatedSpot] = 'X'
		x++
	} else if direction == '>' && lines[y][x+1] == '#' {
		locatedSpots[locatedSpot] = 'X'
		direction = 'v'
	}

	// Handle the bottom move.
	if direction == 'v' && lines[y+1][x] != '#' {
		locatedSpots[locatedSpot] = 'X'
		y++
	} else if direction == 'v' && lines[y+1][x] == '#' {
		locatedSpots[locatedSpot] = 'X'
		direction = '<'
	}

	// Handle the left move.
	if direction == '<' && lines[y][x-1] != '#' {
		locatedSpots[locatedSpot] = 'X'
		x--
	} else if direction == '<' && lines[y][x-1] == '#' {
		locatedSpots[locatedSpot] = 'X'
		direction = '^'
	}

	return getGuardsRoute(lines, x, y, locatedSpots, direction)
}
