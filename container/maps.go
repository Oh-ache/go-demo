package main

import "fmt"

func main() {
	m := map[string]string{
		"name":    "zhaobo",
		"course":  "golang",
		"site":    "imooc",
		"quality": "not bad",
	}

	fmt.Println(m)

	m2 := make(map[string]int)

	var m3 map[string]int

	fmt.Println(m2, m3)

	for k, v := range m {
		fmt.Println(k, v)
	}
	null, ok := m["null"]
	fmt.Println(null, ok)
}
