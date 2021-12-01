package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

const row_size = 20
const col_size = 20
var allkeys = generateKeyList('a', 'p')

func main() {
	var input = readFile("data.txt")
	var data = input[:len(input)-1]
	dist := shortestPath(data)
	fmt.Println(dist)

}

func shortestPath(data [][]string) int {
	var rowNum = []int{-1, 0, 0, 1}
	var colNum = []int{0, -1, 1, 0}
	x, y := findPosition(data,'@')

	var q []queueNode
	var currNode queueNode
	var pt Point
	var visited []Point
	visited = append(visited, Point{x, y})
	var collectedKeys []rune
	data[x][y] = "."

	//append this node to the list of nodes
	q = append(q, queueNode{Point{x, y}, 0, visited, collectedKeys})

	for len(q) != 0 {
		currNode, q = q[0], q[1:]
		fmt.Println(currNode)
		pt = currNode.pt
		visitedPoints := currNode.visited

		for i := 0; i < 4; i++ {
			//we check all directions
			var row = pt.x + rowNum[i]
			var col = pt.y + colNum[i]
			var step = data[row][col]

			if step == "#" || hasVisited(row, col, visitedPoints) || (isDoor(step) && !hasKey(currNode.collectedKeys, []rune(step)[0])) {
				//if next point is a wall (#) or we already visited the point, or it's a door and we dont have the key, we cannot pass
				continue
			} else if isKey(step) && !keyInList([]rune(step)[0], currNode.collectedKeys) {
				//if next point is a key, and it is not yet in out keylist, we add the key to our keylist and reset all visited points
				currNode.collectedKeys = append(currNode.collectedKeys, []rune(step)[0])
				fmt.Println(string(currNode.collectedKeys), len(currNode.collectedKeys), len(allkeys))
				visitedPoints = nil
				//check if we found all keys
				if len(currNode.collectedKeys) == len(allkeys) {
					//all keys found in this node, so we return distance of this node!
					fmt.Println(string(currNode.collectedKeys))
					return currNode.dist+1
				}

			}

			//else, we can go to point: add point to visited points, and append to q
			visitedPoints = append(visitedPoints, Point{row, col})
			q = append(q, queueNode{Point{row, col}, currNode.dist+1, visitedPoints, currNode.collectedKeys})
		}
	}

	//we check all directions: if it is an #, we block that direction. If it is a ., we add the potential point to the Point list can go that direction, and remember the move we performed. For the next move we want to do, we cannot go backwards
	// If the move goes onto a keouy, we remove the door from the data map (A -> .). For the next move, backwards is allowed!
	err := errors.New("Cannot find all keys, so cannot find end of the maze!")
	panic(err)
	return 0

}

func generateKeyList(startKey rune, endKey rune) []rune {
	var keys []rune

	for key:=startKey; key <= endKey; key++ {
		keys = append(keys, key)
	}

	return keys
}

func printGrid(data [][]string, node queueNode) {
	for row:=0; row < len(data); row++{
		for col:=0; col < len(data[row]); col++{
			if row == node.pt.x && col == node.pt.y {
				fmt.Print("@")
			} else {
				fmt.Print(data[row][col])
			}

		}
		fmt.Print("\n")
	}
}

func hasKey(keys []rune, door rune) bool {
	for _, key := range keys {
		if key - 32 == door {
			return true
		}
	}
	return false
}

func hasVisited(row int, col int, visitedPoints []Point) bool {
	for _, visitedPoint := range visitedPoints {
		if visitedPoint.x == row && visitedPoint.y == col {
			return true
		}
	}
	return false
}

func keyInList(a rune, list []rune) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func hasDoor(data [][]string, door rune) bool {
	for row:=0; row < len(data); row++ {
		for col:=0; col < len(data[row]); col++ {
			rune_door := []rune(data[row][col])[0]
			if rune_door == door {
				//door is in the map!
				return true
			}
		}
	}
	return false
}

func isDoor(input string) bool {
	r := []rune(input)[0]
	if r < 'A' || r > 'Z' {
		return false
	}
	return true
}

func isKey(input string) bool {
	r := []rune(input)[0]
	if r < 'a' || r > 'z' {
		return false
	}
	return true
}

func isValid(row int, col int) bool {
	return (row >= 0) && (row < row_size) && (col >= 0) && (col < col_size)
}

func findPosition(data [][]string, symbol rune) (int, int){
	for row:=0; row < len(data); row++ {
		for col:=0; col < len(data[row]); col++ {
			if data[row][col] == string(symbol) {
				return row, col
			}
		}
	}
	err := errors.New("Starting/key position not found!")
	panic(err)
	return 0,0
}

func readFile(fname string) [][]string {

	var data [][]string

	f, err := os.Open(fname)
	if err != nil {
		fmt.Printf("error opening file: %v\n", err)
		os.Exit(1)
	}
	r := bufio.NewReader(f)
	for {
		s, e := Readln(r)
		elements := strings.Split(s, "")
		data = append(data, elements)

		if e != nil {
			break
		}
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

type Point struct {
	x, y int
}

type queueNode struct {
	pt 			  Point
	dist          int
	visited       []Point
	collectedKeys []rune
}




