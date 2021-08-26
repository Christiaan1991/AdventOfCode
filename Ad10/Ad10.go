package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strings"
)

func main() {
	data := readFile("data.txt")
	CeresMonitoringStation(data)

}

func CeresMonitoringStation(data [][]string) {

	var asteroids [][2]float64
	var coords [2]float64
	var location int

	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data[i]); j++ {
			if isAsteroid(data[i][j]) {
				asteroids = append(asteroids, [2]float64{float64(i), float64(j)})
			}
		}
	}

	var tracked_asteroids [353][][2]float64
	var highest_num int = 0
	var angle float64
	var angles []float64
	var asteroids_angles [][3]float64

	for i := 0; i < len(asteroids); i++ {
		for j := i + 1; j < len(asteroids); j++ {
			if canSee(asteroids[i], asteroids[j], tracked_asteroids[i]) {
				tracked_asteroids[i] = append(tracked_asteroids[i], asteroids[j])
				tracked_asteroids[j] = append(tracked_asteroids[j], asteroids[i])
			}
		}

		if len(tracked_asteroids[i]) > highest_num {
			highest_num = len(tracked_asteroids[i])
			coords = asteroids[i]
			location = i
		}
	}

	fmt.Println(highest_num, coords)

	for j := 0; j < len(tracked_asteroids[location]); j++ {
		rel_x := tracked_asteroids[location][j][1] - coords[1]
		rel_y := tracked_asteroids[location][j][0] - coords[0]
		//fmt.Println(rel_x, rel_y)
		if rel_x == 0 && rel_y < 0 {
			angle = 0
		}

		if rel_x == 0 && rel_y > 0 {
			angle = 180
		}
		if rel_x > 0 && rel_y == 0 {
			angle = 90
		}

		if rel_x < 0 && rel_y == 0 {
			angle = 270
		}

		if rel_x > 0 && rel_y < 0 {
			angle = 360 * math.Atan(math.Abs(rel_x)/math.Abs(rel_y)) / (2 * math.Pi)
		}

		if rel_x > 0 && rel_y > 0 {
			angle = 90 + 360*math.Atan(math.Abs(rel_y)/math.Abs(rel_x))/(2*math.Pi)
		}

		if rel_x < 0 && rel_y > 0 {
			angle = 180 + 360*math.Atan(math.Abs(rel_x)/math.Abs(rel_y))/(2*math.Pi)
		}

		if rel_x < 0 && rel_y < 0 {
			angle = 270 + 360*math.Atan(math.Abs(rel_y)/math.Abs(rel_x))/(2*math.Pi)
		}
		angles = append(angles, angle)
		asteroids_angles = append(asteroids_angles, [3]float64{angle, tracked_asteroids[location][j][1], tracked_asteroids[location][j][0]})
		//fmt.Println(angle, tracked_asteroids[location][j][1], tracked_asteroids[location][j][0], rel_x, rel_y)
	}

	sort.Float64s(angles)
	//fmt.Println(angles)

	for i := 0; i < len(asteroids_angles); i++ {
		if angles[199] == asteroids_angles[i][0] {
			//fmt.Println(angles[2], asteroids_angles[i][1], asteroids_angles[i][2])
			fmt.Println(angles[199], asteroids_angles[i][1]*100+asteroids_angles[i][2])
		}
	}

	fmt.Printf("Number of asteroids detected: %d at position (%d, %d)", highest_num, int(coords[1]), int(coords[0]))
	//fmt.Println(angle)
}

func isAsteroid(pos string) bool { return pos == "#" }

func distance(a1 [2]float64, a2 [2]float64) float64 {
	return float64(math.Sqrt(math.Pow(a1[0]-a2[0], 2) + math.Pow(a1[1]-a2[1], 2)))
}

func canSee(ast1 [2]float64, ast2 [2]float64, tracked_asteroids [][2]float64) bool {
	if ast1 == ast2 {
		return false
	}
	for i := 0; i < len(tracked_asteroids); i++ {
		if ast1 == tracked_asteroids[i] || ast2 == tracked_asteroids[i] {
			continue
		}
		if math.Abs(distance(ast1, tracked_asteroids[i])+distance(ast2, tracked_asteroids[i])-distance(ast1, ast2)) < 0.000001 {
			return false
		}
	}
	return true
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
