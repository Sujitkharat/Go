package array

// reverse reverses elements in arr[start:end+1]
func reverse(arr []int, start, end int) {
	for start < end {
		arr[start], arr[end] = arr[end], arr[start]
		start++
		end--
	}
}

// RotateArray rotates the array to the right by k steps
func RotateArray(arr []int, k int) []int {
	n := len(arr)
	if n == 0 || k%n == 0 {
		return arr
	}

	k = k % n // Correct modulus logic

	// Reverse the three segments
	reverse(arr, 0, n-k-1)
	reverse(arr, n-k, n-1)
	reverse(arr, 0, n-1)

	return arr
}
