package array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// The ForEach function doesn't return anything, so it's a bit tricky to test directly. However, you can test its side effects. For example, you can use it to modify a variable in the outer scope and then check that variable.

func TestForEach(t *testing.T) {
	type args struct {
		a []int
	}

	type test struct {
		name string
		args args
		want int
	}

	tests := []test{
		{
			name: "should loop over slice",
			args: args{a: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
			want: 45,
		},
		{
			name: "should loop over empty slice",
			args: args{a: []int{}},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sum := 0
			ForEach[int](tt.args.a, func(e, i int, s []int) {
				sum += e
			})
			assert.Equal(t, tt.want, sum) // The sum of the numbers 1 through 9 is 45
		})
	}
}
