package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

func main() {
	fmt.Println("hello go")
	variableZoreValue()
	euler()
	triangle()
	consts()
	enums()
}

func enums() {
	const (
		cpp = iota
		java
		python
		golang
		php
	)
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(cpp, java, python, golang, php)
	fmt.Printf("%d %d %d %d %d %d\n", b, kb, mb, gb, tb, pb)
}

func consts() {
	const filename = "abc.txt"
	const a, b = 3, 4
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)
}

func triangle() {
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Printf("%d\n", c)
}

func euler() {
	fmt.Printf("%.3f %.3f\n",
		cmplx.Exp(1i*math.Pi)+1,
		cmplx.Pow(math.E, 1i*math.Pi)+1)
}

func variableZoreValue() {
	var i int = 3
	var s string = "123"
	fmt.Printf("%d %q\n", i, s)
}
