package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

var base_pattern = [4]int{0, 1, 0, -1}

const PHASES int = 100
const REPEAT int = 10000

//since the input is now on repeat, we just have to calculate the last digits of the original output

func main() {
	var input_signal = readFile()
	offset := offset(input_signal[0:7])

	output_signal := part2(input_signal, offset)
	fmt.Println(output_signal[offset : offset+8])

}

func part2(input_signal []int, offset int) []int {
	var count int = 0
	for count < PHASES {
		input_signal = phase2(input_signal, offset)
		count++
	}

	return input_signal
}

func part1(input_signal []int, offset int) {
	var count int = 0
	for count < PHASES {

		input_signal = phase(input_signal)
		count++
		fmt.Println(count)
	}

	fmt.Println(input_signal[offset : offset+8])
}

func offset(input_signal []int) int {
	var num int = 0
	for j := 0; j < len(input_signal); j++ {
		num += input_signal[(len(input_signal)-1)-j] * int(math.Pow(10, float64(j)))
	}
	return num
}

func phase2(input_signal []int, offset int) []int {
	var val int = 0

	for i := len(input_signal) - 1; i >= offset; i-- {
		val += input_signal[i]
		input_signal[i] = val % 10
	}
	return input_signal
}

func insert(a []int, index int, value int) []int {
	if len(a) == index { // nil or empty slice or after last element
		return append(a, value)
	}
	a = append(a[:index+1], a[index:]...) // index < len(a)
	a[index] = value
	return a
}

func phase(input_signal []int) []int {
	var output_signal []int

	for i := 0; i < len(input_signal); i++ {
		pattern := createPattern(input_signal, i)
		var num int = 0

		for j := 0; j < len(input_signal); j++ {
			num += pattern[j] * input_signal[j]
		}
		output_signal = append(output_signal, Abs(num%10))
	}

	return output_signal
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func createPattern2(input_signal []int, i int, offset int) []int {
	var pattern []int
	var k int = 0

	for j := 0; j < offset+1; j++ {
		for k = 0; k <= i; k++ {
			pattern = append(pattern, base_pattern[j%len(base_pattern)])
			if len(pattern) > offset+1 {
				return pattern[1 : offset+1]
			}
		}
	}
	return pattern[1 : offset+1]
}

func createPattern(input_signal []int, i int) []int {
	var pattern []int
	var k int = 0

	for j := 0; j < len(input_signal)+1; j++ {
		for k = 0; k <= i; k++ {
			pattern = append(pattern, base_pattern[j%len(base_pattern)])
			if len(pattern) > len(input_signal)+1 {
				return pattern[1 : len(input_signal)+1]
			}
		}
	}
	return pattern[1 : len(input_signal)+1]
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

	s := strings.Split(string(data), "")

	var count int = 0
	for count < REPEAT {
		for _, i := range s {
			j, err := strconv.Atoi(i)

			if err != nil {
				panic(err)
			}

			nums = append(nums, j)
		}
		count++
	}

	return nums
}
