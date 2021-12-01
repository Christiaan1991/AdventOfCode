package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var data [][][]string = readFile("data.txt")

func main() {
	tril := math.Pow(10, 12)
	low := int(tril / findOreAmount(1, "FUEL"))
	high := 10 * low

	for findOreAmount(high, "FUEL") < tril { //rough adjustment
		low = high
		high = low * 10
		fmt.Println(findOreAmount(high, "FUEL"), low, high)
	}

	for low < (high - 1) {
		mid := (low + high) / 2
		num_ore := findOreAmount(mid, "FUEL")
		fmt.Println(findOreAmount(mid, "FUEL"), mid)
		if num_ore > tril {
			high = mid
		} else if num_ore < tril {
			low = mid
		} else { //exactly one trillion

			break
		}
	}

}

func findOreAmount(number int, comp string) float64 {

	var needed_chem = make(map[string]int)
	var avaible_chem = make(map[string]int)
	needed_chem[comp] = number //add to needed_chem list
	var num_reactions int
	var total_ore int = 0

	for len(needed_chem) != 0 { //while we still need chem

		item := getFirstValue(needed_chem)

		if needed_chem[item] <= avaible_chem[item] {
			avaible_chem[item] -= needed_chem[item]
			delete(needed_chem, item)
			continue
		}

		needed_item := needed_chem[item] - avaible_chem[item]
		delete(needed_chem, item)
		delete(avaible_chem, item)

		//check all reactions
		for i := 0; i < len(data); i++ {
			if item == data[i][len(data[i])-1][1] {
				produced_item, _ := strconv.Atoi(data[i][len(data[i])-1][0])

				if needed_item%produced_item == 0 {
					num_reactions = needed_item / produced_item
				} else {
					num_reactions = (needed_item / produced_item) + 1
				}

				avaible_chem[item] += (num_reactions * produced_item) - needed_item

				for j := 0; j < len(data[i])-1; j++ {
					toInt, _ := strconv.Atoi(data[i][j][0])
					if data[i][j][1] == "ORE" {
						total_ore += toInt * num_reactions
					} else {
						needed_chem[data[i][j][1]] += toInt * num_reactions
					}
				}
			}
		}
	}
	return float64(total_ore)
}

func getFirstValue(m map[string]int) string {
	i := 0
	for key, _ := range m {
		if i == 0 {
			return key
		}
	}
	return "0"
}

func readFile(fname string) [][][]string {

	var data [][][]string

	f, err := os.Open(fname)
	if err != nil {
		fmt.Printf("error opening file: %v\n", err)
		os.Exit(1)
	}
	r := bufio.NewReader(f)
	for {
		var line [][]string
		s, e := Readln(r)

		if e != nil {
			break
		}

		elements := strings.Split(s, " => ")

		components := strings.Split(elements[0], ", ")
		for _, w := range components {
			num := strings.Split(w, " ")
			line = append(line, num)
		}

		num := strings.Split(elements[1], " ")
		line = append(line, num)

		data = append(data, line)

	}

	return data

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
