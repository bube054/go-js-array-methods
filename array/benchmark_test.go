package array

import (
	"testing"
)

// makeInts returns a deterministic []int of length n: [0, 1, 2, ..., n-1].
// Used as a stable input across benchmarks so results are comparable.
func makeInts(n int) []int {
	out := make([]int, n)
	for i := 0; i < n; i++ {
		out[i] = i
	}
	return out
}

// makeNestedAny returns []any{0, 1, ..., n-1} where every other element is
// wrapped in a []int{x, x+1}, producing a depth-1 nested structure for Flat.
func makeNestedAny(n int) []any {
	out := make([]any, n)
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			out[i] = []int{i, i + 1}
		} else {
			out[i] = i
		}
	}
	return out
}

const benchN = 10_000

// ----- Package-level function benchmarks -----

func BenchmarkFilter(b *testing.B) {
	src := makeInts(benchN)
	pred := func(n, _ int, _ []int) bool { return n%2 == 0 }
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Filter(src, pred)
	}
}

func BenchmarkMap(b *testing.B) {
	src := makeInts(benchN)
	fn := func(n, _ int, _ []int) int { return n * 2 }
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Map(src, fn)
	}
}

func BenchmarkMapStrict(b *testing.B) {
	src := makeInts(benchN)
	fn := func(n, _ int, _ []int) int { return n * 2 }
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = MapStrict(src, fn)
	}
}

func BenchmarkReduce(b *testing.B) {
	src := makeInts(benchN)
	fn := func(acc any, n, _ int, _ []int) any { return acc.(int) + n }
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = Reduce(src, fn, 0)
	}
}

func BenchmarkReduceStrict(b *testing.B) {
	src := makeInts(benchN)
	initial := 0
	fn := func(acc, n, _ int, _ []int) int { return acc + n }
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = ReduceStrict(src, fn, &initial)
	}
}

func BenchmarkReverse(b *testing.B) {
	src := makeInts(benchN)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Reverse(src)
	}
}

func BenchmarkConcat(b *testing.B) {
	a := makeInts(benchN / 2)
	c := makeInts(benchN / 2)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Concat(a, c)
	}
}

func BenchmarkIncludes(b *testing.B) {
	src := makeInts(benchN)
	target := benchN - 1 // worst case: last element
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Includes(src, target, nil)
	}
}

func BenchmarkJoin(b *testing.B) {
	src := makeInts(1_000)
	sep := ","
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Join(src, &sep)
	}
}

// Flat uses reflect under the hood — worth tracking.
func BenchmarkFlat(b *testing.B) {
	src := makeNestedAny(1_000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = Flat[int](src)
	}
}

// ----- Array[T] method benchmarks (overhead vs. raw functions) -----

func BenchmarkArrayFilter(b *testing.B) {
	src := Array[int](makeInts(benchN))
	pred := func(n, _ int, _ []int) bool { return n%2 == 0 }
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = src.Filter(pred)
	}
}

func BenchmarkArrayMapStrict(b *testing.B) {
	src := Array[int](makeInts(benchN))
	fn := func(n, _ int, _ []int) int { return n * 2 }
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = src.MapStrict(fn)
	}
}

// Pipeline benchmark — what a real chained call costs end-to-end.
func BenchmarkArrayChain(b *testing.B) {
	src := Array[int](makeInts(benchN))
	pred := func(n, _ int, _ []int) bool { return n%2 == 0 }
	double := func(n, _ int, _ []int) int { return n * 2 }
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = src.Filter(pred).MapStrict(double).Reverse()
	}
}
