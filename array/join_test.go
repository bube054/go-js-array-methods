package array

import (
	"testing"
)

func TestJoin(t *testing.T) {
	tests := []struct {
		name      string
		slice     []string
		separator *string
		expected  string
	}{
		{
			name:      "join with custom separator",
			slice:     []string{"Banana", "Orange", "Apple", "Mango"},
			separator: func() *string { s := " and "; return &s }(),
			expected:  "Banana and Orange and Apple and Mango",
		},
		{
			name:      "join with comma separator",
			slice:     []string{"Banana", "Orange", "Apple"},
			separator: func() *string { s := ","; return &s }(),
			expected:  "Banana,Orange,Apple",
		},
		{
			name:      "join with nil separator (default comma)",
			slice:     []string{"Banana", "Orange", "Apple"},
			separator: nil,
			expected:  "Banana,Orange,Apple",
		},
		{
			name:      "join empty array",
			slice:     []string{},
			separator: func() *string { s := ","; return &s }(),
			expected:  "",
		},
		{
			name:      "join single element",
			slice:     []string{"Apple"},
			separator: func() *string { s := ","; return &s }(),
			expected:  "Apple",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Join(tt.slice, tt.separator)
			if result != tt.expected {
				t.Errorf("Join() = %s, expected %s", result, tt.expected)
			}
		})
	}
}
