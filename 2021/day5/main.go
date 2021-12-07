package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

const SIZE = 1000
var grid [SIZE][SIZE]int

func main(){
	input := readFile("input")
	coords := getCoords(input)

	part1(coords)
	part2(coords)
}

func getCoords(input []string) [][4]int {
	var coords [][4]int
	for i, _ := range input {
		var line [4]int
		s := strings.Split(input[i], " -> ")
		pos1 := strings.Split(s[0], ",")
		pos2 := strings.Split(s[1], ",")
		line[0], _ = strconv.Atoi(pos1[0])
		line[1], _ = strconv.Atoi(pos1[1])
		line[2], _ = strconv.Atoi(pos2[0])
		line[3], _ = strconv.Atoi(pos2[1])
		coords = append(coords, line)
	}
	return coords
}

func part2(coords [][4]int) {
	drawLines(coords)
	determineOverlap()
}

func part1(coords [][4]int) {
	drawLines(coords)
	determineOverlap()
}

func printGrid() {
	for i:=0; i < SIZE; i++ {
		for j:=0; j<SIZE; j++{
			fmt.Print(grid[j][i])
		}
		fmt.Print("\n")
	}
	fmt.Print("\n\n")
}

func determineOverlap() {
	var counter = 0
	for i:=0; i < SIZE; i++ {
		for j:=0; j < SIZE; j++ {
			if grid[i][j] > 1 {
				counter++
			}
		}
	}
	fmt.Printf("Number of overlaps in grid: %d\n", counter)
}

func drawLines(coords [][4]int) {
	for i:=0; i < len(coords); i++ {
		x1 := coords[i][0]
		y1 := coords[i][1]
		x2 := coords[i][2]
		y2 := coords[i][3]
		if x1 == x2 || y1 == y2 {
			//horizontal or vertical line
			if x1 == x2 {
				drawHorizontalLine(x1, y1, y2)
			} else {
				drawVerticalLine(x1, y1, x2)
			}
		} else if (x1 == y1 && x2 == y2) || (x1 == y2 && x2 == y1) || Abs(x2 - x1) == Abs(y2 - y1){
			//exactly diagonal
			if (x1 == y1 && x2 == y2) || (x2 - x1 == y2 - y1) {
				drawDiagonalLine1(x1, y1, x2, y2)
			} else {
				drawDiagonalLine2(x1, y1, x2, y2)
			}
		}
	}
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func drawDiagonalLine2(x1, y1, x2, y2 int) {
	if x1 <= x2 {
		for x1 <= x2 {
			grid[x1][y1]++
			x1++
			y1--
		}
		return
	} else {
		for x1 >= x2 {
			grid[x1][y1]++
			x1--
			y1++
		}
		return
	}
}

func drawDiagonalLine1(x1, y1, x2, y2 int) {
	if x1 >= x2 {
		for x1 >= x2 {
			grid[x1][y1]++
			x1--
			y1--
		}
		return
	} else {
		for x1 <= x2 {
			grid[x1][y1]++
			x1++
			y1++
		}
		return
	}
}

func drawHorizontalLine(x1, y1, y2 int) {
	if y1 <= y2 {
		for y1 <= y2 {
			grid[x1][y1]++
			y1++
		}
		return
	} else {
		for y1 >= y2 {
			grid[x1][y1]++
			y1--
		}
	}

}

func drawVerticalLine(x1, y1, x2 int) {
	if x1 <= x2 {
		for x1 <= x2 {
			grid[x1][y1]++
			x1++
		}
		return
	} else {
		for x1 >= x2 {
			grid[x1][y1]++
			x1--
		}
	}
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