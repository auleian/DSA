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