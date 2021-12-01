package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	nums := readFile()
	ans := 0

	for x := 0; x < len(nums); x++ {
		ans += toLaunchAdv(nums[x])
	}

	fmt.Println(ans)

}

func toLaunch(mass int) int {
	mass = mass/3 - 2

	return mass
}

func toLaunchAdv(mass int) int {
	totalMass := 0
	for mass > 0 {
		mass = mass/3 - 2

		if mass < 0 {
			mass = 0
		}

		totalMass += mass
	}
	return totalMass
}

func readFile() []int {
	file, err := os.Open("data.txt")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var perline int
	var nums []int

	for {
		_, err := fmt.Fscanf(file, "%d\n", &perline)

		if err != nil {
			if err == io.EOF {
				break
			}

			fmt.Println(err)
			os.Exit(1)
		}

		nums = append(nums, perline)
	}

	return nums
}
