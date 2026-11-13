package main

import "testing"

func build(vals ...int) *ListNode {
	dummy := &ListNode{}
	current := dummy
	for _, v := range vals {
		current.Next = &ListNode{Val: v}
		current = current.Next
	}
	return dummy.Next
}

func toSlice(head *ListNode) []int {
	result := []int{}
	for current := head; current != nil; current = current.Next {
		result = append(result, current.Val)
	}
	return result
}

func TestReverseList(t *testing.T) {
	tests := []struct {
		input []int
		want  []int
	}{
		{[]int{10, 20, 30}, []int{30, 20, 10}},
		{[]int{1, 2}, []int{2, 1}},
		{[]int{5}, []int{5}},        
		{[]int{}, []int{}},          
	}

	for _, tc := range tests {
		list := build(tc.input...)        
		reversed := reverseList(list)      
		got := toSlice(reversed)           

		if len(got) != len(tc.want) {
			t.Errorf("reverseList(%v) = %v, want %v", tc.input, got, tc.want)
			continue
		}
		for i := range tc.want {
			if got[i] != tc.want[i] {
				t.Errorf("reverseList(%v) = %v, want %v", tc.input, got, tc.want)
				break
			}
		}
	}
}