// arrays_test.go
package main
import "testing"

func TestReverse(t *testing.T) {
    arr := []int{1, 2, 3}
    reverse(arr)
    want := []int{3, 2, 1}
    for i := range arr {
        if arr[i] != want[i] {
            t.Errorf("got %v, want %v", arr, want)
        }
    }
}

func TestWindowAverages(t *testing.T) {
    tests := []struct {
        name string
        arr  []int
        k    int
        want []float64
    }{
        {"basic", []int{1, 3, 2, 6, 4}, 2, []float64{2, 2.5, 4, 5}},
        {"k equals 1", []int{5, 10}, 1, []float64{5, 10}},
        {"whole array", []int{2, 4}, 2, []float64{3}},
    }

    for _, tc := range tests {
        got := windowAverages(tc.arr, tc.k)
        for i := range tc.want {
            if got[i] != tc.want[i] {
                t.Errorf("%s: index %d got %v, want %v", tc.name, i, got[i], tc.want[i])
            }
        }
    }
}

func TestIsPalindrome(t *testing.T) {
    tests := []struct {
        input string
        want  bool
    }{
        {"race a car", false},
        {"abba", true},
        {"", true},
        {"A man, a plan, a canal: Panama", true},
    }

    for _, tc := range tests {
        got := isPalindrome(tc.input)
        if got != tc.want {
            t.Errorf("isPalindrome(%q) = %v, want %v", tc.input, got, tc.want)
        }
    }
}

func TestLongestUnique(t *testing.T){
    tests := []struct {
        input string
        want int
    }{
        {"abcabcbb", 3},
        {"bbbbb", 1},
        {"pwwkew", 3}, 
        {"", 0},
        {"abcdef", 6},
        {"a", 1}, 
        {"au", 2}, 
        {"dvdf", 3}, 
        {"tmmzuxt", 5},
    }

    for _, tc := range tests {
        got := longestUnique(tc.input)
        if got != tc.want {
            t.Errorf("longestUnique(%q) = %v, want %v", tc.input, got, tc.want)
        }
    }
}