package slicex

import (
	"cmp"
	"fmt"
	"math/rand/v2"
	"reflect"
	"slices"
	"sort"
	"strconv"
	"testing"
)

func TestMinK(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		tests := []struct {
			s        []int
			k        int
			expected []int
		}{
			{s: nil, k: 1, expected: nil},
			{s: []int{}, k: 1, expected: nil},
			{s: []int{0}, k: 1, expected: []int{0}},
			{s: []int{0, 3, 1}, k: -1, expected: nil},
			{s: []int{0, 3, 1}, k: 10, expected: []int{0, 1, 3}},
			{s: []int{0, 3, 1, 5, 1, -1, 2, 99, 32, -11}, k: 3, expected: []int{-11, -1, 0}},
		}

		for i, test := range tests {
			mins := MinK(test.s, test.k)
			if !reflect.DeepEqual(mins, test.expected) {
				t.Fatalf("test %d: expected mins %v, got %v", i, test.expected, mins)
			}
		}
	})

	t.Run("fuzzy", func(t *testing.T) {
		const iter = 1000
		const size = 1000

		for range iter {
			k := rand.IntN(size)
			s := RandomSlice(rand.IntN(size)+1, rand.IntN(size)+1)

			mins := MinK(s, k)
			expected := MinKNaive(s, k)

			if !reflect.DeepEqual(mins, expected) {
				t.Errorf("len(s) = %d; k = %d", len(s), k)
				t.Fatalf("expected mins %v, got %v", expected, mins)
			}
		}
	})
}

func TestTopK(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		tests := []struct {
			s        []int
			k        int
			expected []int
		}{
			{s: nil, k: 1, expected: nil},
			{s: []int{}, k: 1, expected: nil},
			{s: []int{0}, k: 1, expected: []int{0}},
			{s: []int{0, 3, 1}, k: -1, expected: nil},
			{s: []int{0, 3, 1}, k: 10, expected: []int{3, 1, 0}},
			{s: []int{0, 3, 1, 5, 1, -1, 2, 99, 32, -11}, k: 3, expected: []int{99, 32, 5}},
		}

		for i, test := range tests {
			top := TopK(test.s, test.k)
			if !reflect.DeepEqual(top, test.expected) {
				t.Fatalf("test %d: expected top %v, got %v", i, test.expected, top)
			}
		}
	})

	t.Run("fuzzy", func(t *testing.T) {
		const iter = 1000
		const size = 1000

		for range iter {
			k := rand.IntN(size)
			s := RandomSlice(rand.IntN(size)+1, rand.IntN(size)+1)

			top := TopK(s, k)
			expected := TopKNaive(s, k)

			if !reflect.DeepEqual(top, expected) {
				t.Errorf("len(s) = %d; k = %d", len(s), k)
				t.Fatalf("expected top %v, got %v", expected, top)
			}
		}
	})
}

func TestMinKPairs(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		tests := []struct {
			pairs    Pairs[string, int]
			k        int
			expected Pairs[string, int]
		}{
			{pairs: nil, k: 1, expected: nil},
			{pairs: Pairs[string, int]{}, k: 1, expected: nil},
			{pairs: Pairs[string, int]{{Key: "0", Val: 0}}, k: 1, expected: Pairs[string, int]{{Key: "0", Val: 0}}},
			{pairs: Pairs[string, int]{{Key: "0", Val: 0}, {Key: "3", Val: 3}}, k: -1, expected: nil},
			{pairs: Pairs[string, int]{{Key: "0", Val: 0}, {Key: "3", Val: 3}, {Key: "1", Val: 1}}, k: 10, expected: Pairs[string, int]{{Key: "0", Val: 0}, {Key: "1", Val: 1}, {Key: "3", Val: 3}}},
			{pairs: Pairs[string, int]{{Key: "-1", Val: -1}, {Key: "3", Val: 3}, {Key: "1", Val: 1}}, k: 2, expected: Pairs[string, int]{{Key: "-1", Val: -1}, {Key: "1", Val: 1}}},
		}

		for i, test := range tests {
			mins := test.pairs.MinK(test.k)
			if !reflect.DeepEqual(mins, test.expected) {
				t.Fatalf("test %d: expected mins %v, got %v", i, test.expected, mins)
			}
		}
	})

	t.Run("fuzzy", func(t *testing.T) {
		const iter = 1000
		const size = 1000

		for range iter {
			k := rand.IntN(size)
			s := RandomSlice(rand.IntN(size)+1, rand.IntN(size)+1)
			p := toPairs(s)

			mins := p.MinK(k)
			expected := p.MinKNaive(k)

			if !reflect.DeepEqual(mins, expected) {
				t.Errorf("len(p) = %d; k = %d", len(p), k)
				t.Fatalf("expected mins %v, got %v", expected, mins)
			}
		}
	})
}

func TestTopKPairs(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		tests := []struct {
			pairs    Pairs[string, int]
			k        int
			expected Pairs[string, int]
		}{
			{pairs: nil, k: 1, expected: nil},
			{pairs: Pairs[string, int]{}, k: 1, expected: nil},
			{pairs: Pairs[string, int]{{Key: "0", Val: 0}}, k: 1, expected: Pairs[string, int]{{Key: "0", Val: 0}}},
			{pairs: Pairs[string, int]{{Key: "0", Val: 0}, {Key: "3", Val: 3}}, k: -1, expected: nil},
			{pairs: Pairs[string, int]{{Key: "0", Val: 0}, {Key: "3", Val: 3}, {Key: "1", Val: 1}}, k: 10, expected: Pairs[string, int]{{Key: "3", Val: 3}, {Key: "1", Val: 1}, {Key: "0", Val: 0}}},
			{pairs: Pairs[string, int]{{Key: "-1", Val: -1}, {Key: "3", Val: 3}, {Key: "1", Val: 1}}, k: 2, expected: Pairs[string, int]{{Key: "3", Val: 3}, {Key: "1", Val: 1}}},
		}

		for i, test := range tests {
			top := test.pairs.TopK(test.k)
			if !reflect.DeepEqual(top, test.expected) {
				t.Fatalf("test %d: expected top %v, got %v", i, test.expected, top)
			}
		}
	})

	t.Run("fuzzy", func(t *testing.T) {
		const iter = 1000
		const size = 1000

		for range iter {
			k := rand.IntN(size)
			s := RandomSlice(rand.IntN(size)+1, rand.IntN(size)+1)
			p := toPairs(s)

			top := p.TopK(k)
			expected := p.TopKNaive(k)

			if !reflect.DeepEqual(top, expected) {
				t.Errorf("len(p) = %d; k = %d", len(p), k)
				t.Fatalf("expected mins %v, got %v", expected, top)
			}
		}
	})
}

func toPairs(s []int) Pairs[string, int] {
	p := make(Pairs[string, int], len(s))
	for i, e := range s {
		p[i] = Pair[string, int]{Key: strconv.Itoa(e), Val: e}
	}
	return p
}

// -------------------------------- benchmarks --------------------------------

func BenchmarkSort(b *testing.B) {
	for _, bench := range benchs {
		b.Run(fmt.Sprintf("size=%d", bench.size), func(b *testing.B) {
			for range b.N {
				slices.Sort(bench.s1)
			}
		})
	}
}

func BenchmarkSortAndReverse(b *testing.B) {
	for _, bench := range benchs {
		b.Run(fmt.Sprintf("size=%d", bench.size), func(b *testing.B) {
			for range b.N {
				slices.Sort(bench.s1)
				slices.Reverse(bench.s1)
			}
		})
	}
}

func BenchmarkSortFunc(b *testing.B) {
	for _, bench := range benchs {
		b.Run(fmt.Sprintf("size=%d", bench.size), func(b *testing.B) {
			for range b.N {
				slices.SortFunc(bench.s1, func(a, b int) int { return a - b })
			}
		})
	}
}

func BenchmarkSortSlice(b *testing.B) {
	for _, bench := range benchs {
		s := bench.s1
		b.Run(fmt.Sprintf("size=%d", bench.size), func(b *testing.B) {
			for range b.N {
				sort.Slice(s, func(i, j int) bool { return s[i] > s[j] })
			}
		})
	}
}

func BenchmarkMinK(b *testing.B) {
	for _, bench := range benchs {
		b.Run(fmt.Sprintf("min_10/%d", bench.size), func(b *testing.B) {
			for range b.N {
				MinK(bench.s1, 10)
			}
		})
	}
}

func BenchmarkMinKNaive(b *testing.B) {
	for _, bench := range benchs {
		b.Run(fmt.Sprintf("min_10/%d", bench.size), func(b *testing.B) {
			for range b.N {
				MinKNaive(bench.s1, 10)
			}
		})
	}
}

func BenchmarkTopK(b *testing.B) {
	for _, bench := range benchs {
		b.Run(fmt.Sprintf("top_10/%d", bench.size), func(b *testing.B) {
			for range b.N {
				TopK(bench.s1, 10)
			}
		})
	}
}

func BenchmarkTopKNaive(b *testing.B) {
	for _, bench := range benchs {
		b.Run(fmt.Sprintf("top_10/%d", bench.size), func(b *testing.B) {
			for range b.N {
				TopKNaive(bench.s1, 10)
			}
		})
	}
}

func BenchmarkMinKPairs(b *testing.B) {
	for _, bench := range benchs {
		p := toPairs(bench.s1)
		b.Run(fmt.Sprintf("min_10/%d", bench.size), func(b *testing.B) {
			for range b.N {
				p.MinK(10)
			}
		})
	}
}

func BenchmarkMinKPairsNaive(b *testing.B) {
	for _, bench := range benchs {
		p := toPairs(bench.s1)
		b.Run(fmt.Sprintf("min_10/%d", bench.size), func(b *testing.B) {
			for range b.N {
				p.MinKNaive(10)
			}
		})
	}
}

func BenchmarkTopKPairs(b *testing.B) {
	for _, bench := range benchs {
		p := toPairs(bench.s1)
		b.Run(fmt.Sprintf("top_10/%d", bench.size), func(b *testing.B) {
			for range b.N {
				p.TopK(10)
			}
		})
	}
}

func BenchmarkTopKPairsNaive(b *testing.B) {
	for _, bench := range benchs {
		p := toPairs(bench.s1)
		b.Run(fmt.Sprintf("top_10/%d", bench.size), func(b *testing.B) {
			for range b.N {
				p.TopKNaive(10)
			}
		})
	}
}

// ---------------------------- naive variants --------------------------------

func MinKNaive[E cmp.Ordered](s []E, k int) []E {
	if k < 1 || len(s) == 0 {
		return nil
	}

	slices.Sort(s)
	if k >= len(s) {
		return s
	}
	return s[:k]
}

func TopKNaive[E cmp.Ordered](s []E, k int) []E {
	if k < 1 || len(s) == 0 {
		return nil
	}

	slices.SortFunc(s, descending)
	if k >= len(s) {
		return s
	}
	return s[:k]
}

func (p Pairs[K, V]) MinKNaive(k int) Pairs[K, V] {
	if k < 1 || len(p) == 0 {
		return nil
	}

	slices.SortFunc(p, ascendingPairs)
	if k >= len(p) {
		return p
	}
	return p[:k]
}

func (p Pairs[K, V]) TopKNaive(k int) Pairs[K, V] {
	if k < 1 || len(p) == 0 {
		return nil
	}

	slices.SortFunc(p, descendingPairs)
	if k >= len(p) {
		return p
	}
	return p[:k]
}

func descending[E cmp.Ordered](a, b E) int { return cmp.Compare(b, a) }

func ascendingPairs[K comparable, V cmp.Ordered](a, b Pair[K, V]) int {
	return cmp.Compare(a.Val, b.Val)
}

func descendingPairs[K comparable, V cmp.Ordered](a, b Pair[K, V]) int {
	return cmp.Compare(b.Val, a.Val)
}
