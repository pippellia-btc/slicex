package slicex

// Unique returns a new slice with no duplicates, preserving the order of elements.
func Unique[E comparable](s []E) []E {
	seen := make(map[E]struct{}, len(s))
	unique := make([]E, 0, len(s))

	for _, e := range s {
		if _, exists := seen[e]; !exists {
			seen[e] = struct{}{}
			unique = append(unique, e)
		}
	}
	return unique
}

// Intersection returns a new slice containing unique elements found in all slices.
// The order of elements is the same as the one of the first slice.
func Intersection[E comparable](inputs ...[]E) []E {
	switch len(inputs) {
	case 0:
		return nil

	case 1:
		return Unique(inputs[0])

	default:
		// intersecting from the smallest set
		i, min := 0, len(inputs[0])
		for j, s := range inputs {
			if len(s) < min {
				i = j
				min = len(s)
			}
		}

		if min == 0 {
			return []E{}
		}

		interSet := toSet(inputs[i])
		for j, s := range inputs {
			if j == i {
				continue
			}

			current := toSet(s)
			for e := range interSet {
				if _, found := current[e]; !found {
					delete(interSet, e)
				}
			}

			if len(interSet) == 0 {
				return []E{}
			}
		}

		inter := make([]E, 0, len(interSet))
		for _, e := range inputs[0] {
			if _, found := interSet[e]; found {
				inter = append(inter, e)
				delete(interSet, e) // remove duplicates
			}
		}

		return inter
	}
}

// Union returns a new slice containing unique elements found in any of the slices.
// The order of elements reflects the one in the original slices.
func Union[E comparable](inputs ...[]E) []E {
	var size int
	for _, s := range inputs {
		size += len(s)
	}

	if size == 0 {
		return []E{}
	}

	union := make([]E, 0, size)
	seen := make(map[E]struct{}, size)

	for _, s := range inputs {
		for _, e := range s {
			if _, found := seen[e]; !found {
				union = append(union, e)
				seen[e] = struct{}{}
			}
		}
	}

	return union
}

// Difference returns a new slice containing unique elements of s1 not in s2.
// The order of elements is the same as s1.
func Difference[E comparable](s1, s2 []E) []E {
	u2 := toSet(s2)
	diff := make([]E, 0, len(s1))

	for _, e := range s1 {
		if _, found := u2[e]; !found {
			diff = append(diff, e)
			u2[e] = struct{}{} // removing successive duplicates
		}
	}
	return diff
}

// SymmetricDifference returns a new slice of unique elements present in either s1 or s2, but not both.
// The order of elements is the same as the one of append(s1, s2...).
func SymmetricDifference[E comparable](s1, s2 []E) []E {
	u2 := toSet(s2)
	seen := make(map[E]struct{}, len(s1)+len(u2))
	diff := make([]E, 0, len(s1)+len(u2))

	for _, e := range s1 {
		if _, found := seen[e]; !found {
			seen[e] = struct{}{} // removing duplicates

			if _, in2 := u2[e]; !in2 {
				diff = append(diff, e)
			}
		}
	}

	for _, e := range s2 {
		if _, found := seen[e]; !found {
			// this element is unique to s2 since it was not found in the iteration over s1,
			// so we mark it as seen and add it to the symmetric difference
			seen[e] = struct{}{}
			diff = append(diff, e)
		}
	}

	return diff
}

// Partition returns three unique and non-overlapping slices: elements only in s1, elements in both, and elements only in s2.
// The order of elements is preserved from their first appearance in s1, then s2.
func Partition[E comparable](s1, s2 []E) (d12, inter, d21 []E) {
	u2 := toSet(s2)
	seen := make(map[E]struct{}, len(s1)+len(u2))

	d12 = make([]E, 0, len(s1))
	inter = make([]E, 0)
	d21 = make([]E, 0, len(u2))

	for _, e := range s1 {
		if _, found := seen[e]; !found {
			seen[e] = struct{}{} // removing duplicates

			_, in2 := u2[e]
			if in2 {
				inter = append(inter, e)
			} else {
				d12 = append(d12, e)
			}
		}
	}

	for _, e := range s2 {
		if _, found := seen[e]; !found {
			// this element is unique to s2 since it was not found in the iteration over s1,
			// so we mark it as seen and add it to d21
			seen[e] = struct{}{}
			d21 = append(d21, e)
		}
	}

	return d12, inter, d21
}

func toSet[E comparable](s []E) map[E]struct{} {
	m := make(map[E]struct{}, len(s))
	for _, e := range s {
		m[e] = struct{}{}
	}
	return m
}
