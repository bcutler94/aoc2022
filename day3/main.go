package main

import (
	"bufio"
	"fmt"
	"os"
)

func main()  {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	reader := bufio.NewScanner(file)
	reader.Split(bufio.ScanLines)
	array := []string{}
	total := 0
	for reader.Scan() {
		array = append(array, reader.Text())
		if len(array) == 3 {
			total += getScore2(array[0], array[1], array[2])
			array = []string{}
		}
	}
	fmt.Println(total)
}

// func getScore(line string) int {
// 	length := len(line)
// 	half := length / 2
// 	halfMap := make(map[byte]int)
// 	for i := 0; i < half; i++ {
// 		charcode := int(line[i])
// 		if charcode < 96 {
// 			halfMap[line[i]] = charcode - 64 + 26
// 		} else {
// 			halfMap[line[i]] = charcode - 96
// 		}
// 	}
// 	for i := half; half < length; i++ {
// 		score, ok := halfMap[line[i]]
// 		if ok {
// 			return score
// 		}
// 	}

// 	panic("No match")
// }

func getScore2(line1 string, line2 string, line3 string) int {
	strMap := make(map[rune]int)
	for _, char := range line1 {
		strMap[char] = 1
	}
	for _, char := range line2 {
		_, ok := strMap[char]
		if ok {
			strMap[char] = 2
		}
	}
	for _, char := range line3 {
		num := strMap[char]
		if num == 2 {
			charcode := int(char)

			if charcode < 96 {
				return charcode - 64 + 26
			} else {
				return charcode - 96
			}
		}
	}

	panic("No match")
}