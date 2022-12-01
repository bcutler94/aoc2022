package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	input, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(input)
	fileScanner.Split(bufio.ScanLines)

	var max int = 0
	var current int = 0
	for fileScanner.Scan() {
		text := fileScanner.Text()
		if text == "" {
			if current > max {
				max = current
			}
			current = 0
			continue
		} else {
			num, err := strconv.Atoi(text)
			if err != nil {
				panic(err)
			}
			current += num
		}
	}
	fmt.Print(max)
}
