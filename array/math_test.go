package array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdd(t *testing.T) {
	type args struct {
		a int
		b int
	}

	type test struct {
		name string
		args args
		want int
	}

	tests := []test{
		{
			name: "should return added number",
			args: args{a: 10, b: 20},
			want: 30,
		},
		{
			name: "should return added number",
			args: args{a: 100, b: 20},
			want: 120,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Add(tt.args.a, tt.args.b)
			assert.Equal(t, tt.want, got)
		})
	}
}
