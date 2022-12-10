package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Knot struct {
	x, y int
	knot *Knot
	tail bool
}

func (c *Knot) ToKey() string {
	return strconv.Itoa(c.x) + "." + strconv.Itoa(c.y)
}

func BuildKnot() *Knot {
	knot := new(Knot)
	knot.x = 0
	knot.y = 0
	return knot
}

func BuildRope(count int) *Knot {
	rope := BuildKnot()
	currentKnot := rope
	for i := 1; i <= count; i++ {
		currentKnot.knot = BuildKnot()
		currentKnot = currentKnot.knot
		currentKnot.tail = false
	}
	currentKnot.tail = true
	return rope
}

func main() {
	fd, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(fd)
	scanner.Split(bufio.ScanLines)

	moves := map[string]int{ "0.0": 1 }
	rope := BuildRope(9)

	
	for scanner.Scan() {
		line := scanner.Text()

		arr := strings.Split(line, " ")

		steps, err := strconv.Atoi(arr[1])
		if err != nil {
			panic(err)
		}

		for i := 1; i <= steps; i++ {

			// Move head
			switch arr[0] {
			case "R":
				rope.x++
			case "L":
				rope.x--
			case "D":
				rope.y--
			case "U":
				rope.y++
			default:
				panic("no move available")
			}

			head := rope
			tail := rope.knot

			for {
				if head.tail {
					break
				}

				xDelta := head.x - tail.x
				yDelta := head.y - tail.y
	
				// Move tail
				if xDelta == 2 {
					tail.y = head.y
					tail.x = head.x - 1
				} else if xDelta == -2 {
					tail.y = head.y
					tail.x = head.x + 1
				} else if yDelta == 2 {
					tail.y = head.y - 1
					tail.x = head.x
				} else if yDelta == -2 {
					tail.y = head.y + 1
					tail.x = head.x
				}

				head = tail
				tail = head.knot
			}
			moves[head.ToKey()]++
		}
	}
	fmt.Println(len(moves))

	// 5519
	// 5527
	// 5523

			// 2083
		// 2259

}

	// 5710

// func PrintGrid(grid *[][]int) {
// 	for j := range *grid {
// 		fmt.Println((*grid)[len(*grid) - j - 1])
// 	}
// }

// func CountMoves(grid *[][]int) int {
// 	count := 0
// 	gridPointer := *grid
// 	for j := range gridPointer {
// 		for i := range gridPointer[j] {
// 			if gridPointer[j][i] == 1 {
// 				count++
// 			}
// 		}
// 	}
// 	return count
// }

// func FillGrid(grid *[][]int) {
// 	maxI := 0
// 	for j := range *grid {
// 		if len((*grid)[j]) > maxI {
// 			maxI = len((*grid)[j])
// 		}
// 	}
	
// 	for j := range *grid {
// 		diff := maxI - len((*grid)[j])
// 		for i := 0; i < diff; i++ {
// 			(*grid)[j] = append((*grid)[j], 0)
// 		}
// 	}
// }

// func AdjustTail(grid *[][]int, h, t *[]int) {


// 	if math.Abs(float64((*h)[0]-(*t)[0])) <= 1 && math.Abs(float64((*h)[1]-(*t)[1])) <= 1 {
// 		(*grid)[(*t)[1]][(*t)[0]] = 1
// 	} else {
// 		// if too far left
// 		if (*h)[0]-(*t)[0] == 2 {
// 			(*t)[1] = (*h)[1]
// 			(*t)[0] = (*h)[0] - 1
// 		}
	
// 		// if too far right
// 		if (*t)[0]-(*h)[0] == 2 {
// 			(*t)[1] = (*h)[1]
// 			(*t)[0] = (*h)[0] + 1
// 		}
	
// 		// if too high
// 		if (*t)[1]-(*h)[1] == 2 {
// 			(*t)[1] = (*h)[1] + 1
// 			(*t)[0] = (*h)[0]
// 		}
	
// 		// if too low
// 		if (*h)[1]-(*t)[1] == 2 {
// 			(*t)[1] = (*h)[1] - 1
// 			(*t)[0] = (*h)[0]
// 		}
	
// 		(*grid)[(*t)[1]][(*t)[0]] = 1
// 	}


// }

// func MoveR(grid *[][]int, h, t *[]int, steps int) {
// 	for i := 1; i <= steps; i++ {
// 		// Move Head
// 		length := len((*grid)[(*h)[1]])
// 		if 1+(*h)[0] >= length {
// 			(*grid)[(*h)[1]] = append((*grid)[(*h)[1]], 0)
// 		} 
// 		(*h)[0]++
// 		AdjustTail(grid, h, t)
// 		FillGrid(grid)
// 	}
// }

// func MoveL(grid *[][]int, h, t *[]int, steps int) {
// 	for i := 1; i <= steps; i++ {
// 		// Move Head
// 		if (*h)[0]-1 < 0 {
// 			(*grid)[(*h)[1]] = append([]int{0}, (*grid)[(*h)[1]]...)
// 			(*h)[0] = 0
// 			(*t)[0]++
// 		} else {
// 			(*h)[0]--
// 		}
// 		AdjustTail(grid, h, t)
// 		FillGrid(grid)
// 	}
// }

// func MoveU(grid *[][]int, h, t *[]int, steps int) {
// 	for i := 1; i <= steps; i++ {
// 		// Move Head
// 		length := len(*grid)
// 		if 1+(*h)[1] >= length {
// 			row := []int{}
// 			for k := 0; k <= (*h)[0]; k++ {
// 				row = append(row, 0)
// 			}
// 			(*grid) = append((*grid), row)
// 		} 
// 		(*h)[1]++
// 		AdjustTail(grid, h, t)
// 		FillGrid(grid)
// 	}
// }

// func MoveD(grid *[][]int, h, t *[]int, steps int) {
// 	for i := 1; i <= steps; i++ {
// 		// Move Head
// 		if (*h)[1]-1 < 0 {
// 			row := []int{}
// 			for k := 0; k <= (*h)[0]; k++ {
// 				row = append(row, 0)
// 			}
// 			(*grid) = append([][]int{row}, *grid...)
// 			(*h)[1] = 0
// 			(*t)[1]++
// 		} else {
// 			(*h)[1]--
// 		}
// 		AdjustTail(grid, h, t)
// 		FillGrid(grid)
// 	}
// }


