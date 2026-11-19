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

func maxDepth(s string) int {
	max := 0
	open := 0

	for _, char := range s {
		switch char {
			case '(':
				open++
				if open > max {
					max = open
				}
			case ')':
				open--
			default:
				continue
		}
	}

	return max
}

func removeOuterParentheses(s string) string