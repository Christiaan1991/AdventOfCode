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
var new_image [size][size]int
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
	var program = readFile()
	var intcomp = intcomp.Computer{program, false, 0, 0}

	createImage(intcomp)
	findIntersections()

	robot := startingPosition()
	shortestPath := findPath(robot)

	signal := createMethods(shortestPath)

	program[0] = 2

	var k int = 0

	for {
		value := intcomp.Compute(int(signal[k]))
		showMap(value)
		if value == 0 {
			break
		}
		k++
		// if k == len(signal) {
		// 	k = 0
		// }

	}
}

func createMethods(path string) []byte {

	var A string = "L,6,R,8,R,12,L,6,L,8"
	var B string = "L,10,L,8,R,12"
	var C string = "L,8,L,10,L,6,L,6"
	var next_path string = ""
	var i int = 0

	for i < len(path) {
		if (i+len(A) <= len(path)) && path[i:i+len(A)] == A {
			next_path += "A"
			i += len(A)
		} else if (i+len(B) <= len(path)) && path[i:i+len(B)] == B {
			next_path += "B"
			i += len(B)
		} else if (i+len(C) <= len(path)) && path[i:i+len(C)] == C {
			next_path += "C"
			i += len(C)
		} else {
			next_path += string(path[i])
			i++
		}
	}

	output := next_path + "\n" + A + "\n" + B + "\n" + C + "\n" + "y" + "\n"
	fmt.Println(output)
	signal := []byte(output)

	return signal
}

func findPath(robot Robot) string {
	//we find a path for the robot by going straight until there is no other option but to turn. The path is saved in string format, and print back
	var path string
	var endpoint_found bool = true
	var steps int = 0

	for endpoint_found {
		if robot.canMoveForward() {
			//robot can move forward
			steps++
		} else if canRot, rotation := robot.canRotate(); canRot {
			//robot can rotate left or right. Rotate to right direction, and reset steps to 1, since we perform step in canRotate method
			if steps != 0 {
				path = path + strconv.Itoa(steps) + ","
			}
			path = path + rotation + ","

			steps = 1

		} else {
			//cannot move forward or rotate, meaning that we reached endpoint!
			path = path + strconv.Itoa(steps)
			endpoint_found = false
		}
	}

	return path

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

func createImage(intcomp intcomp.Computer) {
	var i int = 0
	var j int = 0

	for {
		value := intcomp.Compute(0)

		image[i][j] = value
		r := string(rune(value))
		fmt.Printf("%v", r)

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

func showMap(value int) {
	r := string(rune(value))
	fmt.Printf("%v", r)
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

	for len(nums) < 100000 {
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

func (robot *Robot) move_forward() {
	if robot.direction == "UP" {
		robot.pt.x = robot.pt.x - 1
		return
	} else if robot.direction == "RIGHT" {
		robot.pt.y = robot.pt.y + 1
		return
	} else if robot.direction == "DOWN" {
		robot.pt.x = robot.pt.x + 1
		return
	} else if robot.direction == "LEFT" {
		robot.pt.y = robot.pt.y - 1
		return
	}
}

func (robot *Robot) move_backward() {
	if robot.direction == "UP" {
		robot.pt.x = robot.pt.x + 1
		return
	} else if robot.direction == "RIGHT" {
		robot.pt.y = robot.pt.y - 1
		return
	} else if robot.direction == "DOWN" {
		robot.pt.x = robot.pt.x - 1
		return
	} else if robot.direction == "LEFT" {
		robot.pt.y = robot.pt.y + 1
		return
	}
}

func (robot *Robot) canMoveForward() bool {
	robot.move_forward()

	if isValid(robot.pt.x, robot.pt.y) && image[robot.pt.x][robot.pt.y] == 35 {
		//move is valid (not out of bounds), and empty space
		return true
	} else {
		robot.move_backward()
		return false
	}
}

func (robot *Robot) canRotate() (bool, string) {
	robot.toLeft()

	if robot.canMoveForward() {
		return true, "L"
	}

	robot.toRight()
	robot.toRight()

	if robot.canMoveForward() {
		return true, "R"
	}

	return false, "N"
}
