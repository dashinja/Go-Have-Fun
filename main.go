package main

import (
	"fmt"
)

func main() {

	inputs := []int {1000, 100000, 1000000000}
	reverseInputs := ReverseInts(inputs)
	fmt.Printf("inputs: %d\n", inputs)
	fmt.Printf("reverseInputs %d\n: ", reverseInputs)
	answer := Adder(inputs...)
	complexAnswer := Operation(Adder, inputs...)
	
	fmt.Println("Adder number is: ", answer)
	fmt.Println("Adder Equivalence: ", complexAnswer)
	fmt.Println("operationAdd5: ", operationAdd5(inputs...))
	fmt.Println("Subtract number is: ", Subtract(inputs...))
	fmt.Println("Multiply number is: ", Multiply(inputs...))
	fmt.Println("Divide number is: ", Divide(inputs...))
	fmt.Println("Divide reverse slice number is: ", Divide(reverseInputs...))
	fmt.Println("max number is: ", Max(inputs...))
	fmt.Println("min number is: ", Min(inputs...))
	fmt.Println("Square number is: ", Square(inputs[2]))
}


var operationAdd5 = func (args ...int) int {
	sliceOf5 := []int {5}
	appendedSlice := append(args, sliceOf5...)
	return Operation(Adder, Adder(appendedSlice...))
}

