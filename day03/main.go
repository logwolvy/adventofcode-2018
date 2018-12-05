package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var input = make([]string, 0)

type coordinate struct {
	x int
	y int
}

type mappedPts map[coordinate]string

func main() {
	processInputFile("day03/input.txt")

	part1Result := 0
	used := make(map[coordinate][]int)
	unTouchedClaims := make(map[string]bool)
	for _, claim := range input {
		claimID, coordinates := inputParser(claim)
		unTouchedClaims[strconv.Itoa(claimID)] = true
		for _, coordinate := range coordinates {
			if len(used[coordinate]) > 0 {
				used[coordinate][0]++
				if used[coordinate][0] == 2 {
					part1Result++
				}
				if used[coordinate][0] > 1 {
					prevClaimID := used[coordinate][1]
					used[coordinate][1] = claimID
					delete(unTouchedClaims, strconv.Itoa(prevClaimID))
					delete(unTouchedClaims, strconv.Itoa(claimID))
				}
			} else {
				used[coordinate] = append(used[coordinate], 1, claimID)
			}
		}
	}
	fmt.Println(part1Result)
	fmt.Println(unTouchedClaims)
}

func processInputFile(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, strings.TrimSpace(scanner.Text()))
	}
}

func inputParser(s string) (int, []coordinate) {
	claimID, _ := strconv.Atoi(strings.Split(strings.Split(s, " ")[0], "#")[1])
	x, _ := strconv.Atoi(strings.Split(strings.Split(s, " ")[2], ",")[0])
	y, _ := strconv.Atoi(strings.Split(strings.Split(strings.Split(s, " ")[2], ",")[1], ":")[0])
	width, _ := strconv.Atoi(strings.Split(strings.Split(s, " ")[3], "x")[0])
	height, _ := strconv.Atoi(strings.Split(strings.Split(s, " ")[3], "x")[1])

	points := make([]coordinate, 0)
	for j := 1; j <= height; j++ {
		yPt := y + j
		for i := 1; i <= width; i++ {
			xPt := x + i
			points = append(points, coordinate{xPt, yPt})
		}
	}
	return claimID, points
}
