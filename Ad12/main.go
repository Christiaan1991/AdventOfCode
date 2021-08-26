package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var pos [4][3]int
var vel [4][3]int

const NDIM = 3

var state [3][4][2]int
var init_state [3][4][2]int

var periods []int
var time int = 0

func main() {
	readFile("data.txt")

	for i := 0; i < NDIM; i++ {

		//every dimension seperately
		appendState(i)
		time = 0

		for {
			applyGravity(i)
			updatePosition(i)
			//fmt.Println(init_state[i], state[i])

			if check(i) {
				fmt.Println("copy found at", time+1)
				fmt.Println(init_state[i], state[i])
				periods = append(periods, time+1)
				break
			}
			time++

		}
	}

	fmt.Println(periods)

	answer := lcm(periods[0], periods[1])
	fmt.Println(answer)
	answer = lcm(answer, periods[2])
	fmt.Println(answer)

}

func gcd(x int, y int) int {
	if x > y {
		x, y = y, x
	}

	var i int
	for i = x; i >= 1; i-- {
		if x%i == 0 && y%i == 0 {
			break
		}
	}
	return i
}

func lcm(x int, y int) int {
	return x * y / gcd(x, y)
}

func appendState(i int) {
	for j := 0; j < 4; j++ {
		init_state[i][j][0] = pos[j][i]
		init_state[i][j][1] = vel[j][i]
	}
}

func check(dim int) bool {

	for i := 0; i < 4; i++ {
		state[dim][i][0] = pos[i][dim]
		state[dim][i][1] = vel[i][dim]
	}

	for i := 0; i < 4; i++ {
		for j := 0; j < 2; j++ {
			if init_state[dim][i][j] != state[dim][i][j] {
				return false
			}
		}
	}
	return true
}

func updatePosition(k int) {
	for i := 0; i < len(pos); i++ {
		pos[i][k] += vel[i][k]
	}
}

func applyGravity(k int) {
	for i := 0; i < len(pos); i++ {
		for j := i + 1; j < len(pos); j++ {
			if pos[i][k] > pos[j][k] {
				vel[i][k]--
				vel[j][k]++
			}
			if pos[i][k] < pos[j][k] {
				vel[i][k]++
				vel[j][k]--
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

func EnergyPrint() {
	fmt.Println("Energy Print:")
	var pot [4]int
	var kin [4]int
	var total [4]int
	for j := 0; j < len(pos); j++ {
		pot[j] = Abs(pos[j][0]) + Abs(pos[j][1]) + Abs(pos[j][2])
		kin[j] = Abs(vel[j][0]) + Abs(vel[j][1]) + Abs(vel[j][2])
		total[j] = pot[j] * kin[j]
		fmt.Printf("pot: %d + %d + %d = %d; kin: %d + %d + %d = %d; total: %d * %d = %d\n", Abs(pos[j][0]), Abs(pos[j][1]), Abs(pos[j][2]), pot[j], Abs(vel[j][0]), Abs(vel[j][1]), Abs(vel[j][2]), kin[j], pot[j], kin[j], total[j])
	}
	sum := total[0] + total[1] + total[2] + total[3]
	fmt.Printf("Sum of total energy: %d + %d + %d + %d = %d\n", total[0], total[1], total[2], total[3], sum)
}

func Print() {
	fmt.Printf("After %d steps:\n", time)
	for j := 0; j < len(pos); j++ {
		fmt.Printf("pos=<x=%d, y=%d, z=%d>, vel=<x=%d, y=%d, z=%d>\n", pos[j][0], pos[j][1], pos[j][2], vel[j][0], vel[j][1], vel[j][2])
	}
}

func readFile(fname string) {
	var i int = 0

	f, err := os.Open(fname)
	if err != nil {
		fmt.Printf("error opening file: %v\n", err)
		os.Exit(1)
	}
	r := bufio.NewReader(f)

	for {

		s, e := Readln(r)

		if e != nil {
			break
		}

		elements := strings.Split(s[1:len(s)-1], ", ")

		for j := 0; j < 3; j++ {
			num := elements[j][2:]
			pos[i][j], _ = strconv.Atoi(num)
		}
		i++
	}

}

func Readln(r *bufio.Reader) (string, error) {
	var (
		isPrefix bool  = true
		err      error = nil
		line, ln []byte
	)
	for isPrefix && err == nil {
		line, isPrefix, err = r.ReadLine()
		ln = append(ln, line...)
	}
	return string(ln), err
}
