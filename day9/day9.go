package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	fd, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(fd)
	scanner.Split(bufio.ScanLines)

	grid := [][]int{{1}}
	h := []int{0, 0}
	t := []int{0, 0}
	for scanner.Scan() {
		line := scanner.Text()

		arr := strings.Split(line, " ")

		steps, err := strconv.Atoi(arr[1])
		if err != nil {
			panic(err)
		}

		switch arr[0] {
		case "R":
			MoveR(&grid, &h, &t, steps)
		case "L":
			MoveL(&grid, &h, &t, steps)
		case "D":
			MoveD(&grid, &h, &t, steps)
		case "U":
			MoveU(&grid, &h, &t, steps)
		default:
			panic("no move available")
		}
	}
	fmt.Println(CountMoves(&grid))
	// 5519
	// 5527
	// 5523
}

func PrintGrid(grid *[][]int) {
	for j := range *grid {
		fmt.Println((*grid)[len(*grid) - j - 1])
	}
}

func CountMoves(grid *[][]int) int {
	count := 0
	gridPointer := *grid
	for j := range gridPointer {
		for i := range gridPointer[j] {
			if gridPointer[j][i] == 1 {
				count++
			}
		}
	}
	return count
}

func FillGrid(grid *[][]int) {
	maxI := 0
	for j := range *grid {
		if len((*grid)[j]) > maxI {
			maxI = len((*grid)[j])
		}
	}
	
	for j := range *grid {
		diff := maxI - len((*grid)[j])
		for i := 0; i < diff; i++ {
			(*grid)[j] = append((*grid)[j], 0)
		}
	}
}

func AdjustTail(grid *[][]int, h, t *[]int) {


	if math.Abs(float64((*h)[0]-(*t)[0])) <= 1 && math.Abs(float64((*h)[1]-(*t)[1])) <= 1 {
		(*grid)[(*t)[1]][(*t)[0]] = 1
	} else {
		// if too far left
		if (*h)[0]-(*t)[0] == 2 {
			(*t)[1] = (*h)[1]
			(*t)[0] = (*h)[0] - 1
		}
	
		// if too far right
		if (*t)[0]-(*h)[0] == 2 {
			(*t)[1] = (*h)[1]
			(*t)[0] = (*h)[0] + 1
		}
	
		// if too high
		if (*t)[1]-(*h)[1] == 2 {
			(*t)[1] = (*h)[1] + 1
			(*t)[0] = (*h)[0]
		}
	
		// if too low
		if (*h)[1]-(*t)[1] == 2 {
			(*t)[1] = (*h)[1] - 1
			(*t)[0] = (*h)[0]
		}
	
		(*grid)[(*t)[1]][(*t)[0]] = 1
	}


}

func MoveR(grid *[][]int, h, t *[]int, steps int) {
	for i := 1; i <= steps; i++ {
		// Move Head
		length := len((*grid)[(*h)[1]])
		if 1+(*h)[0] >= length {
			(*grid)[(*h)[1]] = append((*grid)[(*h)[1]], 0)
		} 
		(*h)[0]++
		AdjustTail(grid, h, t)
		FillGrid(grid)
	}
}

func MoveL(grid *[][]int, h, t *[]int, steps int) {
	for i := 1; i <= steps; i++ {
		// Move Head
		if (*h)[0]-1 < 0 {
			(*grid)[(*h)[1]] = append([]int{0}, (*grid)[(*h)[1]]...)
			(*h)[0] = 0
			(*t)[0]++
		} else {
			(*h)[0]--
		}
		AdjustTail(grid, h, t)
		FillGrid(grid)
	}
}

func MoveU(grid *[][]int, h, t *[]int, steps int) {
	for i := 1; i <= steps; i++ {
		// Move Head
		length := len(*grid)
		if 1+(*h)[1] >= length {
			row := []int{}
			for k := 0; k <= (*h)[0]; k++ {
				row = append(row, 0)
			}
			(*grid) = append((*grid), row)
		} 
		(*h)[1]++
		AdjustTail(grid, h, t)
		FillGrid(grid)
	}
}

func MoveD(grid *[][]int, h, t *[]int, steps int) {
	for i := 1; i <= steps; i++ {
		// Move Head
		if (*h)[1]-1 < 0 {
			row := []int{}
			for k := 0; k <= (*h)[0]; k++ {
				row = append(row, 0)
			}
			(*grid) = append([][]int{row}, *grid...)
			(*h)[1] = 0
			(*t)[1]++
		} else {
			(*h)[1]--
		}
		AdjustTail(grid, h, t)
		FillGrid(grid)
	}
}


