package main

import (
	"fmt"
)

var Inputs = []int{2, 40, 2500}

func main() {

	reverseInputs := ReverseInts(Inputs)
	fmt.Printf("inputs: %d\n", Inputs)
	fmt.Printf("reverseInputs %d\n: ", reverseInputs)

	for _, v := range OperationSlice {
		var res = 0
		var inputUsed = Inputs

		if GetFunctionName(v) == "main.Divide" {
			println("Function name is 'Divide' so reverse the input order for division:")

			res = Operation(v, reverseInputs...)
			inputUsed = reverseInputs
		} else {
			res = Operation(v, Inputs...)
		}
		fmt.Printf("%s(%d): %v\n\n", GetFunctionName(v),inputUsed, res)
	}
}


