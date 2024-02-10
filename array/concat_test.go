package array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConcat(t *testing.T) {
	type args struct {
		a []int
		b []int
		c []int
	}

	type test struct {
		name string
		args args
		want []int
	}

	tests := []test{
		{
			name: "should combine 2 slices",
			args: args{a: []int{1, 2, 3}, b: []int{4, 5, 6}},
			want: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name: "should combine more than 2 slices",
			args: args{a: []int{1, 2, 3}, b: []int{4, 5, 6}, c: []int{7, 8, 9}},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name: "should return empty slices",
			args: args{a: []int{}, b: []int{}},
			want: []int{},
		},
		{
			name: "should combine more than 2 empty slices",
			args: args{a: []int{}, b: []int{}, c: []int{}},
			want: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.name == "should combine more than 2 slices" {
				got := Concat(tt.args.a, tt.args.b, tt.args.c)
				assert.Equal(t, tt.want, got)
			} else {
				got := Concat(tt.args.a, tt.args.b)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}
