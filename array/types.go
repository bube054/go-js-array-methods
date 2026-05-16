package array

type Array[T any] []T
type ArrayComp[T comparable] []T

type Entry[T any] struct {
	index   int
	element T
}

type Predicate[S ~[]T, T any] func(element T, index int, slice S) bool

type ForEachFunc[S ~[]T, T any] func(element T, index int, slice S)

type MapFunc[S ~[]T, T any, V any] func(element T, index int, slice S) V

type MapFuncStrict[S ~[]T, T any] func(element T, index int, slice S) T

type ReduceFunc[S ~[]T, T any] func(total any, currentValue T, currentIndex int, slice S) any

type ReduceStrictFunc[S ~[]T, T any] func(total T, currentValue T, currentIndex int, slice S) T
