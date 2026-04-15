package array_test

import (
	"fmt"

	array "github.com/bube054/go-js-array-methods/array"
)

// ----- Package-level function examples -----

func ExampleAt() {
	fruits := []string{"Banana", "Orange", "Apple", "Mango"}
	first, _ := array.At(fruits, 0)
	last, _ := array.At(fruits, -1)
	fmt.Println(first, last)
	// Output: Banana Mango
}

func ExampleConcat() {
	a := []int{1, 2, 3}
	b := []int{4, 5}
	c := []int{6}
	fmt.Println(array.Concat(a, b, c))
	// Output: [1 2 3 4 5 6]
}

func ExampleFilter() {
	nums := []int{1, 2, 3, 4, 5}
	even := array.Filter(nums, func(n, _ int, _ []int) bool { return n%2 == 0 })
	fmt.Println(even)
	// Output: [2 4]
}

func ExampleMap() {
	nums := []int{1, 2, 3}
	doubled := array.Map(nums, func(n, _ int, _ []int) int { return n * 2 })
	fmt.Println(doubled)
	// Output: [2 4 6]
}

func ExampleMap_typeChange() {
	nums := []int{1, 2, 3}
	labels := array.Map(nums, func(n, _ int, _ []int) string {
		return fmt.Sprintf("#%d", n)
	})
	fmt.Println(labels)
	// Output: [#1 #2 #3]
}

func ExampleReduce() {
	nums := []int{1, 2, 3, 4}
	sum, _ := array.Reduce(nums, func(acc any, n, _ int, _ []int) any {
		return acc.(int) + n
	}, 0)
	fmt.Println(sum)
	// Output: 10
}

func ExampleReduceStrict() {
	nums := []int{1, 2, 3, 4}
	initial := 0
	sum, _ := array.ReduceStrict(nums, func(acc, n, _ int, _ []int) int {
		return acc + n
	}, &initial)
	fmt.Println(sum)
	// Output: 10
}

func ExampleReverse() {
	fmt.Println(array.Reverse([]int{1, 2, 3}))
	// Output: [3 2 1]
}

func ExamplePush() {
	fmt.Println(array.Push([]int{1, 2}, 3, 4))
	// Output: [1 2 3 4]
}

func ExamplePop() {
	rest, popped := array.Pop([]int{1, 2, 3})
	fmt.Println(rest, *popped)
	// Output: [1 2] 3
}

func ExampleIncludes() {
	fmt.Println(array.Includes([]string{"a", "b", "c"}, "b", nil))
	// Output: true
}

func ExampleSlice() {
	out, _ := array.Slice([]int{10, 20, 30, 40, 50}, 1, 4)
	fmt.Println(out)
	// Output: [20 30 40]
}

func ExampleFlat() {
	nested := []any{1, []int{2, 3}, []int{4, 5}}
	flat, _ := array.Flat[int](nested)
	fmt.Println(flat)
	// Output: [1 2 3 4 5]
}

func ExampleJoin() {
	sep := " - "
	fmt.Println(array.Join([]int{1, 2, 3}, &sep))
	// Output: 1 - 2 - 3
}

// ----- Array[T] receiver-method examples -----

func ExampleArray() {
	arr := array.Array[int]{1, 2, 3, 4, 5}
	fmt.Println(arr.Reverse())
	// Output: [5 4 3 2 1]
}

func ExampleArray_Filter() {
	arr := array.Array[int]{1, 2, 3, 4, 5}
	even := arr.Filter(func(n, _ int, _ []int) bool { return n%2 == 0 })
	fmt.Println(even)
	// Output: [2 4]
}

// Demonstrates method chaining: Filter -> Push -> Reverse, all on Array[T].
func ExampleArray_chaining() {
	arr := array.Array[int]{1, 2, 3, 4, 5}
	result := arr.
		Filter(func(n, _ int, _ []int) bool { return n%2 == 0 }).
		Push(6).
		Reverse()
	fmt.Println(result)
	// Output: [6 4 2]
}

func ExampleArray_MapStrict() {
	arr := array.Array[int]{1, 2, 3}
	doubled := arr.MapStrict(func(n, _ int, _ []int) int { return n * 2 })
	fmt.Println(doubled)
	// Output: [2 4 6]
}

func ExampleArray_Join() {
	arr := array.Array[int]{1, 2, 3}
	fmt.Println(arr.Join(" | "))
	// Output: 1 | 2 | 3
}

func ExampleArray_ToString() {
	arr := array.Array[string]{"a", "b", "c"}
	fmt.Println(arr.ToString())
	// Output: a,b,c
}
