package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

type Submarine struct {
	Depth int
	Forward int
}

func (s *Submarine) MoveForward(input int) {
	s.Forward += input
}

func (s *Submarine) MoveUp(input int) {
	s.Depth -= input
}

func (s *Submarine) MoveDown(input int) {
	s.Depth += input
}

type NewSubmarine struct {
	Depth int
	Forward int
	Aim int
}

func (s *NewSubmarine) MoveForward(input int) {
	s.Forward += input
	s.Depth += s.Aim*input
}

func (s *NewSubmarine) MoveUp(input int) {
	s.Aim -= input
}

func (s *NewSubmarine) MoveDown(input int) {
	s.Aim += input
}

func main(){
	input := readFile("input")
	part1(input)
	part2(input)
}

func part1(input []string){
	sub := Submarine{0,0}

	for _, s := range input {
		move, num := readCommand(s)
		if move == "forward" {
			sub.MoveForward(num)
		} else if move == "down" {
			sub.MoveDown(num)
		} else if move == "up" {
			sub.MoveUp(num)
		}
	}
	fmt.Println(sub.Depth * sub.Forward)
}

func part2(input []string){
	sub := NewSubmarine{0,0, 0}

	for _, s := range input {
		move, num := readCommand(s)
		if move == "forward" {
			sub.MoveForward(num)
		} else if move == "down" {
			sub.MoveDown(num)
		} else if move == "up" {
			sub.MoveUp(num)
		}
	}
	fmt.Println(sub.Depth * sub.Forward)
}

func readCommand(input string) (string, int) {
	s := strings.Split(input, " ")
	i, err := strconv.Atoi(string(s[1]))
	if err != nil {
		panic(err)
	}

	return string(s[0]), i
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

