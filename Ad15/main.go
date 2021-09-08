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
var full_maze bool = false
var steps int = 0

const size int = 50

var maps [size][size]int
var xstart = size / 2
var ystart = size / 2
var xpos = xstart
var ypos = ystart
var xend int
var yend int

var rowNum = [4]int{-1, 0, 0, 1}
var colNum = [4]int{0, -1, 1, 0}

//1 is north, 2 is south, 3 is west, 4 is east

func main() {
	var previous_command int = 0
	maps[xpos][ypos] = 3
	remoteControlProgram(previous_command)
	maps[xstart][ystart] = 2
	maps[xend][yend] = 3
	showMap()
	maps[xstart][ystart] = 1
	maps[xend][yend] = 1
	dist := findShortestPath()
	fmt.Println(dist)
}

func findShortestPath() int {
	var oxygen = 0
	var visited [size][size]bool
	var q []queueNode
	var currNode queueNode
	var pt Point
	startpoint := Point{xend, yend}
	//des := Point{xend, yend}

	visited[xstart][ystart] = true
	s := queueNode{startpoint, 0}
	q = append(q, s)

	for len(q) != 0 {
		fmt.Println(q)
		currNode, q = q[0], q[1:]
		pt = currNode.pt
		// if pt.x == des.x && pt.y == des.y {
		// 	return currNode.dist
		// }

		for i := 0; i < 4; i++ {
			var row = pt.x + rowNum[i]
			var col = pt.y + colNum[i]

			if isValid(row, col) && maps[row][col] == 1 && !visited[row][col] {
				visited[row][col] = true
				q = append(q, queueNode{Point{row, col}, currNode.dist + 1})

			}
		}
		oxygen++
	}

	//Destination cannot be reached
	return oxygen
}

func isValid(row int, col int) bool {
	return (row >= 0) && (row < size) && (col >= 0) && (col < size)
}

func remoteControlProgram(previous_move int) {

	for move := 1; move <= 4; move++ {

		if steps == 1100 {
			return
		}

		if oppositeDirection(move) == previous_move { //if move is opposite direction from the previous move, move cannot be performed to avoid looping
			continue
		}

		updatePosition(move)

		output := comp.Compute(move)
		maps[xpos][ypos] = output

		if output == 0 { //crash to wall, check further position
			updatePosition(oppositeDirection(move))
			continue
		} else if output == 1 { //possible path! go find next path
			steps++
			remoteControlProgram(move)
		} else if output == 2 { //oxygen station found
			xend = xpos
			yend = ypos
			return
		}
	}

	//when it comes here, all possibilties are checked and no move is allowed except for the opposite direction. we perform move back, calculate  in the intcomp, set move to previous move and steps--
	move := oppositeDirection(previous_move)
	updatePosition(move)
	comp.Compute(move)
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
				fmt.Printf("#")
			} else if maps[j][i] == 1 {
				fmt.Printf("-")
			} else if maps[j][i] == 2 { //end point
				fmt.Printf("X")
			} else if maps[j][i] == 3 { //starting point
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

//-------------------------------------------------------------------------------------------------------------------------------------------

type Point struct {
	x int
	y int
}

type queueNode struct {
	pt   Point
	dist int
}
