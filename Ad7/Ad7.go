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

func main() {
	codes := getCodes()
	highest_output := 0

	//code := []int{9, 7, 8, 5, 6}

	for i := 0; i < len(codes); i++ {
		output := ampControlSoft(codes[i])

		if output > highest_output {
			highest_output = output
		}
	}

	fmt.Println(highest_output)

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

func getParam(output []int, num int, mode string) int {
	if mode == "0" {
		return output[output[num]]
	}
	return output[num]
}

func addition(output []int, param1 int, param2 int, num int, mode string) []int {
	if mode == "0" {
		output[output[num]] = param1 + param2
		return output
	}
	output[num] = param1 + param2
	return output
}

func multiplication(output []int, param1 int, param2 int, num int, mode string) []int {
	if mode == "0" {
		output[output[num]] = param1 * param2
		return output
	}
	output[num] = param1 * param2
	return output
}

func ampControlSoft(phases []int) int {
	var (
		data     = [][]int{readFile(), readFile(), readFile(), readFile(), readFile()}
		counters = []int{0, 0, 0, 0, 0}
		output   = 0
	)

	output, data[0], counters[0] = amplifier(data[0], counters[0], []int{phases[0], output})
	output, data[1], counters[1] = amplifier(data[1], counters[1], []int{phases[1], output})
	output, data[2], counters[2] = amplifier(data[2], counters[2], []int{phases[2], output})
	output, data[3], counters[3] = amplifier(data[3], counters[3], []int{phases[3], output})
	output, data[4], counters[4] = amplifier(data[4], counters[4], []int{phases[4], output})
	tothruster := output

	for {
		for i := 0; i < len(phases); i++ {
			output, data[i], counters[i] = amplifier(data[i], counters[i], []int{output})
			if output == 0 {
				return tothruster
			}
		}

		tothruster = output
	}
}

func amplifier(code []int, counter int, inputs []int) (int, []int, int) {

	var output_signal int

	for {
		var first_param int
		var second_param int
		instruction := generateInstruction(strconv.Itoa(code[counter]))
		opcode := string(instruction[3:5])

		if opcode == "99" { //exit ampControl, part of 99
			return 0, []int{0}, 0
		}

		mode1 := string(instruction[2])
		mode2 := string(instruction[1])
		mode3 := string(instruction[0])

		if opcode == "01" { //addition
			first_param = getParam(code, counter+1, mode1)
			second_param = getParam(code, counter+2, mode2)
			addition(code, first_param, second_param, counter+3, mode3)

			counter += 4
			continue
		}

		if opcode == "02" { //multiplication
			first_param = getParam(code, counter+1, mode1)
			second_param = getParam(code, counter+2, mode2)
			multiplication(code, first_param, second_param, counter+3, mode3)

			counter += 4
			continue
		}

		if opcode == "03" { //input

			if mode1 == "0" {
				code[code[counter+1]] = inputs[0]
			}

			if mode1 == "1" {
				code[counter+1] = inputs[0]
			}
			inputs = RemoveIndex(inputs, 0)
			counter += 2
			continue
		}

		if opcode == "04" { //output
			if mode1 == "0" {
				output_signal = code[code[counter+1]]
				inputs = append(inputs, output_signal)
			}
			if mode1 == "1" {
				output_signal = code[counter+1]
				inputs = append(inputs, output_signal)
			}

			counter += 2
			return output_signal, code, counter
		}

		if opcode == "05" {
			first_param = getParam(code, counter+1, mode1)
			if first_param != 0 {
				counter = getParam(code, counter+2, mode2)
				continue
			}
			counter += 3
			continue
		}
		if opcode == "06" {

			first_param = getParam(code, counter+1, mode1)
			if first_param == 0 {
				counter = getParam(code, counter+2, mode2)
				continue
			}
			counter += 3
			continue

		}

		if opcode == "07" {
			first_param = getParam(code, counter+1, mode1)
			second_param = getParam(code, counter+2, mode2)
			less_than(code, first_param, second_param, counter+3, mode3)

			counter += 4
			continue

		}

		if opcode == "08" {
			first_param = getParam(code, counter+1, mode1)
			second_param = getParam(code, counter+2, mode2)
			equals(code, first_param, second_param, counter+3, mode3)

			counter += 4
			continue
		}
	}

	fmt.Println("should not come here")
	os.Exit(1)
	return output_signal, code, counter
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

func less_than(output []int, param1 int, param2 int, num int, mode string) []int {
	if mode == "0" {
		if param1 < param2 {
			output[output[num]] = 1
			return output
		}
		output[output[num]] = 0
		return output

	}

	if param1 < param2 {
		output[num] = 1
		return output
	}
	output[num] = 0
	return output
}

func equals(output []int, param1 int, param2 int, num int, mode string) []int {
	if mode == "0" {
		if param1 == param2 {
			output[output[num]] = 1
			return output
		}
		output[output[num]] = 0
		return output

	}

	if param1 == param2 {
		output[num] = 1
		return output
	}
	output[num] = 0
	return output
}

func readFile() []int {

	file, err := os.Open("testdata.txt")
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

	return nums
}
