package array

// The concat() function concatenates (joins) two or more slices.
// The concat() function returns a new array, containing the joined slices.
// The concat() function does not change the existing slices.
func Concat[T comparable] (slice1 []T, slice2 ...[]T) []T {
	var all []T = slice1

	for _, list := range slice2 {
		all = append(all, list...)
	}

	return all
}

// func ConcatMut [T comparable] (s1 *[]T, s2 ...[]T) {
// 	var all []T = *s1

// 	for _, list := range s2 {
// 		for _, val := range list {
// 			all = append(all, val)
// 		}
// 	}

// 	*s1 = all
// }