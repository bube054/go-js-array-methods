package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIncludes(t *testing.T) {
	type cust struct {
		x int
		y string
		z bool
	}

	type args struct {
		a []cust
		b cust
	}

	type test struct {
		name string
		args args
		want bool
	}

	tests := []test{
		{
			name: "In slice",
			args: args{a: []cust{{x: 1, y: "1", z: true}, {x: 0, y: "0", z: false}}, b: cust{x: 1, y: "1", z: true}},
			want: true,
		},
		{
			name: "Not in slice",
			args: args{a: []cust{{x: 1, y: "1", z: true}, {x: 0, y: "0", z: false}}, b: cust{x: 1, y: "2", z: true}},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Includes(tt.args.a, tt.args.b)
			assert.Equal(t, tt.want, got)
		})
	}
}
