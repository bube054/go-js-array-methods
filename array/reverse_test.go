package array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverse(t *testing.T) {
	type cust struct {
		a int
		b int
	}

	tests := []struct {
		name string
		arg  interface{} // input argument
		want interface{} // expected output
	}{
		{
			name: "ReverseStringSlice",
			arg:  []string{"a", "b", "c"},
			want: []string{"c", "b", "a"},
		},
		{
			name: "ReverseIntSlice",
			arg:  []int{1, 2, 3},
			want: []int{3, 2, 1},
		},
		{
			name: "ReverseBoolSlice",
			arg:  []bool{true, false},
			want: []bool{false, true},
		},
		{
			name: "ReverseCustomStructSlice",
			arg:  []cust{{a: 2, b: 1}, {a: 5, b: 6}},
			want: []cust{{a: 5, b: 6}, {a: 2, b: 1}},
		},
		{
			name: "EmptySlice",
			arg:  []int{},
			want: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Type assertion to specify the type explicitly
			switch v := tt.arg.(type) {
			case []string:
				got := Reverse(v)
				assert.Equal(t, tt.want, got, tt.name)
			case []int:
				got := Reverse(v)
				assert.Equal(t, tt.want, got)
			case []bool:
				got := Reverse(v)
				assert.Equal(t, tt.want, got)
			case []cust:
				got := Reverse(v)
				assert.Equal(t, tt.want, got)
			case []interface{}:
				got := Reverse(v)
				assert.Equal(t, tt.want, got)
			default:
				t.Errorf("unsupported type: %T", v)
			}
		})
	}
}
