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
	part1()
	part2()
}

func part2() {
	var fish = readFile2("input")
	for i:=0; i < 256; i++ {
		fish = step(fish)
	}
	fmt.Println(fish)
	var count = 0
	for _, i := range fish {
		count+=i
	}
	fmt.Println(count)
}

func step(fish []int) []int  {
	var next = make([]int, 9)
	for i:=1; i < 9; i++ {
		next[i-1] = fish[i]
	}
	next[6] += fish[0]
	next[8] += fish[0]
	return next
}

func part1() {
	const days int = 80
	var day = 1

	for day <= days {
		checkForNewFish()
		day++
	}
	fmt.Println(len(input))
}

func checkForNewFish() {
	num := len(input) //we do not want to change length of input in for loop
	for i:=0; i < num; i++ {
		if input[i] == 0 {
			input[i] = 6 //change to 6
			input = append(input, 8) //add 8
			continue
		}
		input[i]-- //else, we reduce number by 1
	}
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

	for i:=0; i < len(s); i++ {
		num, _ := strconv.Atoi(s[i])
		nums = append(nums, num)
	}
	return nums
}

func readFile2(filename string) []int {
	file, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	s := strings.Split(string(data), ",")
	var fish = make ([]int, 9)

	for i:=0; i < len(s); i++ {
		out, _ := strconv.Atoi(s[i])
		fish[out]++
	}
	return fish
}