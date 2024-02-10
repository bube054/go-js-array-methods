package array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEvery(t *testing.T) {
	type args struct {
		a []int
		b Predicate[int]
	}

	type test struct {
		name string
		args args
		want bool
	}

	tests := []test{
		{
			name: "All pass the predicate, all less than 4",
			args: args{a: []int{1, 2, 3}, b: func(e, i int, s []int) bool {
				return e < 4
			}},
			want: true,
		},
		{
			name: "At least one passes the predicate, is even",
			args: args{a: []int{1, 2, 3}, b: func(e, i int, s []int) bool {
				if e % 2 == 0 {
					return true
				}else{
					return false
				}
			}},
			want: false,
		},
		{
			name: "None passes the predicate, is greater than 10",
			args: args{a: []int{1, 2, 3, 8}, b: func(e, i int, s []int) bool {
				if e > 10 {
					return true
				}else{
					return false
				}
			}},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Every(tt.args.a, tt.args.b)
			assert.Equal(t, tt.want, got)
		})
	}
}
