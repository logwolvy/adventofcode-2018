package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var filePath = "day02/input.txt"
var part2 = true // flag only

func main() {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	part1Result := []int{0, 0} // 0 holds pair count, 1 holds triplet count

	scanner := bufio.NewScanner(file)
	input := make([]string, 0)
	for scanner.Scan() {
		input = append(input, strings.TrimSpace(scanner.Text()))
	}

	inputCount := len(input)
outer:
	for i, line := range input {
		// Part 2 code
		if part2 {
			for _, l := range input[i+1 : inputCount-1] {
				if checkMatch(l, line) {
					break outer
				}
			}
		}
		//////////////

		charCounts := make(map[string]int)
		for _, c := range strings.Split(line, "") {
			charCounts[c]++
		}

		hasPair := 0
		hasTriplet := 0
		for _, v := range charCounts {
			if v == 2 {
				hasPair = 1
			} else if v == 3 {
				hasTriplet = 1
			} else if hasPair == 1 && hasTriplet == 1 {
				break
			}
		}
		part1Result[0] += hasPair
		part1Result[1] += hasTriplet
	}
	fmt.Println(part1Result[0] * part1Result[1])
}

// Used for part 2 string matching
func checkMatch(source string, target string) bool {
	var match []string
	s := strings.Split(source, "")
	t := strings.Split(target, "")
	for i, j := range s {
		if j == t[i] {
			match = append(match, j)
		}
	}
	if len(s)-len(match) == 1 {
		fmt.Println(strings.Join(match, ""))
		return true
	}
	return false
}
