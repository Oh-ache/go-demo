package main

import "fmt"

func lengthOfNomReaptingSubStr(s string) int {
	lastOccurres := make(map[rune]int)
	start := 0
	maxLength := 0
	for i, ch := range []rune(s) {
		if lastI, ok := lastOccurres[ch]; ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurres[ch] = i
	}
	return maxLength
}

func main() {
	fmt.Println(lengthOfNomReaptingSubStr("123456"))
	fmt.Println(lengthOfNomReaptingSubStr("一二三二一"))
}
