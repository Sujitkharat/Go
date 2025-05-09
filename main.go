package main

import (
	"fmt"
	"practice/dsa/array"
)

func main() {
	ints := []int{1, 2, 3, 4, 2, 3, 5}
	strs := []string{"apple", "banana", "apple", "cherry", "banana", "date"}

	fmt.Println("Unique ints: ", array.UniqueOnly[int](ints))
	fmt.Println("Unique strings: ", array.RemoveDuplicates(strs))

	fmt.Println("Rotated Array: ", array.RotateArray(ints, 3))
}
