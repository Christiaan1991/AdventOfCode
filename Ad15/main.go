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

const size int = 50

var maps [size][size]int
var xpos = size / 2
var ypos = size / 2

//1 is north, 2 is south, 3 is west, 4 is east

func main() {
	var previous_command int = 0
	maps[xpos][ypos] = 4
	remoteControlProgram(previous_command)
	showMap()
}

func remoteControlProgram(previous_move int) {

	// if steps%size/2 == 0 {
	// 	showMap()
	// }

	if fullMaze() {
		return
	}

	for move := 1; move <= 4; move++ {

		// if found {
		// 	return
		// }

		if oppositeDirection(move) == previous_move { //if move is opposite direction from the previous move, move cannot be performed to avoid looping
			continue
		}

		updatePosition(move)
		output := comp.Compute(move)
		maps[xpos][ypos] = output + 1

		//fmt.Println(xpos, ypos, previous_move, move, output)

		if output == 0 { //crash to wall, check further position
			updatePosition(oppositeDirection(move))
			continue
		} else if output == 1 { //possible path! we add +1 to the total number of steps and go find next path
			steps++
			remoteControlProgram(move)

		} else if output == 2 {
			fmt.Println("found endpoint")
			return
		}
	}

	//when it comes here, all possibilties are checked and no move is allowed except for the opposite direction. we perform move back, calculate  in the intcomp, set move to previous move and steps--
	move := oppositeDirection(previous_move)
	updatePosition(move)
	comp.Compute(move)
	steps--

}

func fullMaze() bool {
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if maps[i][j] == 0 {
				return false
			}
		}
	}

	return true
}

func updatePosition(move int) {
	if move == 1 { //north
		ypos--
	} else if move == 2 { //south
		ypos++
	} else if move == 3 { //west
		xpos--
	} else { //east
		xpos++
	}
}

func oppositeDirection(i int) int {
	if i == 1 {
		return 2
	} else if i == 2 {
		return 1
	} else if i == 3 {
		return 4
	} else {
		return 3
	}
}

func showMap() {
	for i := 0; i < len(maps); i++ {
		for j := 0; j < len(maps[i]); j++ {
			if maps[j][i] == 0 {
				fmt.Printf(" ")
			} else if maps[j][i] == 1 {
				fmt.Printf("#")
			} else if maps[j][i] == 2 {
				fmt.Printf("-")
			} else if maps[j][i] == 3 { //end point
				fmt.Printf("X")
			} else if maps[j][i] == 4 { //starting point
				fmt.Printf("0")
			}

		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
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
