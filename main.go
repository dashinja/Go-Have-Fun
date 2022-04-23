package main



import (
	"fmt"
)

func main() {

	answer := Adder(1,2,3,4,5)
	inputs := []int {1,2,3,4,5}
	complexAnswer := Operation(Adder, inputs...)
	
	fmt.Print(answer)
	fmt.Println("complexAnswer: ", complexAnswer)
	fmt.Println("wow answer dot dot dot: ", operationAdd5(inputs...))
}


var operationAdd5 = func (args ...int) int {
	sliceOf5 := []int {5}
	appendedSlice := append(args, sliceOf5...)
	return Operation(Adder, Adder(appendedSlice...))
}

