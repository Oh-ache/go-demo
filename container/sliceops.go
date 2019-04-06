package main

import "fmt"

func printSlice(arr []int) {
	fmt.Printf("arr is %v\n", arr)
	fmt.Printf("len is %d, cap is %d\n", len(arr), cap(arr))
}

func main() {
	var arr []int
	for i := 0; i < 20; i++ {
		arr = append(arr, i)
		printSlice(arr)
	}

	s1 := []int{2, 3, 4, 5}
	fmt.Println(s1)

	s2 := make([]int, 9)
	fmt.Println(s2)
	fmt.Println(len(s2), cap(s2))

	s3 := make([]int, 10, 16)
	fmt.Println(s3)
	fmt.Println(len(s3), cap(s3))

	fmt.Println("---------------")
	copy(s2, s1)
	printSlice(s2)

	fmt.Println("---------------")
	s2 = append(s2[:3], s2[4:]...)
	printSlice(s2)

	fmt.Println("---------------")
	font := s2[0]
	s2 = s2[1:]

	end := s2[len(s2)-1]
	s2 = s2[:len(s2)-2]

	fmt.Println(s2, font, end, len(s2), cap(s2))
}
