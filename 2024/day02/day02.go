package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := readFile("input.txt")

	fmt.Printf("The answer to part 1 is: %d\n", part1(input))
	fmt.Printf("The answer to part 2 is: %d\n", part2(input))
}

func part1(input [][]int) int {
	numSafe := 0
	for _, row := range input {
		if decreasing(row) || increasing(row) {
			numSafe++
		}
	}

	return numSafe
}

func part2(input [][]int) int {
	numSafe := 0

	for _, row := range input {
		if decreasingWithPop(row) || increasingWithPop(row) {
			numSafe++
		}
	}

	return numSafe
}

func decreasingWithPop(row []int) bool {
	levelTolerated := false
	for i := 0; i < len(row)-1; i++ {
		if (row[i] <= row[i+1]) || (row[i]-row[i+1]) > 3 {
			if !levelTolerated {
				levelTolerated = true
				continue
			}
			return false
		}
	}
	return true
}

func increasingWithPop(row []int) bool {
	levelTolerated := false
	for i := 0; i < len(row)-1; i++ {
		if (row[i] >= row[i+1]) || (row[i+1]-row[i]) > 3 {
			if !levelTolerated {
				levelTolerated = true
				continue
			}
			return false
		}
	}
	return true
}

func increasing(row []int) bool {
	for i := 0; i < len(row)-1; i++ {
		if (row[i] >= row[i+1]) || (row[i+1]-row[i]) > 3 {
			return false
		}
	}
	return true
}

func decreasing(row []int) bool {
	for i := 0; i < len(row)-1; i++ {
		if (row[i] <= row[i+1]) || (row[i]-row[i+1]) > 3 {
			return false
		}
	}
	return true
}

func readFile(name string) [][]int {
	file, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	var input [][]int

	for fileScanner.Scan() {
		toString := fileScanner.Text()
		ss := strings.Split(toString, " ")
		var array []int
		for _, s := range ss {
			num, err := strconv.Atoi(s)
			if err != nil {
				panic(err)
			}
			array = append(array, num)
		}
		input = append(input, array)
	}
	return input
}
