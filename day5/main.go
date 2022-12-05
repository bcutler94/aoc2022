package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

type Crate map[int][]string

func main () {
	fd, err := os.Open("./instructions.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(fd)
	scanner.Split(bufio.ScanLines)

	crates := getCrates()

	for scanner.Scan() {
		line := scanner.Text()
		arr := strings.Split(line, " ")
		count, err := strconv.Atoi(arr[1])
		if err != nil {
			panic(err)
		}
		from, err := strconv.Atoi(arr[3])
		if err != nil {
			panic(err)
		}
		to, err := strconv.Atoi(arr[5])
		if err != nil {
			panic(err)
		}
		crates = movePart2(crates, count, to - 1, from - 1)
	}
	fmt.Println(crates)
	fmt.Println(getFinalStr(crates))
}

func getFinalStr(crates Crate) string {
	keys := make([]int, 0, len(crates))
	for key := range crates {
		keys = append(keys, key)
	}

	sort.Ints(keys)

	var final string = ""
	for key := range keys {
		final += crates[key][len(crates[key]) - 1]
	}
	return final
}



func move(crates Crate, count int, to int, from int) Crate {
	fromLen := len(crates[from])
	for i := 1; i <= count; i++ {
		val := crates[from][fromLen - i]
		crates[to] = append(crates[to], val)
	}
	crates[from] = crates[from][:fromLen - count]
	return crates
}

func movePart2(crates Crate, count int, to int, from int) Crate {
	fromLen := len(crates[from])
	sliceToMove := crates[from][fromLen - count:]
	crates[to] = append(crates[to], sliceToMove...)
	crates[from] = crates[from][:fromLen - count]
	return crates
}



func getCrates() Crate {
	fd, err := os.Open("./crates.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(fd)
	scanner.Split(bufio.ScanLines)
	crates := make(Crate)
	for scanner.Scan() {
		line := scanner.Text()
		arr := strings.Split(line, "")
		for index, ele := range arr {
			mapIndex := (index - 1) / 4
			validIndex := (index - 1) % 4
			isLetter := IsLetter(ele)
			if validIndex == 0 && isLetter {
				crates[mapIndex] = append([]string{ele}, crates[mapIndex]...)
			}
		}
		// fmt.Println("*****************************************")
	}
	return crates
}

func IsLetter(s string) bool {
	for _, r := range s {
			if !unicode.IsLetter(r) {
					return false
			}
	}
	return true
}

