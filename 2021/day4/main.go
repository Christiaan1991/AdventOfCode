package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

type Field struct {
	num int
	filled bool
}

type BingoCard struct {
	fields [5][5]Field
}

func (f *Field) fill (i int) {
	f.num = i
	f.filled = false
}

func (b *BingoCard) fillBingoCard(nums [][]int) {
	for i, _ := range nums {
		for j, _ := range nums[i] {
			b.fields[i][j].fill(nums[i][j])
		}
	}
}

func main(){
	input := readFile("input")
	//part1(input)
	part2(input)
}

func part2(input []string) {
	numbers := findNumbers(input)
	bingoCards := getBingoCards(input)


	for _, i := range numbers {
		//fmt.Println(i)
		findBingoNumber(bingoCards, i)
		for num:=0; num < len(bingoCards); num++{
			card := bingoCards[num]
			if isBingo(card) {
				if len(bingoCards) == 1 {
					getScore(card, i)
					return
				}
				//fmt.Printf("removing bingoCard: %d\n", num)
				bingoCards = remove(bingoCards, num)
				num = -1 //restart for loop
			}
		}
	}
}

func part1(input []string) {
	numbers := findNumbers(input)
	bingoCards := getBingoCards(input)
	for _, i := range numbers {
		findBingoNumber(bingoCards, i)
		for num, card := range bingoCards {
			if isBingo(card) {
				getScore(bingoCards[num], i)
				break
			}
		}
	}
}

func remove(s []BingoCard, i int) []BingoCard {
	return append(s[:i], s[i+1:]...)
}

func getScore(card BingoCard, num int) {
	var count int = 0
	for i:=0; i < 5; i++ {
		for j:=0; j < 5; j++ {
			if card.fields[i][j].filled == false {
				count+=card.fields[i][j].num
			}
		}
	}
	fmt.Printf("Final score: %d x %d = %d\n", count, num, count * num)
}

func checkVertical(card BingoCard) bool {
	for i:=0; i < 5; i++ {
		for j:=0; j < 5; j++ {
			if card.fields[j][i].filled == false {
				break
			}
			if j == 4 {
				fmt.Println("Vertical Bingo!")
				return true
			}
		}
	}
	return false
}

func checkHorizontal(card BingoCard) bool {
	for i:=0; i < 5; i++ {
		for j:=0; j < 5; j++ {
			if card.fields[i][j].filled == false {
				break
			}
			if j == 4 {
				fmt.Println("Horizontal Bingo!")
				return true //horizontal bingo!
			}
		}
	}
	return false
}

func isBingo(card BingoCard) bool {
	if checkHorizontal(card) || checkVertical(card) {
		return true
	}
	return false
}

func findBingoNumber(cards []BingoCard, num int) {
	for i:=0; i < len(cards); i++ {
		for j:=0; j < 5; j++ {
			for k:=0; k < 5; k++ {
				if cards[i].fields[j][k].num == num {
					cards[i].fields[j][k].filled = true
				}
			}
		}
	}
}

func findNumbers(input []string) []int {
	n := strings.Split(input[0], ",")
	var nums []int
	for _, i := range n {
		in, _ := strconv.Atoi(i)
		nums = append(nums, in)
	}
	return nums
}

func getBingoCards(input []string) []BingoCard {
	var cards []BingoCard
	var i int = 2
	var lines [][]int
	card := BingoCard{}
	for i < len(input) {
		if isEmpty(input[i]) {
			card.fillBingoCard(lines) // fill in Bingocard
			cards = append(cards, card) //append to list of bingo cards

			card = BingoCard{} //create new bingo card
			lines = nil
			i++
			continue
		}
		line := toIntArray(input[i])
		lines = append(lines, line)

		i++
	}

	return cards
}

func toIntArray(in string) []int {
	var nums []int
	strs := strings.Split(in, " ")
	for _, str := range strs {
		bytes := []byte(str)
		if len(bytes) == 0 {
			continue
		}
		i, _ := strconv.Atoi(str)

		nums = append(nums, i)
	}
	return nums
}

func isEmpty(in string) bool {
	if in == ""{
		return true
	}
	return false
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