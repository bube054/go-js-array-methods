package array

func Concat [T comparable] (s1 []T, s2 ...[]T) []T {
	var all []T = s1

	for _, list := range s2 {
		for _, val := range list {
			all = append(all, val)
		}
	}

	return all
}

func ConcatMut [T comparable] (s1 *[]T, s2 ...[]T) {
	var all []T = *s1

	for _, list := range s2 {
		for _, val := range list {
			all = append(all, val)
		}
	}

	*s1 = all
}