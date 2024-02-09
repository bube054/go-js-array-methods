package array

import (
	"errors"
	"fmt"
)

func Find[T comparable](s []T, f Predicate[T]) (*T, error) {
	for index, val := range s {
		if f(val, index, s) {
			return &val, nil
		}
	}

	return nil, errors.New(fmt.Sprintf("could not find from %v using %T", s, f))
}

	// d := []string{"ant", "bison", "camel", "duck", "elephant"}

	// var p array.FindFunc[string] = func (e string, ind int, arr []string) bool {
	// 	if e == "bison"{
	// 		return true
	// 	}else{
	// 		return false
	// 	}
	// }
	
	// r, err := array.Find[string](d, p)

	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	
	// fmt.Println(*r, err)