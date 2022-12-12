package main

import (
	"fmt"
	"os"
	"strings"
)

func main()  {
	BuildGaggle()
}

Starting items: 

type Monkey struct {	
	id    byte
	items []int
	operation string
	factor int
}

func (monkey *Monkey) Init(input string) {
	arr := strings.Split(input, "/n")

	// ID
	id := arr[0][len(arr[0]) - 2]
	monkey.id = id

	// 
	strItems := strings.Split(arr[1][15:], ",")
	

}

type Gaggle map[int]*Monkey

func (gaggle *Gaggle) AddMonkey(monkey *Monkey) {
 (*gaggle)[monkey.id] = monkey
}

func BuildGaggle() *Gaggle {
	file, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	fileStr := string(file)
	arr := strings.Split(fileStr, "\n\n")
	gaggle := new(Gaggle)
	for _, item := range arr {
		monkey := new(Monkey)
		monkey.Init(item)
		gaggle.AddMonkey(monkey)
	}

	fmt.Println(arr, len(arr))
	return gaggle
}