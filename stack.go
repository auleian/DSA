package main

func isValid(s string) bool {
    pairs := map[rune]rune{')': '(', ']': '[', '}': '{'}
	stack := []rune{}

    for _, parenthese := range s {
		if open, isCloser := pairs[parenthese]; isCloser {
			if len(stack) == 0 || stack[len(stack)-1] != open {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, parenthese)
		}
	}
	return len(stack) == 0
}