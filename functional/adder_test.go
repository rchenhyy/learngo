package main

import (
	"math"
	"testing"
)

func TestAdder(t *testing.T) {
	tests := []struct {
		arr []int
		sum int
	}{
		// normal cases
		{[]int{1, 2}, 3},
		{[]int{1, 2, 3, 4, 5}, 15},
		{[]int{9, 9, 9, 9, 9}, 45},

		// edge cases
		{[]int{math.MaxInt64, 1}, math.MinInt64},
	}

	for _, i := range tests {
		add := adder()
		sum := 0
		for _, v := range i.arr {
			sum = add(v)
		}

		if sum != i.sum {
			t.Errorf("got %d for arr %v, expected %d", sum, i.arr, i.sum)
		}
	}
}

func BenchmarkIntAdder(b *testing.B) {
	arr, expected := []int{math.MaxInt64, 1}, math.MinInt64

	for i := 0; i < b.N; i++ { // b.N
		add := adder()
		sum := 0
		for _, v := range arr {
			sum = add(v)
		}
		if sum != expected {
			b.Errorf("got %d for arr %v, expected %d", sum, arr, expected)
		}
	}
}
