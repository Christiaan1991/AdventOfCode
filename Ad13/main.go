package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

var program = readFile()
var screen [50][50]int
var icons = map[int]string{0: " ", 1: "\u25A1", 2: ".", 3: "_", 4: "o"}

var relativeBase int = 0
var inputs = []int{}
var intCodeFinished bool = false

func main() {

	//fmt.Println("\u25A1")

	counter := 0

	for !intCodeFinished {
		inputs, counter = intCodeProgram(inputs, counter)

		if len(inputs) == 3 {
			screen[inputs[0]][inputs[1]] = inputs[2]

			segmentDisplay()
			inputs = []int{}
		}
	}
}

func segmentDisplay() {

	for y := 0; y < len(screen); y++ {
		var output string = ""
		for x := 0; x < len(screen[0]); x++ {
			output += icons[screen[x][y]]
		}
		fmt.Println(output)
	}

}

func generateInstruction(instr string) string {
	for len(instr) < 5 {
		instr = "0" + instr
	}
	return instr
}

func setParam(result int, counter int, mode string) {
	if mode == "0" {
		program[program[counter]] = result
	}

	if mode == "1" {
		program[counter] = result
	}

	program[relativeBase+program[counter]] = result
}

func getParam(counter int, mode string) int {

	if mode == "0" {
		return program[program[counter]]
	}

	if mode == "1" {
		return program[counter]
	}

	if mode == "2" {
		return program[program[counter]+relativeBase]
	}

	fmt.Println("Should not come here")
	return program[counter]

}

func intCodeProgram(inputs []int, counter int) ([]int, int) {
	var output_signal int

	for {
		var first_param int
		var second_param int
		var result int

		instruction := generateInstruction(strconv.Itoa(program[counter]))
		opcode := string(instruction[3:5])

		mode1 := string(instruction[2])
		mode2 := string(instruction[1])
		mode3 := string(instruction[0])

		switch opcode {
		case "99": //exit ampControl, part of 99
			intCodeFinished = true
			return inputs, counter

		case "01":
			first_param = getParam(counter+1, mode1)
			second_param = getParam(counter+2, mode2)
			result = first_param + second_param
			setParam(result, counter+3, mode3)

			counter += 4
			break

		case "02":
			first_param = getParam(counter+1, mode1)
			second_param = getParam(counter+2, mode2)
			result = first_param * second_param
			setParam(result, counter+3, mode3)

			counter += 4
			break

		case "03":
			setParam(inputs[0], counter+1, mode1)
			inputs = RemoveIndex(inputs, 0)
			counter += 2
			break

		case "04":
			output_signal = getParam(counter+1, mode1)
			inputs = append(inputs, output_signal)
			counter += 2
			return inputs, counter

		case "05":
			first_param = getParam(counter+1, mode1)
			second_param = getParam(counter+2, mode2)
			if first_param != 0 {
				counter = second_param
			} else {
				counter += 3
			}
			break

		case "06":
			first_param = getParam(counter+1, mode1)
			second_param = getParam(counter+2, mode2)
			if first_param == 0 {
				counter = second_param
			} else {
				counter += 3
			}
			break
		case "07":
			first_param = getParam(counter+1, mode1)
			second_param = getParam(counter+2, mode2)
			if first_param < second_param {
				result = 1
			} else {
				result = 0
			}
			setParam(result, counter+3, mode3)
			counter += 4
			break

		case "08":
			first_param = getParam(counter+1, mode1)
			second_param = getParam(counter+2, mode2)
			if first_param == second_param {
				result = 1
			} else {
				result = 0
			}
			setParam(result, counter+3, mode3)
			counter += 4
			break
		case "09":
			first_param = getParam(counter+1, mode1)
			relativeBase += first_param
			counter += 2
			break
		}

	}
}

func RemoveIndex(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}

func RemoveIndex2(s [][2]int, index int) [][2]int {
	return append(s[:index], s[index+1:]...)
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
