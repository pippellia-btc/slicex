package slicex

import (
	"fmt"
	"math/rand/v2"
	"reflect"
	"testing"
)

func TestUnique(t *testing.T) {
	tests := []struct {
		slice    []int
		expected []int
	}{
		{slice: nil, expected: nil},
		{slice: []int{}, expected: []int{}},
		{slice: []int{1, 2, 0}, expected: []int{0, 1, 2}},
		{slice: []int{1, 2, 0, 3, 1, 0}, expected: []int{0, 1, 2, 3}},
	}

	for i, test := range tests {
		unique := Unique(test.slice)
		if !reflect.DeepEqual(unique, test.expected) {
			t.Fatalf("test %d: expected %v, got %v", i, test.expected, unique)
		}
	}
}

func TestUnion(t *testing.T) {
	tests := []struct {
		slices   [][]int
		expected []int
	}{
		{slices: nil, expected: nil},
		{slices: [][]int{{1, 2, 0}, {4, 5, 0}}, expected: []int{0, 1, 2, 4, 5}},
		{slices: [][]int{{1, 2, 0}, {3, 1, 0}, {6, 5, 7}, {-1, -2, 6}}, expected: []int{-2, -1, 0, 1, 2, 3, 5, 6, 7}},
	}

	for i, test := range tests {
		union := Union(test.slices...)
		if !reflect.DeepEqual(union, test.expected) {
			t.Fatalf("test %d: expected %v, got %v", i, test.expected, union)
		}
	}
}

func TestIntersection(t *testing.T) {
	tests := []struct {
		s1       []float64
		s2       []float64
		expected []float64
	}{
		{s1: nil, s2: nil, expected: nil},
		{s1: nil, s2: []float64{1}, expected: nil},
		{s1: []float64{3}, s2: []float64{}, expected: nil},
		{s1: []float64{1, 2, 3, 2}, s2: []float64{2, 2}, expected: []float64{2}},
		{s1: []float64{0, 0, 0, 0}, s2: []float64{1}, expected: nil},
		{s1: []float64{4, 2}, s2: []float64{2, 4}, expected: []float64{2, 4}},
	}

	for i, test := range tests {
		inter := Intersection(test.s1, test.s2)
		if !reflect.DeepEqual(inter, test.expected) {
			t.Errorf("test %d: expected %v, got %v", i, test.expected, inter)
		}
	}
}

func TestDifference(t *testing.T) {
	tests := []struct {
		s1       []float64
		s2       []float64
		expected []float64
	}{
		{s1: nil, s2: nil, expected: nil},
		{s1: nil, s2: []float64{1}, expected: nil},
		{s1: []float64{1, 2, 2}, s2: []float64{}, expected: []float64{1, 2}},
		{s1: []float64{1, 2, 3, 2}, s2: []float64{2}, expected: []float64{1, 3}},
		{s1: []float64{0, 0, 0, 0}, s2: []float64{1}, expected: []float64{0}},
		{s1: []float64{4, 2}, s2: []float64{2, 4}, expected: []float64{}},
	}

	for i, test := range tests {
		diff := Difference(test.s1, test.s2)
		if !reflect.DeepEqual(diff, test.expected) {
			t.Errorf("test %d: expected %v, got %v", i, test.expected, diff)
		}
	}
}

func TestDifferences(t *testing.T) {
	tests := []struct {
		s1     []float64
		s2     []float64
		u1, u2 []float64
	}{
		{s1: nil, s2: nil, u1: nil, u2: nil},
		{s1: nil, s2: []float64{1, 1}, u1: nil, u2: []float64{1}},
		{s1: []float64{1, 2, 2}, s2: []float64{}, u1: []float64{1, 2}, u2: nil},
		{s1: []float64{1, 2, 3, 2}, s2: []float64{2}, u1: []float64{1, 3}, u2: []float64{}},
		{s1: []float64{0, 0, 0, 0}, s2: []float64{1}, u1: []float64{0}, u2: []float64{1}},
		{s1: []float64{4, 2}, s2: []float64{2, 4}, u1: []float64{}, u2: []float64{}},
	}

	for i, test := range tests {
		u1, u2 := Differences(test.s1, test.s2)
		if !reflect.DeepEqual(u1, test.u1) || !reflect.DeepEqual(u2, test.u2) {
			t.Errorf("test %d: expected (%v,%v), got (%v,%v)", i, test.u1, test.u2, u1, u2)
		}
	}
}

func TestPartition(t *testing.T) {
	tests := []struct {
		s1            []float64
		s2            []float64
		u1, inter, u2 []float64
	}{
		{s1: nil, s2: nil, u1: nil, inter: nil, u2: nil},
		{s1: nil, s2: []float64{1}, u1: nil, inter: nil, u2: []float64{1}},
		{s1: []float64{3, 4}, s2: nil, u1: []float64{3, 4}, inter: nil, u2: nil},
		{s1: []float64{4, 3}, s2: []float64{1, 2, 1}, u1: []float64{3, 4}, inter: nil, u2: []float64{1, 2}},
		{s1: []float64{4, 3, 1}, s2: []float64{1, 2}, u1: []float64{3, 4}, inter: []float64{1}, u2: []float64{2}},
		{s1: []float64{4, 3, 1, 0, 0, 0}, s2: []float64{1, 2}, u1: []float64{0, 3, 4}, inter: []float64{1}, u2: []float64{2}},
	}

	for i, test := range tests {
		u1, inter, u2 := Partition(test.s1, test.s2)
		if !reflect.DeepEqual(u1, test.u1) || !reflect.DeepEqual(inter, test.inter) || !reflect.DeepEqual(u2, test.u2) {
			t.Errorf("test %d: expected (%v,%v,%v), got (%v,%v,%v)", i, test.u1, test.inter, test.u2, u1, inter, u2)
		}
	}
}

// ---------------------------------- benchmarks --------------------------------

type bench struct {
	size   int
	s1, s2 []int
}

var (
	sizes  = []int{1000, 10_000, 100_000, 1_000_000}
	benchs []bench
)

// init generates random slices for the benchmark.
func init() {
	for _, size := range sizes {
		bench := bench{
			size: size,
			s1:   RandomSlice(size, rand.IntN(size)),
			s2:   RandomSlice(size, rand.IntN(size)),
		}

		benchs = append(benchs, bench)
	}
}

func BenchmarkUnique(b *testing.B) {
	for _, bench := range benchs {
		b.Run(fmt.Sprintf("size=%d", bench.size), func(b *testing.B) {
			for range b.N {
				Unique(bench.s1)
			}
		})
	}
}

func BenchmarkUniqueMap(b *testing.B) {
	for _, bench := range benchs {
		b.Run(fmt.Sprintf("size=%d", bench.size), func(b *testing.B) {
			for range b.N {
				UniqueMap(bench.s1)
			}
		})
	}
}

func BenchmarkIntersection(b *testing.B) {
	for _, bench := range benchs {
		b.Run(fmt.Sprintf("size=%d", bench.size), func(b *testing.B) {
			for range b.N {
				Intersection(bench.s1, bench.s2)
			}
		})
	}
}

func BenchmarkIntersectionMap(b *testing.B) {
	for _, bench := range benchs {
		b.Run(fmt.Sprintf("size=%d", bench.size), func(b *testing.B) {
			for range b.N {
				IntersectionMap(bench.s1, bench.s2)
			}
		})
	}
}

func BenchmarkUnion(b *testing.B) {
	for _, bench := range benchs {
		b.Run(fmt.Sprintf("size=%d", bench.size), func(b *testing.B) {
			for range b.N {
				Union(bench.s1, bench.s2, bench.s1, bench.s2)
			}
		})
	}
}

func BenchmarkUnionMap(b *testing.B) {
	for _, bench := range benchs {
		b.Run(fmt.Sprintf("size=%d", bench.size), func(b *testing.B) {
			for range b.N {
				UnionMap(bench.s1, bench.s2, bench.s1, bench.s2)
			}
		})
	}
}
func BenchmarkDifference(b *testing.B) {
	for _, bench := range benchs {
		b.Run(fmt.Sprintf("size=%d", bench.size), func(b *testing.B) {
			for range b.N {
				Difference(bench.s1, bench.s2)
			}
		})
	}
}

func BenchmarkDifferenceMap(b *testing.B) {
	for _, bench := range benchs {
		b.Run(fmt.Sprintf("size=%d", bench.size), func(b *testing.B) {
			for range b.N {
				DifferenceMap(bench.s1, bench.s2)
			}
		})
	}
}

func BenchmarkDifferences(b *testing.B) {
	for _, bench := range benchs {
		b.Run(fmt.Sprintf("size=%d", bench.size), func(b *testing.B) {
			for range b.N {
				Differences(bench.s1, bench.s2)
			}
		})
	}
}

func BenchmarkDifferencesMap(b *testing.B) {
	for _, bench := range benchs {
		b.Run(fmt.Sprintf("size=%d", bench.size), func(b *testing.B) {
			for range b.N {
				DifferencesMap(bench.s1, bench.s2)
			}
		})
	}
}

func BenchmarkPartition(b *testing.B) {
	for _, bench := range benchs {
		b.Run(fmt.Sprintf("size=%d", bench.size), func(b *testing.B) {
			for range b.N {
				Partition(bench.s1, bench.s2)
			}
		})
	}
}

func BenchmarkPartitionMap(b *testing.B) {
	for _, bench := range benchs {
		b.Run(fmt.Sprintf("size=%d", bench.size), func(b *testing.B) {
			for range b.N {
				PartitionMap(bench.s1, bench.s2)
			}
		})
	}
}

// ----------------------------- map variants ---------------------------------

func toMap[E comparable](s []E) map[E]struct{} {
	m := make(map[E]struct{}, len(s))
	for _, e := range s {
		m[e] = struct{}{}
	}
	return m
}

func toSlice[E comparable](m map[E]struct{}) []E {
	s := make([]E, 0, len(m))
	for e := range m {
		s = append(s, e)
	}
	return s
}

func UniqueMap[E comparable](s []E) []E {
	unique := toMap(s)

	var write int
	for e := range unique {
		s[write] = e
		write++
	}
	return s[:write]
}

func IntersectionMap[E comparable](s1, s2 []E) []E {
	u2 := toMap(s2)

	var write int
	for _, e := range s1 {
		if _, exists := u2[e]; exists {
			s1[write] = e
			write++
			delete(u2, e) // remove successive duplicates
		}
	}
	return s1[:write]
}

func UnionMap[E comparable](inputs ...[]E) []E {
	if len(inputs) == 0 {
		return nil
	}

	u := inputs[0]
	for _, s := range inputs[1:] {
		u = append(u, s...)
	}
	return UniqueMap(u)

}

func DifferenceMap[E comparable](s1, s2 []E) []E {
	u2 := toMap(s2)

	var write int
	for _, e := range s1 {
		if _, exists := u2[e]; !exists {
			s1[write] = e
			write++
			u2[e] = struct{}{} // remove successive duplicates
		}
	}
	return s1[:write]
}

func DifferencesMap[E comparable](s1, s2 []E) ([]E, []E) {
	u1 := toMap(s1)
	u2 := toMap(s2)

	var write1 int
	for _, e := range s1 {
		if _, exists := u2[e]; !exists {
			s1[write1] = e
			write1++
			u2[e] = struct{}{} // remove successive duplicates
		}
	}

	var write2 int
	for _, e := range s2 {
		if _, exists := u1[e]; !exists {
			s2[write2] = e
			write2++
			u1[e] = struct{}{} // remove successive duplicates
		}
	}

	return s1[:write1], s2[:write2]
}

func SymmetricDifferenceMap[E comparable](s1, s2 []E) []E {
	s1, s2 = DifferencesMap(s1, s2)
	return append(s1, s2...)
}

func PartitionMap[E comparable](s1, s2 []E) ([]E, []E, []E) {
	u1 := toMap(s1)
	u2 := toMap(s2)
	inter := make(map[E]struct{})

	var write1 int
	for e := range u1 {
		_, exists := u2[e]

		if exists {
			inter[e] = struct{}{}
		} else {
			s1[write1] = e
			write1++
		}
	}

	var write2 int
	for e := range u2 {
		_, exists := u1[e]

		if exists {
			inter[e] = struct{}{}
		} else {
			s2[write2] = e
			write2++
		}
	}

	return s1[:write1], toSlice(inter), s2[:write2]
}

func RandomSlice(size, max int) []int {
	s := make([]int, size)
	for i := range size {
		s[i] = rand.IntN(max)
	}
	return s
}
