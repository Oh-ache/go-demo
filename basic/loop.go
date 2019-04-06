package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func converTobin(n int) string {
	if n == 0 {
		return "0"
	}
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func ptintFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scaner := bufio.NewScanner(file)

	for scaner.Scan() {
		fmt.Println(scaner.Text())
	}
}

func main() {
	fmt.Println(converTobin(10))
	fmt.Println(converTobin(0))
	ptintFile("abc.txt")
}
