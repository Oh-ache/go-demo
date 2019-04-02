package main

import (
	"fmt"
	"regexp"
)

const text = `
my email is theache666@gmail.com@aaaa
my email is 32423@gmail.com@aaaa
my email is 3242345@gmail.com@aaaa
my email is 5gfgf@gmail.com@aaaa
my email is 4325@gmail.com.aaaa`

func main() {
	re := regexp.MustCompile(
		`([a-zA-Z0-9]+)@([a-zA-Z0-9]+)(\.[a-zA-Z0-9.]+)`)
	match := re.FindAllStringSubmatch(text, -1)

	for _, m := range match {
		fmt.Println(m)
	}
}
