package main

import "fmt"

func updateSlice(arr []int) {
	arr[0] = 100
}

func printSlice(arr []int, s string) {
	fmt.Printf("%s=%v len(%s)=%d cap(%s)=%d \n", s, arr, s, len(arr), s, cap(arr))
}
func appendSlice(arr []int) []int {
	res := append(arr, 1)
	printSlice(res, "res")
	return res
}

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println("arr[2:6] is  ", arr[2:6])
	fmt.Println("arr[:6] is  ", arr[:6])
	fmt.Println("arr[2:] is  ", arr[2:])
	fmt.Println("arr[:] is  ", arr[:])
	fmt.Println("-------------------")
	s1 := arr[2:]
	updateSlice(s1)
	fmt.Println("after update s1 is", s1)
	s2 := arr[:]
	updateSlice(s2)
	fmt.Println("after update s2 is", s2)

	fmt.Println("-------------------")
	s3 := arr[2:6]
	s4 := s3[3:5]
	printSlice(arr[:], "arr")
	printSlice(s3, "s3")
	printSlice(s4, "s4")

	fmt.Println("-------------------")
	fmt.Println(arr)
	s5 := append(s4, 15)
	s6 := append(s5, 16)
	s7 := append(s6, 17)
	printSlice(s4, "s4")
	printSlice(s5, "s5")
	printSlice(s6, "s6")
	printSlice(s7, "s7")
	fmt.Println(arr)

	fmt.Println("-------------------")
	for i := 0; i < 20; i++ {
		s7 = appendSlice(s7)
	}

}
