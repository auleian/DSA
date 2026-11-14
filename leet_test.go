package main

import "testing"

func TestSwap(t *testing.T){
	tests := []struct{
		input []int
		want []int
	}{
		{[]int{3,9,2,5}, []int{9,9,5,-1}},
	}

	for _, tc := range tests {
		got := swap(tc.input)
		if len(got) != len(tc.want) {
			t.Fatalf("got %v, want %v", got, tc.want)
		}
		for i := range tc.want {
			if got[i] != tc.want[i] {
				t.Errorf("got %v, want %v", got, tc.want)
				break
			}
		}
	}

}