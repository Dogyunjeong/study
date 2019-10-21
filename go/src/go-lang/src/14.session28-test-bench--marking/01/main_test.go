package main

import "testing"

func TestMySum(t *testing.T) {

	type test struct {
		data   []int
		answer int
	}

	tests := []test{
		test{[]int{22, 20}, 42},
		test{[]int{11, 20, 5}, 36},
		test{[]int{-1, 8, 20}, 27},
		test{[]int{3, 4}, 7},
	}

	for _, v := range tests {
		x := mySum(v.data...)
		if x != v.answer {
			t.Error("Expected ", v.answer, "Got ", x)
		}
	}
}
