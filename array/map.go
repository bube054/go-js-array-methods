package array

// import (
// 	"fmt"
// 	"errors"
// )

func Map[T comparable](s []T, f MapFunc[T]) []any {
	newList := []any{}

	for index, value := range s {
		newItem := f(value, index, s)
		newList = append(newList, newItem)
	}

	return newList
}

func MapMut() {

}
