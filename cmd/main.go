package main

import (
	"cmp"
	"fmt"
	"slices"
)

// For slices.Sort

// This function attempts to unique in place without returning the slice
// but it won't change the caller's slice length.
func UniqueInPlaceNoReturn[E cmp.Ordered](s []E) {
	if len(s) == 0 {
		return
	}
	slices.Sort(s)
	writeIndex := 1
	for i := 1; i < len(s); i++ {
		if s[i] != s[writeIndex-1] {
			s[writeIndex] = s[i]
			writeIndex++
		}
	}
	// This line updates the local 's' variable's length,
	// but the caller's variable still has its original length.
	s = s[:writeIndex] // THIS CHANGE IS NOT SEEN BY THE CALLER'S SLICE VARIABLE
}

func main() {
	mySlice := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3}
	fmt.Println("Original slice:", mySlice, "Length:", len(mySlice), "Capacity:", cap(mySlice))

	UniqueInPlaceNoReturn(mySlice) // Call the function

	// Observe the result: Elements are changed, but length is not for the caller!
	fmt.Println("After UniqueInPlaceNoReturn:", mySlice, "Length:", len(mySlice), "Capacity:", cap(mySlice))
	// Expected output for mySlice: [1 2 3 4 5 6 9 6 5 3] (elements changed, but length still 10)
	// The "unique" part is only up to index 6 (length 7: [1 2 3 4 5 6 9])
}
