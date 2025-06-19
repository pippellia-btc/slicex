package slicex

import (
	"cmp"
	"fmt"
	"math/rand/v2"
	"reflect"
	"slices"
	"sort"
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
			// {s: nil, k: 1, expected: nil},
			// {s: []int{}, k: 1, expected: nil},
			// {s: []int{0}, k: 1, expected: []int{0}},
			// {s: []int{0, 3, 1}, k: -1, expected: nil},
			{s: []int{0, 3, 1}, k: 10, expected: []int{3, 1, 0}},
			// {s: []int{0, 3, 1, 5, 1, -1, 2, 99, 32, -11}, k: 3, expected: []int{99, 32, 5}},
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

	slices.SortFunc(s, func(a, b E) int {
		switch {
		case a < b:
			return 1
		case a > b:
			return -1
		default:
			return 0
		}
	})

	if k >= len(s) {
		return s
	}
	return s[:k]
}

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
		b.Run(fmt.Sprintf("size=%d", bench.size), func(b *testing.B) {
			for range b.N {
				sort.Slice(bench.s1, func(i, j int) bool {
					return bench.s1[i] > bench.s1[j]
				})
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
