//Linked Lists

package main

type ListNode struct {
    Val  int
    Next *ListNode
}

func countNodes(head *ListNode) int {
    counter := 0
	current := head

	for current != nil {
		counter++
		current = current.Next
	}

	return counter
}

func reverseList(head *ListNode) *ListNode{
	var prev *ListNode = nil
	current := head

	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}

	return prev
}