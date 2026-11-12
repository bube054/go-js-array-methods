package array_test

import (
	"slices"
	"testing"

	"github.com/bube054/go-js-array-methods/v2/array"
)

// These tests use the package from outside, like an application does, to check
// the README's promise that the original slice is never modified. Each case
// used to change the original: Push appended into spare capacity, and Pop, Shift
// and Slice returned views with room to append into.

func TestPushAfterPopKeepsOriginal(t *testing.T) {
	orig := []int{1, 2, 3}

	popped, _ := array.Pop(orig)
	array.Push(popped, 99)

	if !slices.Equal(orig, []int{1, 2, 3}) {
		t.Errorf("original changed to %v", orig)
	}
}

func TestPushesOnTheSameSliceAreIndependent(t *testing.T) {
	base := make([]int, 2, 10) // spare capacity, as append often leaves

	a := array.Push(base, 1)
	b := array.Push(base, 2)

	if !slices.Equal(a, []int{0, 0, 1}) || !slices.Equal(b, []int{0, 0, 2}) {
		t.Errorf("got a=%v b=%v, want [0 0 1] and [0 0 2]", a, b)
	}
}

func TestPushAfterShiftKeepsOriginal(t *testing.T) {
	orig := make([]int, 3, 4)
	copy(orig, []int{1, 2, 3})

	shifted, _ := array.Shift(orig)
	array.Push(shifted, 9)
	_ = append(shifted, 9)

	if got := orig[:4]; !slices.Equal(got, []int{1, 2, 3, 0}) {
		t.Errorf("original's spare capacity changed: %v", got)
	}
}

func TestAppendAfterSliceKeepsOriginal(t *testing.T) {
	orig := []int{1, 2, 3, 4}

	part, err := array.Slice(orig, 0, 2)
	if err != nil {
		t.Fatal(err)
	}
	array.Push(part, 99)
	_ = append(part, 99)

	if !slices.Equal(orig, []int{1, 2, 3, 4}) {
		t.Errorf("original changed to %v", orig)
	}
}

func TestArrayMethodsKeepOriginal(t *testing.T) {
	orig := array.Array[int]{1, 2, 3}

	popped, _ := orig.Pop()
	popped.Push(99)
	shifted, _ := orig.Shift()
	shifted.Push(98)
	part, _ := orig.Slice(0, 1)
	part.Push(97)

	if !slices.Equal(orig, array.Array[int]{1, 2, 3}) {
		t.Errorf("original changed to %v", orig)
	}
}

// Find and FindLast return a pointer into the slice on purpose, so a match can be
// updated in place. These tests pin that documented behavior.
func TestFindPointerUpdatesTheSlice(t *testing.T) {
	nums := []int{1, 2, 3, 2}

	first := array.Find(nums, func(n, _ int, _ []int) bool { return n == 2 })
	last := array.FindLast(nums, func(n, _ int, _ []int) bool { return n == 2 })
	*first = 20
	*last = 40

	if !slices.Equal(nums, []int{1, 20, 3, 40}) {
		t.Errorf("got %v, want [1 20 3 40]", nums)
	}
}

func TestDereferencedFindIsACopy(t *testing.T) {
	nums := []int{1, 2, 3}

	value := *array.Find(nums, func(n, _ int, _ []int) bool { return n == 2 })
	value = 99
	_ = value

	if !slices.Equal(nums, []int{1, 2, 3}) {
		t.Errorf("changing a dereferenced copy changed the slice: %v", nums)
	}
}

// Entries used to return entries with hidden fields, so code outside the package
// could not read the index or value of any entry.
func TestEntriesAreReadableOutsideThePackage(t *testing.T) {
	entries := array.Entries([]string{"a", "b"})

	want := []array.Entry[string]{{Index: 0, Value: "a"}, {Index: 1, Value: "b"}}
	if !slices.Equal(entries, want) {
		t.Errorf("got %v, want %v", entries, want)
	}
	if entries[1].Index != 1 || entries[1].Value != "b" {
		t.Errorf("got entry %+v, want {Index:1 Value:b}", entries[1])
	}
}
