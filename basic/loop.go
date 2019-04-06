package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
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

	printFileContents(file)
}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	fmt.Println(converTobin(10))
	fmt.Println(converTobin(0))
	ptintFile("basic/abc.txt")

	s := `qqq
www
eee
rrr`
	printFileContents(strings.NewReader(s))
}
