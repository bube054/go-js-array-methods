package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindIndex(t *testing.T) {
	type args struct {
		a []int
		b Predicate[int]
	}

	type test struct {
		name string
		args args
		want int
	}

	tests := []test{
		{
			name: "Element is in list",
			args: args{a: []int{10, 20, 30}, b: func(e, i int, s []int) bool {
				if e < 20 {
					return true
				} else {
					return false
				}
			}},
			want: 0,
		},
		{
			name: "Element is not in list",
			args: args{a: []int{10, 20, 30}, b: func(e, i int, s []int) bool {
				if e > 100 {
					return true
				} else {
					return false
				}
			}},
			want: -1,
		},
		{
			name: "Empty list",
			args: args{a: []int{}, b: func(e, i int, s []int) bool {
				if e > 100 {
					return true
				} else {
					return false
				}
			}},
			want: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FindIndex(tt.args.a, tt.args.b)
			assert.Equal(t, tt.want, got)
		})
	}
}
