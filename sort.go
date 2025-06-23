package slicex

import (
	"cmp"

	"slices"
)

// SortAscending sorts the provided slice in ascending order.
func SortAscending[E cmp.Ordered](s []E) {
	slices.Sort(s)
}

// SortDescending sorts the provided slice in descending order.
func SortDescending[E cmp.Ordered](s []E) {
	slices.Sort(s)
	slices.Reverse(s)
}

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

// MaxK returns the k biggest elements in the slice, sorted in descending order.
//
// The original slice will be modified.
func MaxK[E cmp.Ordered](s []E, k int) []E {
	if k < 1 || len(s) == 0 {
		return nil
	}

	if k >= len(s) {
		slices.Sort(s)
		slices.Reverse(s)
		return s
	}

	maxs := s[:k]
	i, min := Min(maxs)

	for _, e := range s[k:] {
		if e > min {
			// swap out the smallest element with the new one
			maxs[i] = e
			i, min = Min(maxs)
		}
	}

	slices.Sort(maxs)
	slices.Reverse(maxs)
	return maxs
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
// or k-element selection (MaxK/MinK) is performed based solely on the Val field.
//
// The Key field serves as an identifier to retrieve additional and potentially
// large associated data from an external source after sorting/selection is complete.
//
// This design prioritizes performance by minimizing the data that needs to be
// moved and compared during sort and k-element selections.
// For example, adding just one extra field can increase the execution
// time of the [Pairs.MaxK] method by up-to 2x.
type Pair[K comparable, V cmp.Ordered] struct {
	Key K
	Val V
}

type Pairs[K comparable, V cmp.Ordered] []Pair[K, V]

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

// ToPairs converts the map into a slice of key-value [Pairs].
func ToPairs[K comparable, V cmp.Ordered](m map[K]V) Pairs[K, V] {
	pairs := make(Pairs[K, V], 0, len(m))
	for k, v := range m {
		pairs = append(pairs, Pair[K, V]{Key: k, Val: v})
	}
	return pairs
}

func (p Pairs[K, V]) Len() int { return len(p) }

// Keys returns the slice of keys of the pairs.
func (p Pairs[K, V]) Keys() []K {
	keys := make([]K, len(p))
	for i, pair := range p {
		keys[i] = pair.Key
	}
	return keys
}

// Vals returns the slice of values of the pairs.
func (p Pairs[K, V]) Vals() []V {
	vals := make([]V, len(p))
	for i, pair := range p {
		vals[i] = pair.Val
	}
	return vals
}

// Unpack returns the slice of keys and vals that constitute the pairs.
func (p Pairs[K, V]) Unpack() ([]K, []V) {
	keys := make([]K, len(p))
	vals := make([]V, len(p))

	for i, pair := range p {
		keys[i] = pair.Key
		vals[i] = pair.Val
	}
	return keys, vals
}

// ToMap converts the slice of key-value [Pairs] into a map.
func (p Pairs[K, V]) ToMap() map[K]V {
	m := make(map[K]V, len(p))
	for _, pair := range p {
		m[pair.Key] = pair.Val
	}
	return m
}

// Min returns the minimal pair and its position.
// It panics if p is empty.
func (p Pairs[K, V]) Min() (int, Pair[K, V]) {
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
	return i, p[i]
}

// Max returns the maximal pair and its position.
// It panics if p is empty.
func (p Pairs[K, V]) Max() (int, Pair[K, V]) {
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
	return i, p[i]
}

// SortAscending sorts the provided pairs in ascending order.
func (p Pairs[K, V]) SortAscending() {
	slices.SortFunc(p, func(p1, p2 Pair[K, V]) int { return cmp.Compare(p1.Val, p2.Val) })
}

// SortAscending sorts the provided pairs in ascending order.
func (p Pairs[K, V]) SortDescending() {
	slices.SortFunc(p, func(p1, p2 Pair[K, V]) int { return cmp.Compare(p2.Val, p1.Val) })
}

// MinK returns the k smallest pairs by value, sorted in ascending order.
//
// The original pairs will be modified.
func (p Pairs[K, V]) MinK(k int) Pairs[K, V] {
	if k < 1 || len(p) == 0 {
		return nil
	}

	if k >= len(p) {
		p.SortAscending()
		return p
	}

	mins := p[:k]
	i, max := mins.maxVal()

	for _, e := range p[k:] {
		if e.Val < max {
			// swap out the biggest element with the new one
			mins[i] = e
			i, max = mins.maxVal()
		}
	}

	mins.SortAscending()
	return mins
}

// MaxK returns the k biggest pairs by value, sorted in descending order.
//
// The original pairs will be modified.
func (p Pairs[K, V]) MaxK(k int) Pairs[K, V] {
	if k < 1 || len(p) == 0 {
		return nil
	}

	if k >= len(p) {
		p.SortDescending()
		return p
	}

	maxs := p[:k]
	i, min := maxs.minVal()

	for _, e := range p[k:] {
		if e.Val > min {
			// swap out the smallest element with the new one
			maxs[i] = e
			i, min = maxs.minVal()
		}
	}

	maxs.SortDescending()
	return maxs
}

func (p Pairs[K, V]) minVal() (int, V) {
	i, min := 0, p[0].Val
	for j, pair := range p {
		if pair.Val < min {
			i = j
			min = pair.Val
		}
	}
	return i, min
}

func (p Pairs[K, V]) maxVal() (int, V) {
	i, max := 0, p[0].Val
	for j, pair := range p {
		if pair.Val > max {
			i = j
			max = pair.Val
		}
	}
	return i, max
}
