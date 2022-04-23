package main

import (
	"testing"
)

func TestAdder(t *testing.T) {
	type args struct {
		args []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Happy Path",
			args: args{Inputs},
			want: 2542,
		},
		{
			name: "With mixed Signed Ints",
			args: args{[]int{1, 2, 3, 4, -5}},
			want: 5,
		},
		{
			name: "With all negative Ints",
			args: args{[]int{-1, -2, -3, -4, -5}},
			want: -15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Adder(tt.args.args...); got != tt.want {
				t.Errorf("Adder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOperation(t *testing.T) {
	type args struct {
		operationType func(...int) int
		args          []int
	}
	tests := []struct {
		name  string
		args  args
		wantZ int
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotZ := Operation(tt.args.operationType, tt.args.args...); gotZ != tt.wantZ {
				t.Errorf("Operation() = %v, want %v", gotZ, tt.wantZ)
			}
		})
	}
}

func TestSubtract(t *testing.T) {
	type args struct {
		args []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Happy Path",
			args: args{Inputs},
			want: -2542,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Subtract(tt.args.args...); got != tt.want {
				t.Errorf("Subtract() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMultiply(t *testing.T) {
	type args struct {
		args []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Happy Path",
			args: args{Inputs},
			want: 200000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Multiply(tt.args.args...); got != tt.want {
				t.Errorf("Multiply() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDivide(t *testing.T) {
	type args struct {
		args []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Happy Path",
			args: args{ReverseInts(Inputs)},
			want: 31,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Divide(tt.args.args...); got != tt.want {
				t.Errorf("Divide() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMax(t *testing.T) {
	type args struct {
		args []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Happy Path",
			args: args{Inputs},
			want: 2500,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Max(tt.args.args...); got != tt.want {
				t.Errorf("Max() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMin(t *testing.T) {
	type args struct {
		args []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Happy Path",
			args: args{Inputs},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Min(tt.args.args...); got != tt.want {
				t.Errorf("Min() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSquare(t *testing.T) {
	type args struct {
		arg int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Happy Path",
			args: args{12},
			want: 144,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Square(tt.args.arg); got != tt.want {
				t.Errorf("Square() = %v, want %v", got, tt.want)
			}
		})
	}
}
