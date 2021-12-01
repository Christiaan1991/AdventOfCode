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
	data := readFile()
	//data := []int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9}

	counter := 0

	for {
		data, counter = intCode(data, counter)
	}

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

func intCode(output []int, counter int) ([]int, int) {

	var first_param int
	var second_param int
	instruction := generateInstruction(strconv.Itoa(output[counter]))
	opcode := string(instruction[3:5])
	fmt.Println(opcode)

	if opcode == "99" { //exit program, part of 99
		fmt.Println("Exit program succesful!")
		os.Exit(0)
	}

	mode1 := string(instruction[2])
	mode2 := string(instruction[1])
	mode3 := string(instruction[0])

	if opcode == "01" { //addition
		first_param = getParam(output, counter+1, mode1)
		second_param = getParam(output, counter+2, mode2)
		addition(output, first_param, second_param, counter+3, mode3)

		counter += 4
		return output, counter
	}

	if opcode == "02" { //multiplication
		first_param = getParam(output, counter+1, mode1)
		second_param = getParam(output, counter+2, mode2)
		multiplication(output, first_param, second_param, counter+3, mode3)

		counter += 4
		return output, counter
	}

	if opcode == "03" { //input

		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Enter single integer:")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSuffix(input, "\r\n")

		if mode1 == "0" {
			output[output[counter+1]], _ = strconv.Atoi(input)
		}

		if mode1 == "1" {
			output[counter+1], _ = strconv.Atoi(input)
		}

		counter += 2
		return output, counter

	}

	if opcode == "04" { //output
		if mode1 == "0" {
			fmt.Println(output[output[counter+1]])
		}
		if mode1 == "1" {
			fmt.Println(output[counter+1])
		}

		counter += 2
		return output, counter
	}

	if opcode == "05" {
		first_param = getParam(output, counter+1, mode1)

		if first_param != 0 {
			counter = getParam(output, counter+2, mode2)
			return output, counter
		}
		counter += 3
		return output, counter
	}

	if opcode == "06" {

		first_param = getParam(output, counter+1, mode1)
		fmt.Println("first param:", first_param)

		if first_param == 0 {
			counter = getParam(output, counter+2, mode2)
			return output, counter
		}
		counter += 3
		return output, counter

	}

	if opcode == "07" {
		first_param = getParam(output, counter+1, mode1)
		second_param = getParam(output, counter+2, mode2)
		less_than(output, first_param, second_param, counter+3, mode3)

		counter += 4
		return output, counter

	}

	if opcode == "08" {
		first_param = getParam(output, counter+1, mode1)
		second_param = getParam(output, counter+2, mode2)
		equals(output, first_param, second_param, counter+3, mode3)

		counter += 4
		return output, counter
	}

	os.Exit(1) //program terminates if no opcode is found!

	return output, counter
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

	return nums
}
