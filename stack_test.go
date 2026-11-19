package main

import "testing"

func TestIsValid(t *testing.T) {
	tests := []struct{
		input string
		want bool
	}{
		{"()", true},
		{"(]", false},
		{"()[]{}", true},
		{"([])[{}]", true},
		{")(", false},
	}

	for _, tc := range tests {
		got := isValid(tc.input)
		if got != tc.want {
			t.Errorf("isValid(%q) = %v, want %v", tc.input, got, tc.want)
		}
	}
}

func TestMaxDepth(t *testing.T){
	tests := []struct{
		input string
		want int
	}{
		{"(1+(2*3)+((8)/4))+1", 3},
		{"(1)+((2))+(((3)))", 3},
		{"()(())((()()))", 3},
		{"((1))+(2)", 2},
		{"123", 0},
	}

	for _, tc := range tests {
		got := maxDepth(tc.input)
		if got != tc.want {
			t.Errorf("maxDepth(%q) = %v, want %v", tc.input, got, tc.want)
		}
	}
}

func TestRemoveOuterParentheses (t *testing.T){
	tests := []struct{
		input string
		want string
	}{
		{"(()())(())", "()()()"},
		{"(()())(())(()(()))", "()()()()(())"},
		{"()()", ""},
		{"()", ""},
	}

	for _, tc := range tests {
		got := removeOuterParentheses(tc.input)

		if got != tc.want {
			t.Errorf("removeOuterParentheses(%q) = %q, want %q", tc.input, got, tc.want)
		}
	}
}