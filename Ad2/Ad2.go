package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	for i := 0; i <= 99; i++ {
		for j := 0; j <= 99; j++ {
			nums := readFile()
			noun := i
			verb := j
			nums[1] = noun
			nums[2] = verb

			var output []int = nums
			fmt.Println(noun, verb)

			counter := 0
			for {
				if output[counter] == 99 { //if it returns 99, print the array and exit program!
					if output[0] == 19690720 {
						fmt.Println(output[1]*100 + output[2])
						os.Exit(0)
					}
					break
				}

				output = gravityProgram(output, counter)
				counter += 4
			}

		}
	}
	// nums[1] = 12
	// nums[2] = 2

	// counter := 0

	// for {
	// 	if nums[counter] == 99 { //if it returns 99, print the array and exit program!
	// 		fmt.Println(nums[0])
	// 		os.Exit(0)
	// 	}

	// 	nums = gravityProgram(nums, counter)
	// 	counter += 4
	// }
}

func gravityProgram(output []int, counter int) []int {
	instruction := output[counter]

	if instruction == 1 {
		output[output[counter+3]] = output[output[counter+1]] + output[output[counter+2]]
	}

	if instruction == 2 {
		output[output[counter+3]] = output[output[counter+1]] * output[output[counter+2]]
	}

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
