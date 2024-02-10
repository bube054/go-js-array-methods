package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFill(t *testing.T) {
	type args struct {
		a     []int
		b     int
		start uint
		end   uint
	}

	type test struct {
		name   string
		args   args
		want   []int
		hasErr bool
		errMsg string
	}

	tests := []test{
		{
			name: "Normal fill all in range",
			args: args{a: []int{1, 2, 3, 4}, b: 0, start: 2, end: 3},
			want: []int{1, 2, 0, 4},
		},
		{
			name: "End position out of range",
			args: args{a: []int{1, 2, 3, 4}, b: 0, start: 2, end: 4}, // end must be greater than or equals to array length
			want: []int{1, 2, 0, 0},
		},
		{
			name: "fill all elements",
			args: args{a: []int{1, 2, 3, 4}, b: 0, start: 0, end: 3},
			want: []int{0, 0, 0, 0},
		},
		{
			name: "empty slice",
			args: args{a: []int{}, b: 0, start: 2, end: 4},
			want: []int{},
		},
		{
			name:   "Start position out of range",
			args:   args{a: []int{1, 2, 3, 4}, b: 0, start: 4, end: 4},
			want:   []int{},
			hasErr: true,
			errMsg: "start position is out of range",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.hasErr {
				_, err := Fill(tt.args.a, tt.args.b, tt.args.start, tt.args.end)
				assert.Equal(t, tt.errMsg, err.Error())
			} else {
				got, _ := Fill(tt.args.a, tt.args.b, tt.args.start, tt.args.end)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}
