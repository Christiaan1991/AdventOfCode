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

const (
	gamma = "gamma"
	epsilon = "epsilon"
)

func main(){
	input := readFile("input")
	part1(input)
	part2(input)
}

func part2(input []string) {
	value := getRating(input, "oxygen")
	value2 := getRating(input,"CO2")
	fmt.Printf("%d x %d = %d", value, value2, value*value2)
}

func getRating(input []string, element string) int {
	i := 0
	for i < 12 {
		input = findNewBits(input, i, element)
		if len(input) == 1 {
			fmt.Println(input)
			break
		}
		i++
	}
	return toNumber(strToIntArr(input[0]))
}

func strToIntArr(s string) []int {
	strs := strings.Split(s, "")
	res := make([]int, len(strs))
	for i := range res {
		res[i], _ = strconv.Atoi(strs[i])
	}
	return res
}

func findNewBits(input []string, pos int, element string) []string {
	var oneBits []string
	var zeroBits []string
	for _, bits := range input {
		bit := getBit(bits, pos)
		if bit == 0 {
			zeroBits = append(zeroBits, bits)
		} else {
			oneBits = append(oneBits, bits)
		}
	}
	if element == "oxygen" {
		if len(zeroBits) > len(oneBits) {
			return zeroBits
		} else {
			return oneBits
		}
	} else {
		if len(zeroBits) <= len(oneBits) {
			return zeroBits
		} else {
			return oneBits
		}
	}

}

func part1(input []string) {
	gamma := getRate(input, gamma)
	epsilon := getRate(input, epsilon)
	fmt.Println(gamma*epsilon)
}

func powInt(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

func toNumber(s []int) int {
	num := 0
	for i := 0; i < len(s) ; i++ {
		num += s[i] * powInt(2, len(s)-1-i)
	}
	return num
}

func getRate(input []string, rate string) int {
	var finalBit []int
	for i := 0; i < 11; i++ {
		var zeroBit = 0
		var oneBit = 0
		for _, bits := range input {
			Bit := getBit(bits, i)
			if Bit == 0 {
				zeroBit++
			} else {
				oneBit++
			}
		}
		if rate == gamma {
			if zeroBit > oneBit {
				finalBit = append(finalBit, 0)
			} else {
				finalBit = append(finalBit, 1)
			}
		} else {
			if zeroBit < oneBit {
				finalBit = append(finalBit, 0)
			} else {
				finalBit = append(finalBit, 1)
			}
		}

		fmt.Println(finalBit)
	}
	return toNumber(finalBit)
}

func getBit(input string, pos int) int {
	num, err := strconv.Atoi(input[pos:pos+1])
	if err != nil {
		panic(err)
	}
	return num
}

func readFile(filename string) []string {
	file, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	s := strings.Split(string(data), "\n")
	return s
}