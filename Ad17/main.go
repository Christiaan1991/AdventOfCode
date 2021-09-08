package main

import (
	"Ad17/intcomp"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

const size int = 60

var image [size][size]int

var program = readFile()
var comp1 = intcomp.Computer{program, false, 0, 0}
var comp2 = intcomp.Computer{program, false, 0, 0}

func main() {
	//part1()
	part2()
}

func part2() {
	program[0] = 2
}

func part1() {
	createImage()
	findIntersections()
	showMap()
}

func findIntersections() {
	var sum int = 0
	for i := 1; i < size; i++ {
		for j := 1; j < size; j++ {
			if image[i][j] == 35 && image[i+1][j] == 35 && image[i-1][j] == 35 && image[i][j+1] == 35 && image[i][j-1] == 35 {
				image[i][j] = 1
				sum += i * j
			}
		}
	}
	fmt.Println(sum)
}

func createImage() {
	var i int = 0
	var j int = 0

	for {
		image[i][j] = comp1.Compute(0)

		if image[i][j] == 10 {
			i++
			j = 0
		} else if image[i][j] == 0 {
			break
		} else {
			j++
		}

	}
}

func showMap() {
	for i := 0; i < len(image); i++ {
		for j := 0; j < len(image[i]); j++ {
			if image[i][j] == 35 {
				fmt.Printf("#")
			} else if image[i][j] == 46 {
				fmt.Printf(".")
			} else if image[i][j] == 1 {
				fmt.Printf("O")
			}
		}
		fmt.Printf("\n")
	}
	//fmt.Printf("\n")
}

func readFile() []int {

	file, err := os.Open("data.txt")
	var nums []int

	if err != nil {
		log.Fatal(err)
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}

	s := strings.Split(string(data), ",")

	for _, i := range s {
		j, err := strconv.Atoi(i)

		if err != nil {
			panic(err)
		}

		nums = append(nums, j)
	}

	for len(nums) < 20000 {
		nums = append(nums, 0)
	}

	return nums
}
