package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	const filename = "abc.txt"
	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}
	fmt.Println(eval(1, 2, "+"))
	fmt.Println(getGrade(70))
}

func eval(a, b int, op string) int {
	var result int
	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	default:
		panic("unsupport operate " + op)
	}
	return result
}

func getGrade(sorce int) string {
	g := ""
	switch {
	case (sorce < 60 && sorce >= 0):
		g = "F"
	case sorce < 70:
		g = "D"
	case sorce < 80:
		g = "C"
	case sorce < 90:
		g = "B"
	case sorce <= 100:
		g = "A"
	default:
		panic("please get right sorce")
	}
	return g
}
