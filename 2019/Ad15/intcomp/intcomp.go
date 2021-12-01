package intcomp

import (
	"fmt"
	"os"
	"strconv"
)

type Computer struct {
	Program  []int
	Done     bool
	Rel_base int
	Counter  int
}

func New(Program []int) Computer {
	comp := Computer{Program, false, 0, 0}
	return comp
}

func generateInstruction(instr string) string {
	for len(instr) < 5 {
		instr = "0" + instr
	}
	return instr
}

func (comp *Computer) setParam(result int, mode string, pos int) {
	if mode == "0" {
		comp.Program[comp.Program[comp.Counter+pos]] = result
	}

	if mode == "1" {
		comp.Program[comp.Counter+pos] = result
	}

	comp.Program[comp.Rel_base+comp.Program[comp.Counter+pos]] = result
}

func (comp *Computer) getParam(mode string, pos int) int {
	value := comp.Program[comp.Counter+pos]
	if mode == "0" {
		return comp.Program[value]
	}
	if mode == "1" {
		return value
	}
	if mode == "2" {
		return comp.Program[value+comp.Rel_base]
	} else {
		fmt.Println("Should not come here")
		os.Exit(1)
		return 0
	}
}

func RemoveIndex(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}

func RemoveIndex2(s [][2]int, index int) [][2]int {
	return append(s[:index], s[index+1:]...)
}

func (comp *Computer) Compute(signal int) int {

	for {
		var first_param int
		var second_param int
		var result int

		instruction := generateInstruction(strconv.Itoa(comp.Program[comp.Counter]))
		opcode := string(instruction[3:5])

		mode1 := string(instruction[2])
		mode2 := string(instruction[1])
		mode3 := string(instruction[0])

		switch opcode {
		case "99": //exit ampControl, part of 99
			comp.Done = true
			return 0

		case "01":
			first_param = comp.getParam(mode1, 1)
			second_param = comp.getParam(mode2, 2)
			result = first_param + second_param
			comp.setParam(result, mode3, 3)
			comp.Counter += 4
			break

		case "02":
			first_param = comp.getParam(mode1, 1)
			second_param = comp.getParam(mode2, 2)
			result = first_param * second_param
			comp.setParam(result, mode3, 3)
			comp.Counter += 4
			break

		case "03":
			comp.setParam(signal, mode1, 1)
			comp.Counter += 2
			break

		case "04":
			comp.Counter += 2
			return comp.getParam(mode1, 1-2)

		case "05":
			first_param = comp.getParam(mode1, 1)
			second_param = comp.getParam(mode2, 2)
			if first_param != 0 {
				comp.Counter = second_param
			} else {
				comp.Counter += 3
			}
			break

		case "06":
			first_param = comp.getParam(mode1, 1)
			second_param = comp.getParam(mode2, 2)
			if first_param == 0 {
				comp.Counter = second_param
			} else {
				comp.Counter += 3
			}
			break
		case "07":
			first_param = comp.getParam(mode1, 1)
			second_param = comp.getParam(mode2, 2)
			if first_param < second_param {
				result = 1
			} else {
				result = 0
			}
			comp.setParam(result, mode3, 3)
			comp.Counter += 4
			break

		case "08":
			first_param = comp.getParam(mode1, 1)
			second_param = comp.getParam(mode2, 2)
			if first_param == second_param {
				result = 1
			} else {
				result = 0
			}
			comp.setParam(result, mode3, 3)
			comp.Counter += 4
			break
		case "09":
			first_param = comp.getParam(mode1, 1)
			comp.Rel_base += first_param
			comp.Counter += 2
			break
		}

	}
}
