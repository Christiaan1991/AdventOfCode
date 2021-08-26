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

	wire1data, wire2data := readFile()
	//wire1data := []string{"R75", "D30", "R83", "U83", "L12", "D49", "R71", "U7", "L72"}
	//wire2data := []string{"U62", "R66", "U55", "R34", "D71", "R55", "D58", "R83"}

	//wire1data := []string{"R98", "U47", "R26", "D63", "R33", "U87", "L62", "D20", "R33", "U53", "R51"}
	//wire2data := []string{"U98", "R91", "D20", "R16", "D67", "R40", "U7", "R15", "U6", "R7"}

	//wire1data := []string{"R8", "U5", "L5", "D3"}
	//wire2data := []string{"U7", "R6", "D4", "L4"}

	wire1 := getWirePositions(wire1data)
	wire2 := getWirePositions(wire2data)

	compareWires(wire1, wire2)

	compareWiresSteps(wire1, wire2)

}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func compareWiresSteps(wire1 [][2]int, wire2 [][2]int) {
	var steps int
	for i := 0; i < len(wire1); i++ {
		for j := 1; j < len(wire2); j++ {
			if wire1[i][0] == wire2[j][0] && wire1[i][1] == wire2[j][1] {
				if steps == 0 || i+j < steps {
					steps = i + j
					fmt.Println("New smallest steps:", steps)
				}

			}

		}
	}
}

func compareWires(wire1 [][2]int, wire2 [][2]int) {
	var distance int
	for i := 1; i < len(wire1); i++ {
		for j := 1; j < len(wire2); j++ {
			if wire1[i][0] == wire2[j][0] && wire1[i][1] == wire2[j][1] {
				if distance == 0 || Abs(wire1[i][0])+Abs(wire1[i][1]) < distance {
					distance = Abs(wire1[i][0]) + Abs(wire1[i][1])
					fmt.Println("New smallest intersection found with distance:", distance)
				}

			}
		}
	}
}

func getWirePositions(data []string) [][2]int {
	var wire [][2]int
	xpoint := 0
	ypoint := 0

	wire = append(wire, [2]int{xpoint, ypoint})

	for _, i := range data {
		direction := string(i[0])
		num, _ := strconv.Atoi(string(i[1:]))

		if direction == "R" {
			wire, xpoint, ypoint = goRight(wire, num, xpoint, ypoint)
		}
		if direction == "D" {
			wire, xpoint, ypoint = goDown(wire, num, xpoint, ypoint)
		}
		if direction == "L" {
			wire, xpoint, ypoint = goLeft(wire, num, xpoint, ypoint)
		}
		if direction == "U" {
			wire, xpoint, ypoint = goUp(wire, num, xpoint, ypoint)
		}
	}
	return wire
}

func goRight(wire [][2]int, num int, xpoint int, ypoint int) ([][2]int, int, int) {

	for i := 1; i <= num; i++ {
		xpoint++
		wire = append(wire, [2]int{xpoint, ypoint})
	}

	return wire, xpoint, ypoint

}

func goDown(wire [][2]int, num int, xpoint int, ypoint int) ([][2]int, int, int) {

	for i := 1; i <= num; i++ {
		ypoint--
		wire = append(wire, [2]int{xpoint, ypoint})
	}

	return wire, xpoint, ypoint

}

func goLeft(wire [][2]int, num int, xpoint int, ypoint int) ([][2]int, int, int) {

	for i := 1; i <= num; i++ {
		xpoint--
		wire = append(wire, [2]int{xpoint, ypoint})
	}

	return wire, xpoint, ypoint

}

func goUp(wire [][2]int, num int, xpoint int, ypoint int) ([][2]int, int, int) {

	for i := 1; i <= num; i++ {
		ypoint++
		wire = append(wire, [2]int{xpoint, ypoint})
	}

	return wire, xpoint, ypoint

}

func readFile() ([]string, []string) {

	file, err := os.Open("data.txt")

	if err != nil {
		log.Fatal(err)
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}

	s := strings.Fields(string(data))

	wire1 := strings.Split(s[0], ",")
	wire2 := strings.Split(s[1], ",")

	return wire1, wire2

}
