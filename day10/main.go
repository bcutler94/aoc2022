package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main()  {
	fd, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(fd)
	scanner.Split(bufio.ScanLines)

	crt := new(CRT)
	crt.Init(40, 6)

	for scanner.Scan() {

		line := scanner.Text()
		arr := strings.Split(line, " ")

		switch (arr[0]) {
		case "noop":
			crt.Noop()
		case "addx":
			x, err := strconv.Atoi(arr[1])
			if err != nil {
				panic(err)
			}
			crt.Addx(x)
		}

	}
	crt.Print()
}

type CRT struct {
	ui [][]string 
	sprite int
	cycle int
}

func (crt *CRT) Draw() {
	col := (crt.cycle - 1) % 40
	row := (crt.cycle - 1) / 40
	overlap := crt.sprite + 1 >= col && crt.sprite - 1 <= col
	fmt.Println("Col IDX: ", col)
	fmt.Println("Row IDX: ", row)
	fmt.Println("Cycle: ", crt.cycle)
	fmt.Println("Sprite: ", crt.sprite)
	fmt.Println()
	if overlap {
		crt.ui[row][col] = "#"
	}
}

func (crt *CRT) Noop() {
	crt.Draw()
	crt.cycle++
}

func (crt *CRT) Addx(x int) {
	crt.Draw()
	crt.cycle++
	crt.Draw()
	crt.cycle++
	crt.sprite += x
}

func (crt *CRT) Init(len, height int) {
	crt.sprite = 1 // idx
	crt.cycle = 1 // non-zero based
	for j := 0; j < height; j++ {
		row := []string{}
		for i := 0; i < len; i++ {
			row = append(row, ".")
		}
		crt.ui = append(crt.ui, row)
	}
}

func (crt *CRT) Print() {
	fmt.Println(crt.cycle)
	fmt.Println(crt.sprite)
	for i := range crt.ui {
		fmt.Println(crt.ui[i])
	}
}

// func main()  {
// 	fd, err := os.Open("./input.txt")
// 	if err != nil {
// 		panic(err)
// 	}

// 	scanner := bufio.NewScanner(fd)
// 	scanner.Split(bufio.ScanLines)

// 	cycles := 0
// 	val := 1
// 	pts := GetPoints(20, 220, 40)
// 	total := 0

// 	for scanner.Scan() {

// 		line := scanner.Text()
// 		arr := strings.Split(line, " ")

// 		switch (arr[0]) {
// 		case "noop":
// 			cycles++
// 			total += CalcStrength(cycles, val, pts)
// 		case "addx":
// 			cycles++
// 			total += CalcStrength(cycles, val, pts)
// 			cycles++
// 			total += CalcStrength(cycles, val, pts)
// 			x, err := strconv.Atoi(arr[1])
// 			if err != nil {
// 				panic(err)
// 			}
// 			val += x
// 		}

// 	}
// 	fmt.Println(total)
// 	// 13140
// }

func CalcStrength(cycle, x int, pts map[int]int) int {
	_, exists := pts[cycle]
	if exists {
		fmt.Println(cycle, x, pts)
		return x * cycle
	}
	return 0
}

func GetPoints(start, end, interval int) map[int]int {
	pts := map[int]int{}
	for i := start; i <= end; i += interval {
		pts[i] = 1
	}
	return pts
}