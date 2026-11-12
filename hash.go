//hashing (maps and sets)

package main

import "fmt"


func counter() {
	word := "banana"
	count := make(map[byte]int)

	for i := 0; i < len(word); i++ {
		letter := word[i]
		count[letter]++
	}

	fmt.Println(count)
}

//Exercise 1
func containsDuplicate(nums []int) bool {
	seen := make(map[int]bool)

	for _, i := range nums {
		if seen[i] {
			return true
		}
		seen[i] = true
	}

	return false
}