package array

// UniqueOnly returns a slice with only elements that appear exactly once in the input.
func UniqueOnly[T comparable](arr []T) []T {
	frequencyMap := make(map[T]int)
	for _, val := range arr {
		frequencyMap[val]++
	}

	result := []T{}
	for _, val := range arr {
		if frequencyMap[val] == 1 {
			result = append(result, val)
		}
	}
	return result
}

// RemoveDuplicates returns a slice with duplicates removed, preserving order.
func RemoveDuplicates[T comparable](arr []T) []T {
	seen := make(map[T]bool)
	result := []T{}

	for _, val := range arr {
		if !seen[val] {
			seen[val] = true
			result = append(result, val)
		}
	}
	return result
}
