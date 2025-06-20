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

// Pair represents a key-value pair, optimized for scenarios where sorting
// or k-element selection (TopK/MinK) is performed based solely on the Val field.
//
// The Key field serves as an identifier to retrieve additional and potentially
// large associated data from an external source after sorting/selection is complete.
//
// This design prioritizes performance by minimizing the data that needs to be
// moved and compared during sort and k-element selections.
// For example, adding just one extra field can increase the execution
// time of the [Pairs.TopK] method by up-to 2x.
type Pair[K comparable, V cmp.Ordered] struct {
	Key K
	Val V
}

type Pairs[K comparable, V cmp.Ordered] []Pair[K, V]

func (p Pairs[K, V]) Len() int { return len(p) }

// Min returns the position and value of the minimal pair.
// It panics if p is empty.
func (p Pairs[K, V]) Min() (int, V) {
	if len(p) < 1 {
		panic("slicex.Min: pairs is empty")
	}

	i, min := 0, p[0].Val
	for j, pair := range p {
		if pair.Val < min {
			i = j
			min = pair.Val
		}
	}
	return i, min
}

// Max returns the position and value of the maximal pair.
// It panics if p is empty.
func (p Pairs[K, V]) Max() (int, V) {
	if len(p) < 1 {
		panic("slicex.Max: pairs is empty")
	}

	i, max := 0, p[0].Val
	for j, pair := range p {
		if pair.Val > max {
			i = j
			max = pair.Val
		}
	}
	return i, max
}

// Unpack returns the slice of keys and vals that constitute pairs.
func (p Pairs[K, V]) Unpack() ([]K, []V) {
	keys := make([]K, len(p))
	vals := make([]V, len(p))

	for i, pair := range p {
		keys[i] = pair.Key
		vals[i] = pair.Val
	}
	return keys, vals
}

// Pack keys and vals into a [Pairs] structure. It panics if their lenghts are different.
func Pack[K comparable, V cmp.Ordered](keys []K, vals []V) Pairs[K, V] {
	if len(keys) != len(vals) {
		panic("slicex.Pack: keys and vals must have the same lenght")
	}

	p := make(Pairs[K, V], len(keys))
	for i := range keys {
		p[i] = Pair[K, V]{Key: keys[i], Val: vals[i]}
	}
	return p
}

// MinK returns the k smallest pairs by value, sorted in ascending order.
//
// The original pairs will be modified.
func (p Pairs[K, V]) MinK(k int) Pairs[K, V] {
	if k < 1 || len(p) == 0 {
		return nil
	}

	if k >= len(p) {
		sort.Slice(p, func(i, j int) bool { return p[i].Val < p[j].Val })
		return p
	}

	mins := p[:k]
	i, max := mins.Max()

	for _, e := range p[k:] {
		if e.Val < max {
			// swap out the biggest element with the new one
			mins[i] = e
			i, max = mins.Max()
		}
	}

	sort.Slice(mins, func(i, j int) bool { return p[i].Val < p[j].Val })
	return mins
}

// TopK returns the k biggest pairs by value, sorted in descending order.
//
// The original pairs will be modified.
func (p Pairs[K, V]) TopK(k int) Pairs[K, V] {
	if k < 1 || len(p) == 0 {
		return nil
	}

	if k >= len(p) {
		sort.Slice(p, func(i, j int) bool { return p[i].Val > p[j].Val })
		return p
	}

	top := p[:k]
	i, min := top.Min()

	for _, e := range p[k:] {
		if e.Val > min {
			// swap out the smallest element with the new one
			top[i] = e
			i, min = top.Min()
		}
	}

	sort.Slice(top, func(i, j int) bool { return p[i].Val > p[j].Val })
	return top
}
