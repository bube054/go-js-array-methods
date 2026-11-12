package array_test

import (
	"testing"

	"github.com/bube054/go-js-array-methods/v2/array"
)

// Array's Includes, IndexOf and LastIndexOf use == when the target can be
// compared, and reflect.DeepEqual when it cannot. The check used to look only at
// the target's kind, so values that == cannot handle took the == path and
// panicked: arrays of slices, and structs holding a slice in an interface field.

type box struct{ X any }

func TestSearchUncomparableValues(t *testing.T) {
	tests := []struct {
		name   string
		items  array.Array[any]
		target any
		want   int
	}{
		{"array of slices", array.Array[any]{"a", [1][]int{{1}}}, [1][]int{{1}}, 1},
		{"struct with a slice in an interface field", array.Array[any]{"a", box{X: []int{1}}}, box{X: []int{1}}, 1},
		{"array of slices, no match", array.Array[any]{[1][]int{{2}}}, [1][]int{{1}}, -1},
		{"comparable target among uncomparable items", array.Array[any]{[]int{1}, box{X: []int{1}}, "a"}, "a", 2},
		{"comparable struct still uses ==", array.Array[any]{box{X: 1}, box{X: 2}}, box{X: 2}, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.items.IndexOf(tt.target, nil); got != tt.want {
				t.Errorf("IndexOf = %d, want %d", got, tt.want)
			}
			if got := tt.items.LastIndexOf(tt.target, nil); got != tt.want {
				t.Errorf("LastIndexOf = %d, want %d", got, tt.want)
			}
			if got := tt.items.Includes(tt.target, nil); got != (tt.want != -1) {
				t.Errorf("Includes = %v, want %v", got, tt.want != -1)
			}
		})
	}
}

func TestSearchArrayOfSlicesType(t *testing.T) {
	// T itself is an array of slices, so == can never be used.
	items := array.Array[[2][]int]{{{1}, {2}}, {{3}, {4}}}

	if got := items.IndexOf([2][]int{{3}, {4}}, nil); got != 1 {
		t.Errorf("IndexOf = %d, want 1", got)
	}
}
