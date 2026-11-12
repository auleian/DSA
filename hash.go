//hashing (maps and sets)

package main

import (
	"fmt"
	"reflect"
)


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

//Exercise 3
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	countS := make(map[byte]int)
	countT := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		countS[s[i]]++
		countT[t[i]]++   
	}

	return reflect.DeepEqual(countS, countT)
}