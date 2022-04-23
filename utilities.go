package main

func Operation(operationType func(...int) int , args ...int) (z int) {
	z = operationType(args...)
	return z
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

func Multiply(args []int) int {
	z := 0
	for _, v := range args {
		z *= v
	}
	return z
}

func Divide(args []int) int {
	z := 0
	for _, v := range args {
		z /= v
	}
	return z
}