package slicex

import (
	"cmp"
	"sort"

	"slices"
)

// MinK returns the k smallest elements in the slice, sorted in ascending order.
//
// The original slice will be modified.
func MinK[E cmp.Ordered](s []E, k int) []E {
	if k < 1 || len(s) == 0 {
		return nil
	}

	if k >= len(s) {
		slices.Sort(s)
		return s
	}

	mins := s[:k]
	i, max := Max(mins)

	for _, e := range s[k:] {
		if e < max {
			// swap out the biggest element with the new one
			mins[i] = e
			i, max = Max(mins)
		}
	}

	slices.Sort(mins)
	return mins
}

// TopK returns the k biggest elements in the slice, sorted in descending order.
//
// The original slice will be modified.
func TopK[E cmp.Ordered](s []E, k int) []E {
	if k < 1 || len(s) == 0 {
		return nil
	}

	// found that sort.Slice is faster then slices.SortFunc, see benchmarks

	if k >= len(s) {
		sort.Slice(s, func(i, j int) bool { return s[i] > s[j] })
		return s
	}

	top := s[:k]
	i, min := Min(top)

	for _, e := range s[k:] {
		if e > min {
			// swap out the smallest element with the new one
			top[i] = e
			i, min = Min(top)
		}
	}

	sort.Slice(top, func(i, j int) bool { return s[i] > s[j] })
	return top
}

// Min returns the position and value of the minimal element in s.
// It panics if s is empty.
func Min[E cmp.Ordered](s []E) (int, E) {
	if len(s) == 0 {
		panic("slicex.Min: empty slice")
	}

	i, min := 0, s[0]
	for j, e := range s {
		if e < min {
			i = j
			min = e
		}
	}
	return i, min
}

// Max returns the position and value of the maximal element in s.
// It panics if s is empty.
func Max[E cmp.Ordered](s []E) (int, E) {
	if len(s) == 0 {
		panic("slicex.Max: empty slice")
	}

	i, max := 0, s[0]
	for j, e := range s {
		if e > max {
			i = j
			max = e
		}
	}
	return i, max
}
