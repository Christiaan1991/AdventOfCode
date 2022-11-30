package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

var input = readFile("input")

func main() {
	//fmt.Println(input, len(input))
	part1()
	part2()
}

func part2() {
	min, max := MinMax(input)
	var pos int
	minFuel := 100000000000
	for i := min; i <= max; i++ {
		totalFuel := 0
		for j := 0; j < len(input); j++ {
			totalFuel += addFac(Abs(input[j] - i))
		}
		//fmt.Println(i, totalFuel)
		if totalFuel < minFuel {
			minFuel = totalFuel
			pos = i
		}
	}
	fmt.Println("check change for git")

	fmt.Printf("Total minimal fuel costs: %d to pos %d\n", minFuel, pos)
}

func addFac(in int) int {
	var out = 0
	for i := in; i > 0; i-- {
		out += i
	}
	return out
}

func part1() {
	min, max := MinMax(input)
	var pos int
	minFuel := 100000000000
	for i := min; i <= max; i++ {
		totalFuel := 0
		for j := 0; j < len(input); j++ {
			totalFuel += Abs(input[j] - i)
		}
		//fmt.Println(i, totalFuel)
		if totalFuel < minFuel {
			minFuel = totalFuel
			pos = i
		}
	}

	fmt.Printf("Total minimal fuel costs: %d to pos %d\n", minFuel, pos)
}

func MinMax(array []int) (int, int) {
	var max int = array[0]
	var min int = array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func readFile(filename string) []int {
	file, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	s := strings.Split(string(data), ",")
	var nums []int

	for i := 0; i < len(s); i++ {
		num, _ := strconv.Atoi(s[i])
		nums = append(nums, num)
	}
	return nums
}
