package array

type Predicate[T comparable] func(element T, index int, slice []T) bool

type ForEachFunc[T comparable] func(element T, index int, slice []T)

type MapFunc[T comparable, V any] func(element T, index int, slice []T) V

type MapFuncStrict[T comparable] func(element T, index int, slice []T) T

type ReduceFunc[T comparable] func(total any, currentValue T, currentIndex int, slice []T) any

type ReduceStrictFunc[T comparable] func(total T, currentValue T, currentIndex int, slice []T) T

// A is an alias for Array, used in tests to verify type compatibility, also can be used as a shorthand.
type A[T comparable] = Array[T]
