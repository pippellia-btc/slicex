package slicex

import (
	"cmp"
	"slices"
)

// Unique removes duplicates from the input slice in-place.
// The original slice will be modified (sorted and compacted).
func Unique[E cmp.Ordered](s []E) []E {
	slices.Sort(s)
	return slices.Compact(s)
}

// Intersection returns the elements present in all the provided slices.
// The original slices will be modified.
func Intersection[E cmp.Ordered](inputs ...[]E) []E {
	switch len(inputs) {
	case 0:
		return nil

	case 1:
		return Unique(inputs[0])

	default:
		// quick check for empty slices
		for _, s := range inputs {
			if len(s) == 0 {
				return nil
			}
		}

		// intersect two slices at the time
		inter := Unique(inputs[0])
		for _, s := range inputs[1:] {
			inter = intersection(inter, Unique(s))

			if len(inter) == 0 {
				return nil
			}
		}

		return inter
	}
}

// Union the provided slices into a single one, with no duplicates.
// The original slices will be modified.
func Union[E cmp.Ordered](inputs ...[]E) []E {
	var size int
	for _, s := range inputs {
		size += len(s)
	}

	if size == 0 {
		return nil
	}

	u := Unique(inputs[0])
	for _, s := range inputs[1:] {
		// append to u the unique elements in s-u
		u = append(u, Difference(Unique(s), u)...)
		slices.Sort(u)
	}
	return u
}

// Difference returns the elements in s1 not in s2, with no duplicates.
// The original slices will be modified.
func Difference[E cmp.Ordered](s1, s2 []E) []E {
	if len(s1) == 0 {
		return nil
	}

	if len(s2) == 0 {
		return Unique(s1)
	}

	return difference(Unique(s1), Unique(s2))
}

// Differences returns the Difference(s1,s2) and Difference(s2,s1).
// The original slices will be modified.
func Differences[E cmp.Ordered](s1, s2 []E) ([]E, []E) {
	if len(s1) == 0 {
		return nil, Unique(s2)
	}

	if len(s2) == 0 {
		return Unique(s1), nil
	}

	return differences(Unique(s1), Unique(s2))
}

// SymmetricDifference returns the elements in either s1 or s2 but not in both, with no duplicates.
// The original slices will be modified.
func SymmetricDifference[E cmp.Ordered](s1, s2 []E) []E {
	if len(s1) == 0 {
		return Unique(s2)
	}

	if len(s2) == 0 {
		return Unique(s1)
	}

	s1, s2 = differences(Unique(s1), Unique(s2))
	return append(s1, s2...)
}

// Partition the elements of s1 and s2 into three non-overlapping slices with no duplicates:
//   - u1: elements uniquely in s1 (not is s2).
//   - inter: elements in both s1 and s2.
//   - u2: elements uniquely in s2 (not in s1).
//
// The original slices will be modified.
func Partition[E cmp.Ordered](s1, s2 []E) (u1, inter, u2 []E) {
	if len(s1) == 0 {
		return nil, nil, Unique(s2)
	}

	if len(s2) == 0 {
		return Unique(s1), nil, nil
	}

	return partition(Unique(s1), Unique(s2))
}

// intersection returns the elements in both s1 and s2, with no duplicates.
// The original slices will be modified.
// Both slices are assumed to be already sorted with no duplicates.
func intersection[E cmp.Ordered](s1, s2 []E) []E {
	var read1, read2, write int
	for read1 < len(s1) && read2 < len(s2) {
		switch {
		case s1[read1] < s2[read2]:
			read1++

		case s1[read1] > s2[read2]:
			read2++

		default:
			// element in s1 and s2
			s1[write] = s1[read1]
			write++
			read1++
			read2++
		}
	}

	// remove leftover
	return s1[:write]
}

// difference returns s1-s2, with no duplicates.
// s1 will be modified.
// Both slices are assumed to be already sorted with no duplicates.
func difference[E cmp.Ordered](s1, s2 []E) []E {
	var read1, read2, write int
	for read1 < len(s1) && read2 < len(s2) {
		switch {
		case s1[read1] < s2[read2]:
			// element in s1 not in s2
			s1[write] = s1[read1]
			write++
			read1++

		case s1[read1] > s2[read2]:
			read2++

		default:
			// element in s1 and s2
			read1++
			read2++
		}
	}

	// add all remaining elements to s1
	for read1 < len(s1) {
		s1[write] = s1[read1]
		read1++
		write++
	}

	// remove leftover
	return s1[:write]
}

// differences returns s1-s2 and s2-s1, with no duplicates.
// Both slices will be modified.
// Both slices are assumed to be already sorted with no duplicates.
func differences[E cmp.Ordered](s1, s2 []E) ([]E, []E) {
	var read1, read2, write1, write2 int
	for read1 < len(s1) && read2 < len(s2) {
		switch {
		case s1[read1] < s2[read2]:
			// element in s1 not in s2
			s1[write1] = s1[read1]
			write1++
			read1++

		case s1[read1] > s2[read2]:
			// element in s2 not in s1
			s2[write2] = s2[read2]
			write2++
			read2++

		default:
			// element in s1 and s2
			read1++
			read2++
		}
	}

	// add all remaining elements to s1
	for read1 < len(s1) {
		s1[write1] = s1[read1]
		read1++
		write1++
	}

	// add all remaining elements to s2
	for read2 < len(s2) {
		s2[write2] = s2[read2]
		read2++
		write2++
	}

	// remove leftover
	return s1[:write1], s2[:write2]
}

// partition the elements of s1 and s2 into three non-overlapping slices with no duplicates:
//   - elements uniquely in s1 (not is s2).
//   - elements in both s1 and s2.
//   - elements uniquely in s2 (not in s1).
//
// The original slices will be modified.
// Both slices are assumed to be already sorted with no duplicates.
func partition[E cmp.Ordered](s1, s2 []E) ([]E, []E, []E) {
	var inter []E
	var read1, read2, write1, write2 int

	for read1 < len(s1) && read2 < len(s2) {
		switch {
		case s1[read1] < s2[read2]:
			// element in s1 not in s2
			s1[write1] = s1[read1]
			read1++
			write1++

		case s1[read1] > s2[read2]:
			// element in s2 not in s1
			s2[write2] = s2[read2]
			read2++
			write2++

		default:
			// element in s1 and s2
			inter = append(inter, s1[read1])
			read1++
			read2++
		}
	}

	// add all remaining elements to s1
	for read1 < len(s1) {
		s1[write1] = s1[read1]
		read1++
		write1++
	}

	// add all remaining elements to s2
	for read2 < len(s2) {
		s2[write2] = s2[read2]
		read2++
		write2++
	}

	// remove leftovers
	return s1[:write1], inter, s2[:write2]
}
