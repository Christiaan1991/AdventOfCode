package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	data := readFile()

	counter := findShortestPath(data, "YOU", "SAN")
	fmt.Println(counter)
}

func findOrbitToCOM(data [][2]string, i int, counter int) int {
	counter++
	if data[i][0] == "COM" {
		//fmt.Println("COM found!")
		return counter
	}

	for j := 0; j < len(data); j++ {
		if data[i][0] == data[j][1] { //find next orbital
			counter = findOrbitToCOM(data, j, counter)
		}
	}

	return counter
}

func findOrbit(data [][2]string, obj string) string {
	for j := 0; j < len(data); j++ {
		if data[j][1] == obj {
			return data[j][0]
		}
	}

	return "0"
}

func findShortestPath(data [][2]string, obj1 string, obj2 string) int {

	var obj1_list []string
	var obj2_list []string
	for {
		obj1_list = append(obj1_list, obj1)
		obj1 = findOrbit(data, obj1)
		if obj1 == "COM" {
			break
		}
	}

	for {
		obj2_list = append(obj2_list, obj2)
		obj2 = findOrbit(data, obj2)
		if obj2 == "COM" {
			break
		}
	}

	length1 := len(obj1_list)
	length2 := len(obj2_list)

	for {
		if obj1_list[length1-1] != obj2_list[length2-1] {
			distance := length1 + length2 - 2
			return distance
		}
		length1--
		length2--
	}

}

func readFile() [][2]string {

	file, err := os.Open("data.txt")
	var data [][2]string

	if err != nil {
		log.Fatal(err)
	}

	rd := bufio.NewReader(file)

	for {
		line, err := rd.ReadString('\n')
		line = strings.TrimSuffix(line, "\r\n")
		split := strings.Split(line, ")")

		data = append(data, [2]string{split[0], split[1]})

		if err != nil {
			if err == io.EOF {
				break
			}

			log.Fatalf("read file line error: %v", err)
			os.Exit(1)
		}

	}

	return data

}
