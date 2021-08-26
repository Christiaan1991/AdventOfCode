package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	nums := readFile("data.txt")

	//layer := DigitalSendingNetwork(nums)
	//ones, twos := getDigits(nums, layer)
	newImage(nums)

}

func PrintImage(newimage [][]int) {
	width := len(newimage[0])
	height := len(newimage)

	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	// Set color for each pixel.
	for x := 0; x < height; x++ {
		for y := 0; y < width; y++ {
			fmt.Println(x, y)
			if newimage[x][y] == 0 {
				img.Set(y, x, color.Black)
			}
			if newimage[x][y] == 1 {
				img.Set(y, x, color.White)
			}
		}
	}

	// Encode as PNG.
	f, _ := os.Create("image.png")
	png.Encode(f, img)
}

func newImage(nums [][][]int) {
	var row []int
	var image [][]int

	for j := 0; j < len(nums[0]); j++ {
		for k := 0; k < len(nums[0][j]); k++ {
			for layer := 0; layer < len(nums); layer++ {
				if nums[layer][j][k] != 2 {
					row = append(row, nums[layer][j][k])
					break
				}
			}

		}
		image = append(image, row)
		row = nil
	}

	fmt.Println(image)

	PrintImage(image)

}

func getDigits(nums [][][]int, layer int) (int, int) {

	ones := 0
	twos := 0
	for j := 0; j < len(nums[layer]); j++ {
		for i := 0; i < len(nums[j][layer]); i++ {
			if nums[layer][j][i] == 1 {
				ones++
			}
			if nums[layer][j][i] == 2 {
				twos++
			}
		}
	}
	return ones, twos
}

func DigitalSendingNetwork(nums [][][]int) int {

	var layer int
	current_counter := len(nums[0][0]) * len(nums[0])

	for k := 0; k < len(nums); k++ {
		counter := 0
		for j := 0; j < len(nums[k]); j++ {
			for i := 0; i < len(nums[k][j]); i++ {
				if nums[k][j][i] == 0 {
					counter++
				}
			}
		}

		if counter < current_counter { //if number of zeros is smaller than in current counter
			current_counter = counter
			fmt.Println(counter)
			layer = k
		}

	}

	return layer
}

func readFile(fname string) [][][]int {
	var nums [][][]int
	var row2 [][]int
	var row1 []int
	i := 0
	j := 0
	k := 0

	b, err := ioutil.ReadFile(fname)
	if err != nil {
		return nums
	}

	lines := strings.Split(string(b), "")

	for _, l := range lines {
		n, _ := strconv.Atoi(l)
		row1 = append(row1, n)
		i++
		if i == 25 {
			row2 = append(row2, row1)
			i = 0
			row1 = nil
			j++
			if j == 6 {
				nums = append(nums, row2)
				j = 0
				row2 = nil
				k++
			}
		}
	}
	return nums
}
