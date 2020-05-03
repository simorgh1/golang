package main

import (
	"fmt"
)

func main() {

	// a sorted array with duplicate elements
	var nums []int = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4, 5, 6, 6, 7, 7, 8}
	// nums := [...]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}

	fmt.Printf("Sorted array: %d, size:%d\n", nums, len(nums))

	removeDuplicates(nums)
	rotate(nums, 5)
}

// A sorted array with duplicated items is rearranged and printed as a sorted unique members with new size
func removeDuplicates(x []int) {

	i := 0

	for j := 1; j < len(x); j++ {
		if x[j] != x[i] {
			i++
			x[i] = x[j]
		}
	}

	// the new size which contains the unique sorted items is i+1
	fmt.Printf("Cleaned: %d, size:%d, cap:%d\n", x[0:i+1], len(x[0:i+1]), cap(x[0:i+1]))
}

// It rotates the array by k items and prints the result
func rotate(x []int, k int) {

	for j := 0; j < k; j++ {
		last := x[len(x)-1]

		for l := len(x) - 2; l >= 0; l-- {
			x[l+1] = x[l]
		}

		// replace the first items with the last item after rotation
		x[0] = last
	}

	fmt.Printf("Rotated array: %d, by %d elements.\n", x, k)
}
