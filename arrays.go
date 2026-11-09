package main

func reverse(arr []int)  {
	left := 0
	right := len(arr) - 1

	for left < right {
		arr[left], arr[right] = arr[right], arr[left]
		left++
		right--
	}
}

func windowAverages(arr []int, k int) []float64 {
	result := []float64{}

    sum := 0
	for i := 0; i < k; i++ {
		sum += arr[i]
	}
	result = append(result, float64(sum) / float64(k))

	for i := k; i < len(arr); i++ {
		sum += arr[i]
		sum -= arr[i-k]
		result = append(result, float64(sum) / float64(k))
	}
	return result
}

