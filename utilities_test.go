package main

import (
	"reflect"
	"testing"
)

func TestReverseInts(t *testing.T) {
	type args struct {
		input []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "should reverse a slice of ints",
			args: args{[]int{1, 2, 3, 4}},
			want: []int{4, 3, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseInts(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReverseInts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetFunctionName(t *testing.T) {
	// type args struct {
	// 		func(int) int
	// }

	type args struct {
		args func (int) int
	}
	tests := []struct {
		name string
		args operationFuncMulti
		want string
	}{
		{
			name: "should produce function name",
			args: Adder,
			want: "example.Adder",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetFunctionName(tt.args); got != tt.want {
				t.Errorf("GetFunctionName() = %v, want %v", got, tt.want)
			}
		})
	}
}
