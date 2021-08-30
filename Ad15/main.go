package main

import (
	"Ad15/intcomp"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

var program = readFile()
var comp = intcomp.Computer{program, false, 0, 0}
var found bool = false
var steps int = 0
var commands []int

//1 is north, 2 is south, 3 is west, 4 is east

func main() {
	//remoteControlProgram()
	//fmt.Println(steps)
	fmt.Println(comp.Compute(1))
	fmt.Println(comp.Compute(1))
	fmt.Println(comp.Compute(1))
	fmt.Println(comp.Compute(1))
	fmt.Println(comp.Compute(3))
	fmt.Println(comp.Compute(3))
	fmt.Println(comp.Compute(3))
	fmt.Println(comp.Compute(3))
	fmt.Println(comp.Compute(3))
	fmt.Println(comp.Compute(3))
	fmt.Println(comp.Compute(1))
	fmt.Println(comp.Compute(1))
	fmt.Println(comp.Compute(1))
	fmt.Println(comp.Compute(1))
	fmt.Println(comp.Compute(3))
	fmt.Println(comp.Compute(3))
	fmt.Println(comp.Compute(3))
	fmt.Println(comp.Compute(3))
	fmt.Println(comp.Compute(2))
	fmt.Println(comp.Compute(2))
	fmt.Println(comp.Compute(2))
	fmt.Println(comp.Compute(2))
	fmt.Println(comp.Compute(4))
	fmt.Println(comp.Compute(4))
	fmt.Println(comp.Compute(1))
	fmt.Println(comp.Compute(1))
	fmt.Println(comp.Compute(2))
	fmt.Println(comp.Compute(1))
	fmt.Println(comp.Compute(1))
}

func remoteControlProgram() {

	fmt.Println(commands)
	if found {
		return
	}

	for i := 1; i <= 4; i++ {

		if len(commands) != 0 && !checkPreviousMove(i) { //if move was performed, certain moves cannot be performed to avoid looping
			continue
		}

		output := comp.Compute(i)
		//fmt.Println(output)

		if output == 0 { //crash to wall, check further position
			continue
		} else if output == 1 { //possible path
			steps++
			commands = append(commands, i)
			remoteControlProgram()

		} else if output == 2 {
			found = true
			return
		}
	}

	//when it comes here, all possibilties are checked and no move is allowed. we go back one step, remove step from commands, and steps--
	stepBack()
	remove(commands, steps-1)
	steps--

}

func checkPreviousMove(move int) bool { //to check that we are not moving back intentionally
	if commands[steps-1] == 1 && move == 2 {
		return false
	} else if commands[steps-1] == 2 && move == 1 {
		return false
	} else if commands[steps-1] == 3 && move == 4 {
		return false
	} else if commands[steps-1] == 4 && move == 3 {
		return false
	}
	return true
}

func stepBack() {
	if commands[len(commands)-1] == 1 {
		comp.Compute(2)
	} else if commands[len(commands)-1] == 2 {
		comp.Compute(1)
	} else if commands[len(commands)-1] == 3 {
		comp.Compute(4)
	} else if commands[len(commands)-1] == 4 {
		comp.Compute(3)
	}
}

func remove(slice []int, s int) []int {
	return slice[:len(slice)-1]
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

	for len(nums) < 2000000 {
		nums = append(nums, 0)
	}

	return nums
}
