package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"os"
	"strconv"
	"strings"
)

func main()  {
	BuildGaggle()
}

type Monkey struct {	
	id    int
	items []int
	operation ast.Expr
	divisible int
	divTrue int
	divFalse int
}

func (monkey *Monkey) Init(input string) {
	arr := strings.Split(input, "\n")

	// ID
	id := arr[0][len(arr[0]) - 2]
	monkey.id = int(id)

	// Starting items
	strItems := strings.Split(arr[1][18:], ", ")
	startingItems := []int{}
	for _, item := range strItems {
		intItem, err := strconv.Atoi(item)
		if err != nil {
			panic(err)
		}
		startingItems = append(startingItems, intItem)
	}
	monkey.items = startingItems

	// Operation
	strExpr := strings.Split(arr[2], "=")
	expr, err := parser.ParseExpr(strExpr[1])
	if err != nil {
		panic(err)
	}
	monkey.operation = expr

	// Divisble
	strDivisble := strings.Split(arr[3], " ")
	divisible, err := strconv.Atoi(strDivisble[len(strDivisble) - 1])
	if err != nil {
		panic(err)
	}
	fmt.Println("Div", divisible)
	monkey.divisible = divisible

	// DivTrue
	strDivs := strings.Split(arr[4], " ")
	divTrue, err := strconv.Atoi(strDivs[len(strDivs) - 1])
	if err != nil {
		panic(err)
	}
	fmt.Println("DivTrue", divTrue)
	monkey.divTrue = divTrue

	// DivFalse
	strDivsF := strings.Split(arr[5], " ")
	divFalse, err := strconv.Atoi(strDivsF[len(strDivsF) - 1])
	if err != nil {
		panic(err)
	}
	fmt.Println("DivFalse", divFalse, strDivsF)
	monkey.divFalse = divFalse

	fmt.Println("Created monkey", monkey)
}

type Gaggle map[int]*Monkey

func (gaggle *Gaggle) AddMonkey(monkey *Monkey) {
 (*gaggle)[monkey.id] = monkey
}

func BuildGaggle() Gaggle {
	file, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	fileStr := string(file)
	arr := strings.Split(fileStr, "\n\n")
	gaggle := make(Gaggle)
	for _, item := range arr {
		monkey := new(Monkey)
		monkey.Init(item)
		gaggle.AddMonkey(monkey)
	}

	return gaggle
}