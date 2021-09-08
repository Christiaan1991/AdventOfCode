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
var intersectionPoints []Point
var directions = []string{"UP", "DOWN", "LEFT", "RIGHT"}

//ASCII code:
// 35: #
// 46: .
// 60: <
// 62: >
// 94: ^
//
// 10: \n
// 44: ,
// 65: A
// 66: B
// 67: C
// 82: R
// 76: L

func main() {
	//part1()
	part2()
}

func part2() {
	createImage()
	findIntersections()
	showMap()
	// var program = readFile()
	// program[0] = 2
	robot := startingPosition()
	fmt.Println(robot)

}

func findShortestPath(robot Robot) {
	var visited [size][size]bool
	var q []Point
	var rowNum = [4]int{-1, 0, 0, 1}
	var colNum = [4]int{0, -1, 1, 0}
	var currPoint Point

	visited[robot.pt.x][robot.pt.y] = true
	s := robot.pt
	q = append(q, s)

	for len(q) != 0 {
		fmt.Println(q)
		currPoint, q = q[0], q[1:]

		for i := 0; i < 4; i++ {
			var row = currPoint.x + rowNum[i]
			var col = currPoint.y + colNum[i]

			if isValid(row, col) && image[row][col] == 1 && (!visited[row][col] || isIntersection(row, col)) {
				visited[row][col] = true
				q = append(q, Point{row, col})

			}
		}

	}

}

func isValid(row int, col int) bool {
	return (row >= 0) && (row < size) && (col >= 0) && (col < size)
}

func startingPosition() Robot {
	//find starting position of the cleaning robot
	var robot Robot
	for i := 0; i < len(image); i++ {
		for j := 0; j < len(image[i]); j++ {
			if image[i][j] == 94 { //we know it starts upwards
				robot = Robot{"UP", Point{i, j}}
				return robot
			}
		}
	}
	//cannot find starting point for the robot
	os.Exit(1)
	return robot
}

func isIntersection(row int, col int) bool {
	for _, intersectionPoint := range intersectionPoints {
		if intersectionPoint.x == row && intersectionPoint.y == col {
			return true
		}
	}
	return false
}

func findIntersections() {
	for i := 1; i < size; i++ {
		for j := 1; j < size; j++ {
			if image[i][j] == 35 && image[i+1][j] == 35 && image[i-1][j] == 35 && image[i][j+1] == 35 && image[i][j-1] == 35 {
				intersectionPoints = append(intersectionPoints, Point{i, j})
			}
		}
	}
}

func createImage() {
	var i int = 0
	var j int = 0
	var program = readFile()
	var image_comp = intcomp.Computer{program, false, 0, 0}

	for {
		image[i][j] = image_comp.Compute(0)

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
			switch image[i][j] {
			case 35:
				fmt.Printf("#")
			case 46:
				fmt.Printf(".")
			case 60:
				fmt.Printf("<")
			case 62:
				fmt.Printf(">")
			case 94:
				fmt.Printf("^")
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

/*----------------------------------------------------------------------------------*/
type Point struct {
	x int
	y int
}

type Robot struct {
	direction string
	pt        Point
}

func (robot *Robot) toLeft() {
	if robot.direction == "UP" {
		robot.direction = "LEFT"
		return
	} else if robot.direction == "LEFT" {
		robot.direction = "DOWN"
		return
	} else if robot.direction == "DOWN" {
		robot.direction = "RIGHT"
		return
	} else if robot.direction == "RIGHT" {
		robot.direction = "UP"
		return
	}
}

func (robot *Robot) toRight() {
	if robot.direction == "UP" {
		robot.direction = "RIGHT"
		return
	} else if robot.direction == "RIGHT" {
		robot.direction = "DOWN"
		return
	} else if robot.direction == "DOWN" {
		robot.direction = "LEFT"
		return
	} else if robot.direction == "LEFT" {
		robot.direction = "UP"
		return
	}
}

func (robot *Robot) move(steps int) {
	if robot.direction == "UP" {
		robot.y = robot.y - steps
		return
	} else if robot.direction == "RIGHT" {
		robot.x = robot.x + steps
		return
	} else if robot.direction == "DOWN" {
		robot.y = robot.y + steps
		return
	} else if robot.direction == "LEFT" {
		robot.x = robot.x - steps
		return
	}
}
