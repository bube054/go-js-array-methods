package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindLast(t *testing.T) {
	type args struct {
		a []int
		b Predicate[int]
	}

	type test struct {
		name   string
		args   args
		want   int
		hasErr bool
		errMsg string
	}

	tests := []test{
		{
			name: "Element is in list (single choice)",
			args: args{a: []int{10, 20, 30, 40}, b: func(e, i int, s []int) bool {
				if e == 40 {
					return true
				} else {
					return false
				}
			}},
			want: 40,
		},
		{
			name: "Element is in list (multiple choices)",
			args: args{a: []int{10, 20, 30}, b: func(e, i int, s []int) bool {
				if e < 20 {
					return true
				} else {
					return false
				}
			}},
			want: 10,
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
			want:   10,
			hasErr: true,
			errMsg: "could not find element in slice",
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
			want:   10,
			hasErr: true,
			errMsg: "slice is empty",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.hasErr {
				_, err := FindLast(tt.args.a, tt.args.b)
				assert.Equal(t, tt.errMsg, err.Error())
			} else {
				got, _ := FindLast(tt.args.a, tt.args.b)
				assert.Equal(t, tt.want, *got)
			}
		})
	}
}
