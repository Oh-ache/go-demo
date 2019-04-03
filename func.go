package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		return a / b, nil
	default:
		return 0, fmt.Errorf("unspport op " + op)
	}
}

func div(a, b int) (q, r int) {
	return a / b, a % b
}

func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("call function %s with args (%d, %d)\n", opName, a, b)
	return op(a, b)
}

//func pow (a, b int) int {
//	return int(math.Pow(float64(a), float64(b)))
//}

func sum(numbers ...int) int {
	sum := 0
	for i := range numbers {
		sum += i
	}
	return sum
}

func swap(a, b *int) {
	*a, *b = *b, *a
}

func main() {
	if res, err := eval(10, 3, ":"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
	q, r := div(13, 4)
	fmt.Println(q, r)
	fmt.Println(apply(func(a, b int) int {
		return int(math.Pow(float64(a), float64(b)))
	}, 3, 4))
	fmt.Println(sum(1, 2, 3, 4, 5, 6, 8))
	a, b := 1, 3
	swap(&a, &b)
	fmt.Println(a, b)
}
