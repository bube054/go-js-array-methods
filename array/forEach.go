package array

func ForEach[T comparable](slice []T, fn ForEachFunc[T]){
	for index, value := range slice {
		fn(value, index, slice)
	}
}