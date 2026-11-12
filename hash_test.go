package main

import "testing"

func TestContainsDuplicate(t *testing.T){
	tests := []struct{
		input []int
		want bool
	}{
		{[]int{1, 3, 2, 6, 4}, false},
		{[]int{1, 2, 3, 1}, true},
		{[]int{1, 1, 1}, true},
		{[]int{1}, false},
		{[]int{}, false},
	}

	for _, tc := range tests {
		got := containsDuplicate(tc.input)

		if got != tc.want {
			t.Errorf("containsDuplicate(%v) = %v, want %v", tc.input, got, tc.want)
		}
	}
}