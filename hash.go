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

//Exercise 2
func twoSum(nums []int, target int) []int {
	seen := make(map[int]int)
	
	for i, x := range nums {
		compliment := target - x
		if j, ok := seen[compliment]; ok {
			return []int{j, i}
		}
		seen[x] = i
	}

	return nil
}

func isAnagram(s string, t string) bool