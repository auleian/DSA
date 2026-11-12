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

func TestTwoSum(t *testing.T) {
	tests := []struct {
		input  []int
		target int
		want   []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
	}

	for _, tc := range tests {
		got := twoSum(tc.input, tc.target)
		if len(got) != len(tc.want) {
			t.Errorf("twoSum(%v, %d) = %v, want %v", tc.input, tc.target, got, tc.want)
			continue
		}
		for i := range tc.want {
			if got[i] != tc.want[i] {
				t.Errorf("twoSum(%v, %d) = %v, want %v", tc.input, tc.target, got, tc.want)
				break
			}
		}
	}
}

func TestIsAnagram(t *testing.T){
	tests := []struct{
		input_1 string
		input_2 string
		want bool
	}{
		{"anagram", "nagaram", true},
		{"rat", "car", false},
		{"ab", "a", false},
		{"", "", true},
		{"aacc", "ccac", false},
	}

	for _, tc := range tests {
		got := isAnagram(tc.input_1, tc.input_2)
		if got != tc.want {
			t.Errorf("isAnagram(%v, %v) = %v, want %v", tc.input_1, tc.input_2, got, tc.want)
		}
	}
}