package array

type Predicate[T comparable] func(e T, i int, s []T) bool

type MapFunc[T comparable] func(e T, i int, s []T) any

type ReduceFunc[T comparable] func(acc any, e T, i int, s []T) any 

type ReduceStrictFunc[T comparable, K comparable] func(acc K, e T, i int, s []T) K 