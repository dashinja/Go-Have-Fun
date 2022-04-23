package main

func Operation(operationType func(...int) int, args ...int) (z int) {
	if GetFunctionName(operationType) == "main.Square" {
		println("Function name is: ", GetFunctionName(operationType))

		z = Square(args[0])
		return z
	}
	z = operationType(args...)
	return z
}

//TODO: What can I do with the below line?
type operationFuncMulti func(...int) int
// type operationFunc func(int) int
// type operationType interface {
// 	operationFunc
// 	operationFuncMulti
// }

// type operationList interface {
// 	func(...int) int
// 	func(int) int
// }

func Adder(args ...int) int {
	z := 0
	for _, v := range args {
		z += v
	}
	return z
}

func Subtract(args ...int) int {
	z := 0
	for _, v := range args {
		z -= v
	}
	return z
}

func Multiply(args ...int) int {
	z := 1
	for _, v := range args {
		z *= v
	}
	return z
}

func Divide(args ...int) int {
	z := 1
	for i, v := range args {
		if i == 0 {
			z = v
			continue
		}
		z /= v
	}
	return z
}

func Max(args ...int) int {
	z := 0
	for _, v := range args {
		if v > z {
			z = v
		}
	}
	return z
}

func Min(args ...int) int {
	z := args[len(args)-1]
	for _, v := range args {
		if z > v {
			z = v
		}
	}
	return z
}

func Square(arg int) int {
	return arg * arg
}

var operationAdd5 = func(args ...int) int {
	sliceOf5 := []int{5}
	appendedSlice := append(args, sliceOf5...)
	return Operation(Adder, Adder(appendedSlice...))
}

var OperationSlice = []operationFuncMulti{
	Adder,
	Subtract,
	Multiply,
	Divide,
	Max,
	Min,
	operationAdd5,
}
