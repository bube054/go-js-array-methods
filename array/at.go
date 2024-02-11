// Package array provides functions for manipulating arrays. This comprehensive library provides an array of helper functions specifically designed to empower developers in working efficiently with array slices. It encompasses popular methods like map, filter, reduce, forEach, find, some, and many more, offering streamlined functionalities to enhance your golang coding experience. For the javascript methods https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array
package array

import (
	"errors"
	"fmt"
	"math"
)

// At retrieves the element at the specified index in a slice.
//
// The function takes a slice 's' of any comparable type 'T' and an integer 'i'
// representing the index of the element to retrieve. It returns a pointer to the
// element and an error. If the index is out of range, it returns a nil pointer
// and an error indicating the out-of-range condition.
//
// Parameters:
//     s []T:      The slice from which to retrieve the element.
//     i int:       The index of the element to retrieve.
//
// Returns:
//     *T:         A pointer to the element at the specified index.
//     error:      An error indicating any issues encountered, such as an out-of-range index.
//
// Example:
//     intSlice := []int{1, 2, 3, 4, 5}
//     index := 2
//     element, err := At(intSlice, index)
//     if err != nil {
//         fmt.Println("Error:", err)
//     } else {
//         fmt.Println("Element at index", index, ":", *element)
//     }
//     // Output: Element at index 2 : 3
//
//     stringSlice := []string{"apple", "banana", "orange"}
//     index = -1
//     element, err = At(stringSlice, index)
//     if err != nil {
//         fmt.Println("Error:", err)
//     } else {
//         fmt.Println("Element at index", index, ":", *element)
//     }
//     // Output: Element at index -1 : banana
func At[T comparable](s []T, i int) (*T, error) {
	if math.Abs(float64(i)) > float64(len(s)-1) {
		return nil, errors.New(fmt.Sprintf("%v is out of range", i))
	}

	index := i
	if index < 0 {
		index = len(s) + index
	}

	return &s[index], nil
}