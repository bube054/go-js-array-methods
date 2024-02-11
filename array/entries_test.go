package array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEntries(t *testing.T) {
	type args struct {
		a []int
	}

	type test struct {
		name string
		args args
		want []Entry[int]
	}

	tests := []test{
		{
			name: "should give slice of entries",
			args: args{a: []int{1, 2, 3}},
			want: []Entry[int]{{Index: 0, Value: 1}, {Index: 1, Value: 2}, {Index: 2, Value: 3}},
		},
		{
			name: "should give empty slice of entries",
			args: args{a: []int{}},
			want: []Entry[int]{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Entries(tt.args.a)
			assert.Equal(t, tt.want, got)
		})
	}
}
