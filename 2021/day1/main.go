package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := readFile("input")
	part1(input)
	part2(input)
}

func part2(input []int) {

}

func part1(input []int) {
	counter := 0
	var newinput []int
	i := 0

	for i+2 < len(input) {
		newinput = append(newinput, input[i] + input[i+1] + input[i+2])
		i++
	}

	for i := 1; i < len(newinput); i++ {
		if newinput[i] > newinput[i-1] {
			counter++
		}
	}
	fmt.Println(counter)
}

func readFile(filename string) []int {
	file, err := os.Open(filename)
	var nums []int

	if err != nil {
		log.Fatal(err)
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	s := strings.Split(string(data), "\n")

	for _, i := range s {
		j, err := strconv.Atoi(i)

		if err != nil {
			panic(err)
		}

		nums = append(nums, j)
	}

	return nums
}