package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	grid := BuildGrid()
	fmt.Println(GetMaxScore(grid))
}

func GetMaxScore(grid [][]int) int {
	max := 0
	for i := 1; i < len(grid) - 1; i++ {
		for j := 1; j < len(grid[i]) - 1; j++ {
			candidate := GetScore(grid, i, j)
			if candidate > max {
				max = candidate
			} 
		}
	}
	return max
}

func GetScore(grid [][]int, i, j int) int {
	treeHeight := grid[i][j]

	// Check up
	upScore := 0
	for p := -1; i + p >= 0; p-- {
		upScore++
		height := grid[i + p][j]
		if height >= treeHeight {
			break;
		} 
	}

	// Check down
	downScore := 0
	for p := 1; p + i < len(grid); p++ {
		downScore++
		height := grid[i + p][j]
		if height >= treeHeight {
			break
		}
	}

	// Check left
	leftScore := 0
	for p := -1; p + j >= 0; p-- {
		leftScore++
		height := grid[i][j + p]
		if height >= treeHeight {
			break
		}
	}

	// Check right
	rightScore := 0
	for p := 1; p + j < len(grid[0]); p++ {
		rightScore++
		height := grid[i][j + p]
		if height >= treeHeight {
			break
		}
	}

	return upScore * downScore * leftScore * rightScore
}

func CountVisible(grid [][]int) int {
	count := len(grid[0]) + len(grid) + len(grid[0]) + len(grid) - 4
	for i := 1; i < len(grid) - 1; i++ {
		for j := 1; j < len(grid[i]) - 1; j++ {
			if IsVisible(grid, i, j) {
				count++
			} 
		}
	}
	return count
}

func IsVisible(grid [][]int, i, j int) bool {
	treeHeight := grid[i][j]

	// Check up
	leftVisible := true
	for p := -1; i + p >= 0; p-- {
		height := grid[i + p][j]
		if height >= treeHeight {
			leftVisible = false
			break
		}
	}

	if leftVisible {
		return true
	}

	// Check down
	rightVisible := true
	for p := 1; p + i < len(grid); p++ {
		height := grid[i + p][j]
		if height >= treeHeight {
			rightVisible = false
			break
		}
	}

	if rightVisible {
		return true
	}

		// Check left
	upVisible := true
	for p := -1; p + j >= 0; p-- {
		height := grid[i][j + p]
		if height >= treeHeight {
			upVisible = false
			break
		}
	}

	if upVisible {
		return true
	}

	// Check right
	downVisible := true
	for p := 1; p + j < len(grid[0]); p++ {
		height := grid[i][j + p]
		if height >= treeHeight {
			downVisible = false
			break
		}
	}

	return downVisible
}

func BuildGrid() [][]int {
	fs, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(fs);
	scanner.Split(bufio.ScanLines);

	grid := make([][]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		arr := strings.Split(line, "")
		gridLine := make([]int, 0)
		for _, tree := range arr {
			num, err := strconv.Atoi(tree)
			if err != nil {
				panic(err)
			}
			gridLine = append(gridLine, num)
		}
		grid = append(grid, gridLine)
	}
	return grid
}