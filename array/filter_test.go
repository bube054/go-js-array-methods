package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {
	type args struct {
		a     []int
		b     Predicate[int]
	}

	type test struct {
		name   string
		args   args
		want   []int
	}

	tests := []test{
		{
			name: "Are even",
			args: args{a: []int{1, 2, 3, 4}, b: func(e, i int, s []int) bool {
				if e % 2 == 0 {
					return true
				}else{
					return false
				}
			}},
			want: []int{2, 4},
		},
		{
			name: "Are odd",
			args: args{a: []int{1, 2, 3, 4}, b: func(e, i int, s []int) bool {
				if e % 2 != 0 {
					return true
				}else{
					return false
				}
			}},
			want: []int{1, 3},
		},
		{
			name: "Is greater than 10",
			args: args{a: []int{1, 2, 3, 4}, b: func(e, i int, s []int) bool {
				if e > 10 {
					return true
				}else{
					return false
				}
			}},
			want: []int{},
		},
		{
			name: "Empty slice, and are odd",
			args: args{a: []int{}, b: func(e, i int, s []int) bool {
				if e % 2 != 0 {
					return true
				}else{
					return false
				}
			}},
			want: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
				got := Filter(tt.args.a, tt.args.b)
				assert.Equal(t, tt.want, got)
		})
	}
}
