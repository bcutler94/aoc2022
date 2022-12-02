package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main ()  {
	fmt.Println(getResult(1, 3))
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	var totalScore int = 0
	for fileScanner.Scan() {
		text := fileScanner.Text()
		arr := strings.Split(text, " ")
		leftScore := getScore(arr[0])
		rightScore := getScore(arr[1])
		fmt.Println(arr[0], arr[1])
		fmt.Println(leftScore, rightScore)
		fmt.Println(getResult(leftScore, rightScore))
		fmt.Println()
		totalScore += getResult(leftScore, rightScore)
	}
	fmt.Println(totalScore)
}

func getResult(left int, right int) int {
	if left == right {
		return right + 3
	} 
	var killer int = right + 1
	if killer == 4 {
		killer = 1
	}
	if left == killer {
		return right
	}
	return right + 6
}

func getScore(letter string) int {
	switch letter {
	case "A", "X":
		return 1
	case "B", "Y":
		return 2
	case "C", "Z":
		return 3
	}
	panic(letter)
}