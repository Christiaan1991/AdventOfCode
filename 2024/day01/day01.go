package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input := readFile("input.txt")

	fmt.Printf("The answer to part 1 is: %d\n", part1(input))
	fmt.Printf("The answer to part 2 is: %d\n", part2(input))
}

func part1(input []string) int {
	var leftList []int
	var rightList []int
	for _, s := range input {
		ss := strings.Split(s, "   ")
		leftNum, err := strconv.Atoi(ss[0])
		if err != nil {
			panic(err)
		}
		leftList = append(leftList, leftNum)

		rightNum, err := strconv.Atoi(ss[1])
		if err != nil {
			panic(err)
		}
		rightList = append(rightList, rightNum)
	}

	sort.Ints(leftList)
	sort.Ints(rightList)

	ans := 0

	for i := range input {
		diff := int(math.Abs(float64(leftList[i] - rightList[i])))
		ans += diff
	}

	return ans
}

func part2(input []string) int {
	var leftList []int
	rightMap := make(map[int]int)
	for _, s := range input {
		ss := strings.Split(s, "   ")
		leftNum, err := strconv.Atoi(ss[0])
		if err != nil {
			panic(err)
		}
		leftList = append(leftList, leftNum)

		rightNum, err := strconv.Atoi(ss[1])
		if err != nil {
			panic(err)
		}
		rightMap[rightNum]++
	}

	ans := 0

	for _, num := range leftList {
		ans += num * rightMap[num]
	}

	return ans
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
