package array

import (
	"errors"
	"fmt"
)

func Pop[T comparable](s []T, defaultValue T) ([]T, T, error) {
	length := len(s)
 
	if length == 0 {
		return s, defaultValue, errors.New(fmt.Sprintf("%v is empty", s))
	}

	lastItemIndex := length - 1
	lastItem := s[lastItemIndex]
	withoutLastItem := s[:lastItemIndex]
	return withoutLastItem, lastItem, nil
}

func PopMut[T comparable] (s *[]T, defaultValue T) (T, error) {
	length := len(*s)

	if length == 0 {
		return defaultValue, errors.New(fmt.Sprintf("%v is empty", s))
	}

	lastItemIndex := length - 1
	lastItem := (*s)[lastItemIndex]
	*s = (*s)[:lastItemIndex]
	return lastItem, nil
}