package main

import "fmt"

func printArr(arr *[5]int) {
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func main() {
	var arr1 [5]int
	arr2 := [3]int{1, 2, 3}
	arr3 := [...]int{2, 3, 4, 5, 6}
	var grid [4][5]int
	fmt.Println(arr1, arr2, arr3)
	fmt.Println(grid)

	fmt.Println("print arr1")
	printArr(&arr1)

	fmt.Println("print arr3")
	printArr(&arr3)

	fmt.Println("print arr1 and arr3 ")
	fmt.Println(arr1, arr3)

}
