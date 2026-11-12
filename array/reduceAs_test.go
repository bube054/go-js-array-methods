package array

import "testing"

type reduceStats struct {
	total int
	seen  []int
}

func TestReduceAs_WithInitialValue(t *testing.T) {
	words := []string{"go", "js", "array"}
	lengthTotal := ReduceAs(words, func(acc int, word string, _ int, _ []string) int {
		return acc + len(word)
	}, 0)

	if lengthTotal != 9 {
		t.Fatalf("ReduceAs() with initial value = %d, want %d", lengthTotal, 9)
	}

	numbers := []int{1, 2, 3}
	stats := ReduceAs(numbers, func(acc reduceStats, n int, _ int, _ []int) reduceStats {
		acc.total += n
		acc.seen = append(acc.seen, n)
		return acc
	}, reduceStats{})

	if stats.total != 6 {
		t.Fatalf("ReduceAs() custom struct total = %d, want %d", stats.total, 6)
	}

	if len(stats.seen) != len(numbers) {
		t.Fatalf("ReduceAs() custom struct seen length = %d, want %d", len(stats.seen), len(numbers))
	}

	for i, got := range stats.seen {
		if got != numbers[i] {
			t.Fatalf("ReduceAs() custom struct seen[%d] = %d, want %d", i, got, numbers[i])
		}
	}
}

func TestReduceAs_WithoutInitialValue(t *testing.T) {
	numbers := []int{1, 2, 3}
	result := ReduceAs(numbers, func(acc int, n int, _ int, _ []int) int {
		return acc + n
	})

	if result != 6 {
		t.Fatalf("ReduceAs() without initial value = %d, want %d", result, 6)
	}
}

func TestReduceAs_EmptySliceWithInitialValue(t *testing.T) {
	result := ReduceAs([]string{}, func(acc int, value string, _ int, _ []string) int {
		return acc + len(value)
	}, 7)

	if result != 7 {
		t.Fatalf("ReduceAs() on empty slice with initial value = %d, want %d", result, 7)
	}
}

func TestReduceAs_Panics(t *testing.T) {
	t.Run("empty slice without initial value", func(t *testing.T) {
		defer func() {
			r := recover()
			if r == nil {
				t.Fatal("ReduceAs() did not panic on empty slice without initial value")
			}
			if got, ok := r.(string); !ok || got != "array: ReduceAs of empty slice with no initial value" {
				t.Fatalf("ReduceAs() panic = %v, want %q", r, "array: ReduceAs of empty slice with no initial value")
			}
		}()

		_ = ReduceAs([]string{}, func(acc int, value string, _ int, _ []string) int {
			return acc + len(value)
		})
	})

	t.Run("type mismatch without initial value", func(t *testing.T) {
		defer func() {
			r := recover()
			if r == nil {
				t.Fatal("ReduceAs() did not panic on type mismatch without initial value")
			}
			if got, ok := r.(string); !ok || got != "array: ReduceAs needs an initial value when the element type is not the accumulator type" {
				t.Fatalf("ReduceAs() panic = %v, want %q", r, "array: ReduceAs needs an initial value when the element type is not the accumulator type")
			}
		}()

		_ = ReduceAs([]string{"go", "js"}, func(acc int, value string, _ int, _ []string) int {
			return acc + len(value)
		})
	})
}
