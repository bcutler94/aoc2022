package main

import (
	"bufio"
	"fmt"
	"os"
)

func main () {
	fd, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(fd)
	scanner.Split(bufio.ScanWords)

	fmt.Println(GetIndex("mjqjpqmgbljsphdztnvjfqwrcgsmlb", 14))

	for scanner.Scan() {
		word := scanner.Text()
		fmt.Println(GetIndex(word, 14))
		return
	}

	panic("No input")

}

func GetIndex(word string, len int) int {
	for idx, _ := range word {
		set := make(map[rune]int)
		foundIndex := true
		for _, subChar := range word[idx:idx+len] {
			_, exists := set[subChar]
			if exists {
				foundIndex = false
				break
			} else {
				set[subChar] = 1
			}
		}
		if foundIndex {
			return idx + len
		}
	}
	panic("not found")
}