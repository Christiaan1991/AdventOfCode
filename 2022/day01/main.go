package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {

	nums := readFile("input")
	fmt.Printf("the answer to part 1: %d\n", part1(nums))
	fmt.Printf("the answer to part 2: %d", part2(nums))
}

func part1(input []string) int {
	var total = 0
	var highestTotal = 0
	for _, s := range input {
		if s == "" {
			if total > highestTotal {
				highestTotal = total
			}
			total = 0
			continue
		}
		calorie, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		total += calorie
	}
	return highestTotal
}

func part2(input []string) int {
	var total = 0
	var threeHighestTotals = []int{0, 0, 0}
	for _, s := range input {
		if s == "" {
			for i := 0; i < 3; i++ {
				if total > threeHighestTotals[i] {
					threeHighestTotals[i] = total
					break
				}
			}
			sort.Ints(threeHighestTotals)
			total = 0
			continue
		}
		calorie, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		total += calorie
	}

	var highestTotal = 0
	for _, highest := range threeHighestTotals {
		highestTotal += highest
	}
	return highestTotal
}

func readFile(name string) []string {
	file, err := os.Open(name)

	if err != nil {
		log.Fatal(err)
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		toString := fileScanner.Text()
		fileLines = append(fileLines, toString)
	}
	return fileLines
}
