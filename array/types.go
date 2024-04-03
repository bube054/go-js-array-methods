package array

type Predicate[T comparable] func(element T, index int, slice []T) bool

type ForEachFunc[T comparable] func(element T, index int, slice []T)

type MapFunc[T comparable] func(element T, index int, slice []T) any

type MapFuncStrict[T comparable] func(element T, index int, slice []T) T

type ReduceFunc[T comparable] func(total any, currentValue T, currentIndex int, slice []T) any

type ReduceStrictFunc[T comparable] func(total T, currentValue T, currentIndex int, slice []T) T
