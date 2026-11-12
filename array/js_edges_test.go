package array_test

import (
	"slices"
	"testing"

	"github.com/bube054/go-js-array-methods/v2/array"
)

// These cases all work in JavaScript, and each comment shows the JavaScript
// result. Splice and Fill used to return an error for them because a start
// position equal to the length (meaning "the end") was treated as out of range.

func TestSpliceAtTheEnd(t *testing.T) {
	zero := 0
	tests := []struct {
		name  string
		slice []int
		index int
		want  []int
	}{
		// [1, 2].splice(2, 0, 3) -> [1, 2, 3]
		{"insert at len", []int{1, 2}, 2, []int{1, 2, 3}},
		// [].splice(0, 0, 3) -> [3]
		{"insert into empty", []int{}, 0, []int{3}},
		// nil behaves like an empty slice
		{"insert into nil", nil, 0, []int{3}},
		// [1, 2].splice(-2, 0, 3) -> [3, 1, 2]
		{"insert at -len", []int{1, 2}, -2, []int{3, 1, 2}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := array.Splice(tt.slice, tt.index, &zero, 3)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !slices.Equal(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSpliceRemoveFromEmpty(t *testing.T) {
	// [].splice(0) -> []
	got, err := array.Splice([]int{}, 0, nil)
	if err != nil || len(got) != 0 {
		t.Errorf("got %v, %v; want [], nil", got, err)
	}
}

func TestSpliceStillRejectsIndexesPastTheEnd(t *testing.T) {
	for _, index := range []int{3, -4} {
		if _, err := array.Splice([]int{1, 2}, index, nil); err == nil {
			t.Errorf("index %d: expected an error", index)
		}
	}
}

func TestFillEmptyRanges(t *testing.T) {
	tests := []struct {
		name       string
		slice      []int
		start, end int
		want       []int
	}{
		// [].fill(9, 0, 0) -> []
		{"empty slice", []int{}, 0, 0, []int{}},
		// [1, 2].fill(9, 2, 2) -> [1, 2]
		{"start at len", []int{1, 2}, 2, 2, []int{1, 2}},
		// [1, 2].fill(9, -2, 2) -> [9, 9]
		{"whole slice from -len", []int{1, 2}, -2, 2, []int{9, 9}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := array.Fill(tt.slice, 9, tt.start, tt.end)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !slices.Equal(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFillStillRejectsBoundsPastTheEnd(t *testing.T) {
	if _, err := array.Fill([]int{1, 2}, 9, 3, 3); err == nil {
		t.Error("start 3 on a 2-element slice: expected an error")
	}
}
