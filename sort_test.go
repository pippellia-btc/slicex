package slicex

import (
	"cmp"
	"fmt"
	"math/rand/v2"
	"reflect"
	"slices"
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

			s1 := RandomFloats(rand.IntN(size) + 1)
			s2 := make([]float64, len(s1))
			copy(s2, s1)

			mins := MinK(s1, k)
			expected := MinKNaive(s2, k)

			if !reflect.DeepEqual(mins, expected) {
				t.Errorf("len(s) = %d; k = %d", len(s1), k)
				t.Fatalf("expected mins %v, got %v", expected, mins)
			}
		}
	})
}

func TestMaxK(t *testing.T) {
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
			top := MaxK(test.s, test.k)
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

			s1 := RandomFloats(rand.IntN(size) + 1)
			s2 := make([]float64, len(s1))
			copy(s2, s1)

			top := MaxK(s1, k)
			expected := MaxKNaive(s2, k)

			if !reflect.DeepEqual(top, expected) {
				t.Errorf("len(s) = %d; k = %d", len(s1), k)
				t.Fatalf("expected top %v, got %v", expected, top)
			}
		}
	})
}

func TestPairsMinK(t *testing.T) {
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
			s := RandomFloats(rand.IntN(size) + 1)
			p1 := toPairs(s)
			p2 := toPairs(s)

			mins := p1.MinK(k)
			expected := p2.MinKNaive(k)

			if !reflect.DeepEqual(mins, expected) {
				t.Errorf("len(p) = %d; k = %d", len(p1), k)
				t.Fatalf("expected mins %v, got %v", expected, mins)
			}
		}
	})
}

func TestPairsMaxK(t *testing.T) {
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
			top := test.pairs.MaxK(test.k)
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
			s := RandomFloats(rand.IntN(size) + 1)
			p1 := toPairs(s)
			p2 := toPairs(s)

			maxs := p1.MaxK(k)
			expected := p2.MaxKNaive(k)

			if !reflect.DeepEqual(maxs, expected) {
				t.Errorf("len(p) = %d; k = %d", len(p1), k)
				t.Fatalf("expected maxs %v, got %v", expected, maxs)
			}
		}
	})
}

func toPairs[E cmp.Ordered](s []E) Pairs[int, E] {
	p := make(Pairs[int, E], len(s))
	for i, e := range s {
		p[i] = Pair[int, E]{Key: i, Val: e}
	}
	return p
}

func RandomFloats(size int) []float64 {
	s := make([]float64, size)
	for i := range size {
		s[i] = rand.Float64()
	}
	return s
}

// -------------------------------- benchmarks --------------------------------

var (
	SortSizes  = []int{1000, 10_000, 100_000, 1_000_000}
	SortBenchs [][]float64
)

// init generates random slices of ints for the benchmarks
func init() {
	for _, size := range SortSizes {
		SortBenchs = append(SortBenchs, RandomFloats(size))
	}
}

func BenchmarkSortDescending(b *testing.B) {
	for _, bench := range SortBenchs {
		b.Run(fmt.Sprintf("size=%d", len(bench)), func(b *testing.B) {
			for range b.N {
				c := make([]float64, len(bench))
				copy(c, bench)
				SortDescending(c)
			}
		})
	}
}

func BenchmarkSortFunc(b *testing.B) {
	for _, bench := range SortBenchs {
		b.Run(fmt.Sprintf("size=%d", len(bench)), func(b *testing.B) {
			for range b.N {
				c := make([]float64, len(bench))
				copy(c, bench)
				slices.SortFunc(c, func(a, b float64) int { return cmp.Compare(b, a) })
			}
		})
	}
}

func BenchmarkMinK(b *testing.B) {
	for _, bench := range SortBenchs {
		b.Run(fmt.Sprintf("min_10/%d", len(bench)), func(b *testing.B) {
			for range b.N {
				c := make([]float64, len(bench))
				copy(c, bench)
				MinK(c, 10)
			}
		})
	}
}

func BenchmarkMinKNaive(b *testing.B) {
	for _, bench := range SortBenchs {
		b.Run(fmt.Sprintf("min_10/%d", len(bench)), func(b *testing.B) {
			for range b.N {
				c := make([]float64, len(bench))
				copy(c, bench)
				MinKNaive(c, 10)
			}
		})
	}
}

func BenchmarkMaxK(b *testing.B) {
	for _, bench := range SortBenchs {
		b.Run(fmt.Sprintf("max_10/%d", len(bench)), func(b *testing.B) {
			for range b.N {
				c := make([]float64, len(bench))
				copy(c, bench)
				MaxK(c, 10)
			}
		})
	}
}

func BenchmarkMaxKNaive(b *testing.B) {
	for _, bench := range SortBenchs {
		b.Run(fmt.Sprintf("max_10/%d", len(bench)), func(b *testing.B) {
			for range b.N {
				c := make([]float64, len(bench))
				copy(c, bench)
				MaxKNaive(c, 10)
			}
		})
	}
}

func BenchmarkPairsMinK(b *testing.B) {
	for _, bench := range SortBenchs {
		b.Run(fmt.Sprintf("min_10/%d", len(bench)), func(b *testing.B) {
			for range b.N {
				p := toPairs(bench)
				p.MinK(10)
			}
		})
	}
}

func BenchmarkPairsMinKNaive(b *testing.B) {
	for _, bench := range SortBenchs {
		b.Run(fmt.Sprintf("min_10/%d", len(bench)), func(b *testing.B) {
			for range b.N {
				p := toPairs(bench)
				p.MinKNaive(10)
			}
		})
	}
}

func BenchmarkPairsMaxK(b *testing.B) {
	for _, bench := range SortBenchs {
		b.Run(fmt.Sprintf("max_10/%d", len(bench)), func(b *testing.B) {
			for range b.N {
				p := toPairs(bench)
				p.MaxK(10)
			}
		})
	}
}

func BenchmarkPairsMaxKNaive(b *testing.B) {
	for _, bench := range SortBenchs {
		b.Run(fmt.Sprintf("max_10/%d", len(bench)), func(b *testing.B) {
			for range b.N {
				p := toPairs(bench)
				p.MaxKNaive(10)
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

func MaxKNaive[E cmp.Ordered](s []E, k int) []E {
	if k < 1 || len(s) == 0 {
		return nil
	}

	slices.Sort(s)
	if k >= len(s) {
		slices.Reverse(s)
		return s
	}

	top := s[len(s)-k:]
	slices.Reverse(top)
	return top
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

func (p Pairs[K, V]) MaxKNaive(k int) Pairs[K, V] {
	if k < 1 || len(p) == 0 {
		return nil
	}

	slices.SortFunc(p, descendingPairs)
	if k >= len(p) {
		return p
	}
	return p[:k]
}

func ascendingPairs[K comparable, V cmp.Ordered](a, b Pair[K, V]) int {
	return cmp.Compare(a.Val, b.Val)
}

func descendingPairs[K comparable, V cmp.Ordered](a, b Pair[K, V]) int {
	return cmp.Compare(b.Val, a.Val)
}
