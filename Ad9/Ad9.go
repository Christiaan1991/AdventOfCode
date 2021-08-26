package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

var program = readFile()
var relativeBase int = 0
var inputs = []int{2}

func main() {

	output, _ := amplifier(0, inputs)
	fmt.Println(output)
}

func getCodes() [][]int {
	array := []int{5, 6, 7, 8, 9}
	codes := permutations(array)

	return codes
}

func permutations(arr []int) [][]int {
	var helper func([]int, int)
	res := [][]int{}

	helper = func(arr []int, n int) {
		if n == 1 {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
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

func amplifier(counter int, inputs []int) ([]int, int) {

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
			break

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

	fmt.Println("should not come here")
	os.Exit(1)
	return inputs, counter
}

func RemoveIndex(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}

func Reader() int {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter single integer:")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSuffix(input, "\r\n")
	inputint, _ := strconv.Atoi(input)
	return inputint

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

	for len(nums) < 10000 {
		nums = append(nums, 0)
	}

	return nums
}
