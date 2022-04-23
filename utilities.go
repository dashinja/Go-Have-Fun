package main

import (
	"reflect"
	"runtime"
)

func ReverseInts(input []int) []int {
	if len(input) == 0 {
		return input
	}
	return append(ReverseInts(input[1:]), input[0])
}

func GetFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}