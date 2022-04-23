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
		// TODO: Add test cases.
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
	type args struct {
		i interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetFunctionName(tt.args.i); got != tt.want {
				t.Errorf("GetFunctionName() = %v, want %v", got, tt.want)
			}
		})
	}
}
