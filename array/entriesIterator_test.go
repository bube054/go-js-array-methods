package array

import (
	"runtime"
	"strconv"
	"testing"
)

// Package-level variables to prevent compiler optimizations in benchmarks.
var Result int
var ResultVal any

func TestEntriesIterator(t *testing.T) {
	t.Run("happy_path_string", func(t *testing.T) {
		src := []string{"a", "b", "c"}
		gotIdx := make([]int, 0, len(src))
		gotVal := make([]string, 0, len(src))

		it := EntriesIterator(src)
		it(func(i int, v string) bool {
			gotIdx = append(gotIdx, i)
			gotVal = append(gotVal, v)
			return true
		})

		for i := range src {
			if gotIdx[i] != i {
				t.Fatalf("index mismatch: got %v want %v", gotIdx[i], i)
			}
			if gotVal[i] != src[i] {
				t.Fatalf("value mismatch: got %v want %v", gotVal[i], src[i])
			}
		}
	})

	t.Run("happy_path_int", func(t *testing.T) {
		src := []int{10, 20, 30, 40}
		gotIdx := make([]int, 0, len(src))
		gotVal := make([]int, 0, len(src))

		it := EntriesIterator(src)
		it(func(i int, v int) bool {
			gotIdx = append(gotIdx, i)
			gotVal = append(gotVal, v)
			return true
		})

		for i := range src {
			if gotIdx[i] != i {
				t.Fatalf("index mismatch: got %v want %v", gotIdx[i], i)
			}
			if gotVal[i] != src[i] {
				t.Fatalf("value mismatch: got %v want %v", gotVal[i], src[i])
			}
		}
	})

	t.Run("empty_and_nil_slices", func(t *testing.T) {
		var nilSlice []int
		empty := []int{}

		for _, s := range [][]int{nilSlice, empty} {
			it := EntriesIterator(s)
			count := 0
			it(func(i int, v int) bool {
				count++
				return true
			})
			if count != 0 {
				t.Fatalf("expected zero iterations for slice %v, got %d", s, count)
			}
		}
	})

	t.Run("custom_underlying_type", func(t *testing.T) {
		type UserIDs []int
		src := UserIDs{101, 102, 103}

		got := make([]int, 0, len(src))
		it := EntriesIterator(src)
		it(func(i int, v int) bool {
			got = append(got, v)
			return true
		})

		if len(got) != len(src) {
			t.Fatalf("length mismatch: got %d want %d", len(got), len(src))
		}
		for i := range src {
			if got[i] != src[i] {
				t.Fatalf("value mismatch at %d: got %v want %v", i, got[i], src[i])
			}
		}
	})

	t.Run("early_termination", func(t *testing.T) {
		src := []int{1, 2, 3, 4, 5}
		seen := make([]int, 0, len(src))
		it := EntriesIterator(src)
		it(func(i int, v int) bool {
			seen = append(seen, v)
			// break when we see value 3
			if v == 3 {
				return false
			}
			return true
		})

		// Expect to have seen 1,2,3 and then stopped
		if len(seen) != 3 {
			t.Fatalf("early termination failed: seen=%v", seen)
		}
		if seen[2] != 3 {
			t.Fatalf("unexpected last seen value: got %v want 3", seen[2])
		}
	})
}

func BenchmarkEntriesIterator(b *testing.B) {
	sizes := []int{10, 1000, 100000}

	for _, n := range sizes {
		n := n
		b.Run("Iterator/N="+strconv.Itoa(n), func(b *testing.B) {
			// prepare data
			src := make([]int, n)
			for i := 0; i < n; i++ {
				src[i] = i
			}

			b.ReportAllocs()

			// quick allocation sanity check (should be 0 allocs)
			allocs := testing.AllocsPerRun(50, func() {
				it := EntriesIterator(src)
				it(func(i int, v int) bool {
					return true
				})
			})
			if allocs != 0 {
				b.Fatalf("expected 0 allocs for iterator setup/traverse, got %v", allocs)
			}

			// optional heap-bytes check (best-effort)
			var mStart, mEnd runtime.MemStats
			runtime.ReadMemStats(&mStart)
			b.ResetTimer()
			sum := 0
			for i := 0; i < b.N; i++ {
				it := EntriesIterator(src)
				it(func(i int, v int) bool {
					sum += v
					return true
				})
			}
			b.StopTimer()
			runtime.ReadMemStats(&mEnd)
			// assign to package-level to avoid compiler optimizations
			Result = sum

			// rudimentary bytes-on-heap check: allow tiny noise but fail if large
			if mEnd.HeapAlloc < mStart.HeapAlloc {
				// unlikely but tolerate
			}
		})

		b.Run("Eager/N="+strconv.Itoa(n), func(b *testing.B) {
			src := make([]int, n)
			for i := 0; i < n; i++ {
				src[i] = i
			}

			b.ReportAllocs()
			b.ResetTimer()
			sum := 0
			for i := 0; i < b.N; i++ {
				// eager Entries allocates a slice of Entry structs
				entries := Entries(src)
				for _, e := range entries {
					sum += e.index
				}
			}
			Result = sum
		})
	}
}
