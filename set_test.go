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
		{slice: nil, expected: []int{}},
		{slice: []int{}, expected: []int{}},
		{slice: []int{1, 2, 0}, expected: []int{1, 2, 0}},
		{slice: []int{1, 2, 0, 3, 1, 0}, expected: []int{1, 2, 0, 3}},
	}

	for i, test := range tests {
		unique := Unique(test.slice)
		if !reflect.DeepEqual(unique, test.expected) {
			t.Errorf("test %d: expected %v, got %v", i, test.expected, unique)
		}
	}
}

func TestUnion(t *testing.T) {
	tests := []struct {
		slices   [][]int
		expected []int
	}{
		{slices: nil, expected: []int{}},
		{slices: [][]int{{1, 2, 0}, {4, 5, 0}}, expected: []int{1, 2, 0, 4, 5}},
		{slices: [][]int{{1, 2, 0}, {3, 1, 0}, {6, 5, 7}, {-1, -2, 6}}, expected: []int{1, 2, 0, 3, 6, 5, 7, -1, -2}},
	}

	for i, test := range tests {
		union := Union(test.slices...)
		if !reflect.DeepEqual(union, test.expected) {
			t.Errorf("test %d: expected %v, got %v", i, test.expected, union)
		}
	}
}

func TestIntersection(t *testing.T) {
	tests := []struct {
		s1       []float64
		s2       []float64
		expected []float64
	}{
		{s1: nil, s2: nil, expected: []float64{}},
		{s1: nil, s2: []float64{1}, expected: []float64{}},
		{s1: []float64{3}, s2: []float64{}, expected: []float64{}},
		{s1: []float64{1, 2, 3, 2}, s2: []float64{2, 2}, expected: []float64{2}},
		{s1: []float64{0, 0, 0, 0}, s2: []float64{1}, expected: []float64{}},
		{s1: []float64{4, 2, 5}, s2: []float64{2, 4}, expected: []float64{4, 2}},
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
		{s1: nil, s2: nil, expected: []float64{}},
		{s1: nil, s2: []float64{1}, expected: []float64{}},
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

func TestPartition(t *testing.T) {
	tests := []struct {
		s1            []float64
		s2            []float64
		u1, inter, u2 []float64
	}{
		{s1: nil, s2: nil, u1: []float64{}, inter: []float64{}, u2: []float64{}},
		{s1: nil, s2: []float64{1}, u1: []float64{}, inter: []float64{}, u2: []float64{1}},
		{s1: []float64{3, 4}, s2: nil, u1: []float64{3, 4}, inter: []float64{}, u2: []float64{}},
		{s1: []float64{4, 3}, s2: []float64{1, 2, 1}, u1: []float64{4, 3}, inter: []float64{}, u2: []float64{1, 2}},
		{s1: []float64{4, 3, 1}, s2: []float64{1, 2}, u1: []float64{4, 3}, inter: []float64{1}, u2: []float64{2}},
		{s1: []float64{4, 3, 1, 0, 0, 0}, s2: []float64{1, 2}, u1: []float64{4, 3, 0}, inter: []float64{1}, u2: []float64{2}},
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
	SetSizes  = []int{1000, 10_000, 100_000, 1_000_000}
	SetBenchs []bench
)

// init generates random slices of ints for the benchmarks
func init() {
	for _, size := range SetSizes {
		bench := bench{
			size: size,
			s1:   RandomInts(size, rand.IntN(size)),
			s2:   RandomInts(size, rand.IntN(size)),
		}

		SetBenchs = append(SetBenchs, bench)
	}
}

func RandomInts(size, max int) []int {
	s := make([]int, size)
	for i := range size {
		s[i] = rand.IntN(max)
	}
	return s
}

func BenchmarkUnique(b *testing.B) {
	for _, bench := range SetBenchs {
		b.Run(fmt.Sprintf("size=%d", bench.size), func(b *testing.B) {
			for range b.N {
				Unique(bench.s1)
			}
		})
	}
}

func BenchmarkIntersection(b *testing.B) {
	for _, bench := range SetBenchs {
		b.Run(fmt.Sprintf("size=%d", bench.size), func(b *testing.B) {
			for range b.N {
				Intersection(bench.s1, bench.s2, bench.s1[:bench.size/2], bench.s2[:bench.size/2])
			}
		})
	}
}

func BenchmarkUnion(b *testing.B) {
	for _, bench := range SetBenchs {
		b.Run(fmt.Sprintf("size=%d", bench.size), func(b *testing.B) {
			for range b.N {
				Union(bench.s1, bench.s2, bench.s1, bench.s2)
			}
		})
	}
}

func BenchmarkDifference(b *testing.B) {
	for _, bench := range SetBenchs {
		b.Run(fmt.Sprintf("size=%d", bench.size), func(b *testing.B) {
			for range b.N {
				Difference(bench.s1, bench.s2)
			}
		})
	}
}

func BenchmarkSymmetricDifference(b *testing.B) {
	for _, bench := range SetBenchs {
		b.Run(fmt.Sprintf("size=%d", bench.size), func(b *testing.B) {
			for range b.N {
				SymmetricDifference(bench.s1, bench.s2)
			}
		})
	}
}

func BenchmarkPartition(b *testing.B) {
	for _, bench := range SetBenchs {
		b.Run(fmt.Sprintf("size=%d", bench.size), func(b *testing.B) {
			for range b.N {
				Partition(bench.s1, bench.s2)
			}
		})
	}
}
