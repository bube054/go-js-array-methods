package array

// ForEach applies a given function to each element of a comparable slice.
//
// The function takes three arguments: the current element, its index, and the original slice.
// The function is applied in order for each element in the slice.
//
// This function does not return a value.
//
// Example:
//
// ForEach([]int{1, 2, 3, 4}, func(n int, i int, nums []int) {
//     fmt.Println(n, "is at index", i, "in", nums)
// })
//
// This will print:
// 1 is at index 0 in [1 2 3 4]
// 2 is at index 1 in [1 2 3 4]
// 3 is at index 2 in [1 2 3 4]
// 4 is at index 3 in [1 2 3 4]
//
// Parameters:
// s - The slice to iterate over.
// f - The function to apply to each element.
func ForEach[T comparable](s []T, f ForEachFunc[T]) {
	for index, value := range s {
		f(value, index, s)
	}
}
