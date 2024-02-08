package array

import (
	"errors"
	"fmt"
)

func Shift[T comparable](s []T, defaultValue T) ([]T, T, error) {
	length := len(s)

	if length == 0 {
		return s, defaultValue, errors.New(fmt.Sprintf("%v is empty", s))
	}

	firstItemIndex := 0
	firstItem := s[firstItemIndex]
	if length == 1 {
		empty := []T{}
		return empty, firstItem, nil
	} else {
		secondItemIndex := firstItemIndex + 1
		withoutFirstItem := s[secondItemIndex:]
		return withoutFirstItem, firstItem, nil
	}
}

func ShiftMut[T comparable](s *[]T, defaultValue T) (T, error) {
	length := len(*s)

	if length == 0 {
		return defaultValue, errors.New(fmt.Sprintf("%v is empty", s))
	}

	firstItemIndex := 0
	firstItem := (*s)[firstItemIndex]
	if length == 1 {
		*s = []T{}
		return firstItem, nil
	} else {
		secondItemIndex := firstItemIndex + 1
		*s = (*s)[secondItemIndex:]
		return firstItem, nil
	}
}
