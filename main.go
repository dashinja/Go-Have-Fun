package main

//TODO: https://go101.org/article/101.html is YOUR FRIEND
//NOTE: https://go.dev/doc/faq
//NOTE: https://go.dev/blog/slices-intro

import (
	"fmt"
)

func main() {

	answer := Adder(1,2,3,4,5)
	fmt.Print(answer)
	inputs := []int {1,2,3,4,5}
	complexAnswer := Operation(Adder, inputs...)
	fmt.Println("complexAnswer: ", complexAnswer)
	fmt.Println("wow answer: ", operationAdd5(1,2,3,4,5,6,7,8))
	fmt.Println("wow answer dot dot dot: ", operationAdd5(inputs...))
}

func Operation(operationType func(...int) int , args ...int) (z int) {
	z = operationType(args...)
	return z
}
// func Operation(operationType func([]int) int, args ...int) (int z) {
// 	z := operationType(args)
// 	return z
// }
var operationAdd5 = func (args ...int) int {
	sliceOf5 := []int {5}
	appendedSlice := append(args, sliceOf5...)
	return Operation(Adder, Adder(appendedSlice...))
}

func Adder(args ...int) int {
	z := 0
	for _, v := range args {
		z += v
	}
	return z
}

func Subtract(args []int) int {
	z := 0
	for _, v := range args {
		z -= v
	}
	return z
}

func multiply(args []int) int {
	z := 0
	for _, v := range args {
		z *= v
	}
	return z
}

func divide(args []int) int {
	z := 0
	for _, v := range args {
		z /= v
	}
	return z
}
