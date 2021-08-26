package main

import (
	"Ad11/robot"
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

var program = readFile()
var relativeBase int = 0
var inputs = []int{1}

func main() {

	var counter int = 0
	const size int = 200
	var num_panels int = 0
	var panels [size][size]int
	var painted_panels = [][2]int{}

	//create a robot
	paintrobot := robot.Robot{Direction: "NORTH", Position: [2]int{size / 2, size / 2}}

	for {
		inputs, counter = intCodeProgram(inputs, counter)

		if counter == 0 {
			break
		}

		if len(inputs) == 2 {
			x := paintrobot.Position[0]
			y := paintrobot.Position[1]
			panels[x][y] = paintrobot.PaintPanel(inputs[0])
			painted_panels = append(painted_panels, [2]int{x, y})
			num_panels++

			for i := 0; i < len(painted_panels)-1; i++ {

				if painted_panels[i] == painted_panels[len(painted_panels)-1] { //panel is already painted!
					RemoveIndex2(painted_panels, len(painted_panels)-1)
					num_panels--
					break
				}

			}

			paintrobot.TurnRobot(inputs[1])
			paintrobot.Move()

			if isBlack(panels[paintrobot.Position[0]][paintrobot.Position[1]]) {
				inputs = []int{0}
			} else { //white
				inputs = []int{1}
			}
		}
	}

	PrintImage(panels)

}

func PrintImage(newimage [200][200]int) {
	width := len(newimage[0])
	height := len(newimage)

	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	// Set color for each pixel.
	for x := 0; x < height; x++ {
		for y := 0; y < width; y++ {
			if newimage[x][y] == 0 {
				img.Set(x, y, color.Black)
			}
			if newimage[x][y] == 1 {
				img.Set(x, y, color.White)
			}
		}
	}

	// Encode as PNG.
	f, _ := os.Create("image.png")
	png.Encode(f, img)
}

func isBlack(colour int) bool {
	return colour == 0
}

func generateInstruction(instr string) string {
	for len(instr) < 5 {
		instr = "0" + instr
	}
	return instr
}

func setParam(result int, counter int, mode string) {
	if mode == "0" {
		program[program[counter]] = result
	}

	if mode == "1" {
		program[counter] = result
	}

	program[relativeBase+program[counter]] = result
}

func getParam(counter int, mode string) int {

	if mode == "0" {
		return program[program[counter]]
	}

	if mode == "1" {
		return program[counter]
	}

	if mode == "2" {
		return program[program[counter]+relativeBase]
	}

	fmt.Println("Should not come here")
	return program[counter]

}

func intCodeProgram(inputs []int, counter int) ([]int, int) {
	var output_signal int

	for {
		var first_param int
		var second_param int
		var result int

		instruction := generateInstruction(strconv.Itoa(program[counter]))
		opcode := string(instruction[3:5])

		mode1 := string(instruction[2])
		mode2 := string(instruction[1])
		mode3 := string(instruction[0])

		switch opcode {
		case "99": //exit ampControl, part of 99
			counter = 0
			return inputs, counter

		case "01":
			first_param = getParam(counter+1, mode1)
			second_param = getParam(counter+2, mode2)
			result = first_param + second_param
			setParam(result, counter+3, mode3)

			counter += 4
			break

		case "02":
			first_param = getParam(counter+1, mode1)
			second_param = getParam(counter+2, mode2)
			result = first_param * second_param
			setParam(result, counter+3, mode3)

			counter += 4
			break

		case "03":
			setParam(inputs[0], counter+1, mode1)
			inputs = RemoveIndex(inputs, 0)
			counter += 2
			break

		case "04":
			output_signal = getParam(counter+1, mode1)
			inputs = append(inputs, output_signal)
			counter += 2
			return inputs, counter

		case "05":
			first_param = getParam(counter+1, mode1)
			second_param = getParam(counter+2, mode2)
			if first_param != 0 {
				counter = second_param
			} else {
				counter += 3
			}
			break

		case "06":
			first_param = getParam(counter+1, mode1)
			second_param = getParam(counter+2, mode2)
			if first_param == 0 {
				counter = second_param
			} else {
				counter += 3
			}
			break
		case "07":
			first_param = getParam(counter+1, mode1)
			second_param = getParam(counter+2, mode2)
			if first_param < second_param {
				result = 1
			} else {
				result = 0
			}
			setParam(result, counter+3, mode3)
			counter += 4
			break

		case "08":
			first_param = getParam(counter+1, mode1)
			second_param = getParam(counter+2, mode2)
			if first_param == second_param {
				result = 1
			} else {
				result = 0
			}
			setParam(result, counter+3, mode3)
			counter += 4
			break
		case "09":
			first_param = getParam(counter+1, mode1)
			relativeBase += first_param
			counter += 2
			break
		}

	}
}

func RemoveIndex(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}

func RemoveIndex2(s [][2]int, index int) [][2]int {
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

	for len(nums) < 2000 {
		nums = append(nums, 0)
	}

	return nums
}
