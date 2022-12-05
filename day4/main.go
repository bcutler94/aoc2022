package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main () {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}

	scan := bufio.NewScanner(file)
	scan.Split(bufio.ScanLines)
	
	count := 0
	for scan.Scan() {
		text := scan.Text()
		pair := strings.Split(text, ",")
		leftPair := strings.Split(pair[0], "-")
		rightPair := strings.Split(pair[1], "-")
		leftPairInt := convertToI(leftPair)
		rightPairInt := convertToI(rightPair)
		if leftPairInt[0] <= rightPairInt[0] && leftPairInt[1] >= rightPairInt[0] {
			count++
		} else if leftPairInt[0] <= rightPairInt[1] && leftPairInt[1] >= rightPairInt[1] {
			count++
		} else if rightPairInt[0] <= leftPairInt[0] && rightPairInt[1] >= leftPairInt[0] {
			count++
		} else if rightPairInt[0] <= leftPairInt[1] && rightPairInt[1] >= leftPairInt[1] {
			count++
		}
	}
	fmt.Println(count)
}

func convertToI(arr []string) []int {
	res := make([]int, 0)
	for _, ele := range arr {
		val, err := strconv.Atoi(ele)
		if err != nil {
			panic(err)
		}
		res = append(res, val)
	}
	return res
}