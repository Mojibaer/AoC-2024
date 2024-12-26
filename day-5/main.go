package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	file, err := os.ReadFile("data.txt")

	if err != nil {
		log.Fatal(err)
	}

	data := string(file)
	rules := []string{}
	updatePage := []string{}

	result := 0

	lines := strings.Split(data, "\n")

	isRules := false

	// Store rule and pageUpdate data in slices.
	for _, line := range lines {
		if line == "" {
			isRules = true
			continue
		}

		if isRules {
			updatePage = append(updatePage, line)
		} else {
			rules = append(rules, line)
		}
	}

	for _, line := range updatePage {
		updateLine := strings.Split(line, ",")
		if checkForward(rules, updateLine) {
			// Since the values treated as string, need to convert
			// it to int to sum them.
			midValue, err := strconv.Atoi(getMidValue(updateLine))
			if err != nil {
				panic(err)
			}

			result += midValue
		}
	}

	fmt.Println(result)
}

func checkForward(rules []string, pageUpdate []string) bool {
	for i := 0; i < len(pageUpdate)-1; i++ {
		ruleBreak := pageUpdate[i+1] + pageUpdate[i]

		for _, value := range rules {
			value := strings.Split(value, "|")
			valueStr := value[0] + value[1]
			if valueStr == ruleBreak {
				return false
			}
		}
	}
	return true
}

func getMidValue(pageUpdate []string) string {
	if len(pageUpdate) == 3 {
		return pageUpdate[1]
	}

	index := (len(pageUpdate) - 1) / 2

	return pageUpdate[index]
}
