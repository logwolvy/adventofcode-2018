package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var filePath = "day01/input.txt"
var result int
var resultsMap = make(map[int]bool) // Used in part 2
var part2 = false                   // Used as a flag

func main() {
outer:
	for {
		file, err := os.Open(filePath) // Probably bad but couldn't avoid opening file in outer loop
		if err != nil {
			panic(err)
		}
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			num, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
			if err != nil {
				num = 0
			}
			result += num
			if resultsMap[result] && part2 {
				fmt.Println(result)
				break outer
			} else {
				resultsMap[result] = true
			}

		}
		if !part2 {
			fmt.Println(result)
			break
		}
		defer file.Close()
	}
}
