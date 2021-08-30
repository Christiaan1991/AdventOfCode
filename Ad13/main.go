package main

import (
	"Ad13/intcomp"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

var program = readFile()
var icons = map[int]string{0: " ", 1: "\u25A1", 2: ".", 3: "_", 4: "o"}

func main() {
	var screen [50][50]int
	var score int
	var print bool

	program[0] = 2
	joystick := 0

	comp := intcomp.Computer{program, false, 0, 0}

	for !comp.Done {
		x := comp.Compute(joystick)
		y := comp.Compute(joystick)
		value := comp.Compute(joystick)

		if x == -1 && y == 0 {
			score = value
			print = true
		} else {
			screen[x][y] = value
		}

		hasBall, coords_ball := getBall(screen)
		hasPaddle, coords_paddle := getPaddle(screen)

		if hasBall && hasPaddle {
			if coords_paddle[0] > coords_ball[0] {
				joystick = -1
			} else if coords_paddle[0] < coords_ball[0] {
				joystick = 1
			} else {
				joystick = 0
			}
			print = true

		}

		if print {
			//segmentDisplay(screen, score)
		}

	}

	fmt.Println(score)

}

func getPaddle(screen [50][50]int) (bool, [2]int) {
	for x := 0; x < len(screen); x++ {
		for y := 0; y < len(screen[x]); y++ {
			if screen[x][y] == 3 {
				return true, [2]int{x, y}
			}
		}
	}
	return false, [2]int{-1, -1}
}

func getBall(screen [50][50]int) (bool, [2]int) {
	for x := 0; x < len(screen); x++ {
		for y := 0; y < len(screen[x]); y++ {
			if screen[x][y] == 4 {
				return true, [2]int{x, y}
			}
		}
	}
	return false, [2]int{-1, -1}
}

func segmentDisplay(screen [50][50]int, score int) {
	for y := 0; y < len(screen); y++ {
		var output string = ""
		for x := 0; x < len(screen[0]); x++ {
			output += icons[screen[x][y]]
		}
		fmt.Println(output)
	}
	fmt.Println("Score: ", score)
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
