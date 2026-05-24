package array

type Array[T any] []T

type Entry[T any] struct {
	index   int
	element T
}

type Predicate[T any] func(element T, index int, slice []T) bool

type ForEachFunc[T any] func(element T, index int, slice []T)

type MapFunc[T any, V any] func(element T, index int, slice []T) V

type MapFuncStrict[T any] func(element T, index int, slice []T) T

type ReduceFunc[T any] func(accumulator any, currentValue T, index int, slice []T) any

type ReduceAsFunc[T any, R any] func(total R, currentValue T, currentIndex int, slice []T) R

type ReduceStrictFunc[T any] func(total T, currentValue T, currentIndex int, slice []T) T
