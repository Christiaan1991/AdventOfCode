package main

import (
	"fmt"
	"strconv"
)

func main() {
	solvePassword(138241, 674034)
}

func solvePassword(num1 int, num2 int) {
	counter := 0
	for i := num1; i <= num2; i++ {
		if advancedAdjacent(i) && notDecrease(i) {
			fmt.Println(i)
			counter++
		}
	}

	fmt.Println(counter)
}

func adjacent(num int) bool {
	slice := strconv.Itoa(num)
	for i := 0; i < len(slice)-1; i++ {
		if slice[i] == slice[i+1] {
			return true
		}
	}

	return false
}

func advancedAdjacent(num int) bool {
	slice := strconv.Itoa(num)
	for i := 0; i < len(slice)+1; i++ {
		j := 0
		for i < len(slice)-1 && slice[i] == slice[i+1] {
			i++
			j++
		}
		if j == 1 {
			return true
		}
	}

	return false
}

func notDecrease(num int) bool {
	slice := strconv.Itoa(num)
	for i := 0; i < len(slice)-1; i++ {
		if slice[i] > slice[i+1] {
			//fmt.Println(slice)
			return false
		}
	}

	return true
}
