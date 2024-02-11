package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFlat(t *testing.T) {
	type args struct {
		a [][]int
	}

	type test struct {
		name string
		args args
		want []int
	}

	tests := []test{
		{
			name: "flatten slice of 2 slices",
			args: args{a: [][]int{{1,2,3}, {4,5,6}}},
			want: []int{1,2,3,4,5,6},
		},
		{
			name: "flatten slice of both slices empty",
			args: args{a: [][]int{{}, {}}},
			want: []int{},
		},
		{
			name: "flatten slice of first slice empty, then second full",
			args: args{a: [][]int{{}, {4,5,6}}},
			want: []int{4,5,6},
		},
		{
			name: "flatten slice of first slice full, then second empty",
			args: args{a: [][]int{{1,2,3}, {}}},
			want: []int{1,2,3},
		},
		{
			name: "flatten slice of 3 slices",
			args: args{a: [][]int{{1,2,3}, {4,5,6}, {7,8,9}}},
			want: []int{1,2,3,4,5,6,7,8,9},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Flat(tt.args.a)
			assert.Equal(t, tt.want, got)
		})
	}
}
