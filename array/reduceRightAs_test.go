package array

import (
	"strconv"
	"testing"
)

func TestReduceRightAs_OrderOfExecution(t *testing.T) {
	result := ReduceRightAs([]int{1, 2, 3}, func(acc string, n int, _ int, _ []int) string {
		return acc + strconv.Itoa(n)
	}, "")

	if result != "321" {
		t.Fatalf("ReduceRightAs() = %q, want %q", result, "321")
	}
}

func TestReduceRightAs_EmptySliceWithInitialValue(t *testing.T) {
	result := ReduceRightAs([]string{}, func(acc int, value string, _ int, _ []string) int {
		return acc + len(value)
	}, 7)

	if result != 7 {
		t.Fatalf("ReduceRightAs() on empty slice with initial value = %d, want %d", result, 7)
	}
}

func TestReduceRightAs_Panics(t *testing.T) {
	t.Run("empty slice without initial value", func(t *testing.T) {
		defer func() {
			r := recover()
			if r == nil {
				t.Fatal("ReduceRightAs() did not panic on empty slice without initial value")
			}
			if got, ok := r.(string); !ok || got != "array: ReduceRightAs of empty slice with no initial value" {
				t.Fatalf("ReduceRightAs() panic = %v, want %q", r, "array: ReduceRightAs of empty slice with no initial value")
			}
		}()

		_ = ReduceRightAs([]string{}, func(acc int, value string, _ int, _ []string) int {
			return acc + len(value)
		})
	})

	t.Run("type mismatch without initial value", func(t *testing.T) {
		defer func() {
			r := recover()
			if r == nil {
				t.Fatal("ReduceRightAs() did not panic on type mismatch without initial value")
			}
			if got, ok := r.(string); !ok || got != "array: ReduceRightAs needs an initial value when the element type is not the accumulator type" {
				t.Fatalf("ReduceRightAs() panic = %v, want %q", r, "array: ReduceRightAs needs an initial value when the element type is not the accumulator type")
			}
		}()

		_ = ReduceRightAs([]string{"go", "js"}, func(acc int, value string, _ int, _ []string) int {
			return acc + len(value)
		})
	})
}
