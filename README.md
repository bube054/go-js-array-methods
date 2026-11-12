# go-js-array-methods

[![Go Reference](https://pkg.go.dev/badge/github.com/bube054/go-js-array-methods/v2.svg)](https://pkg.go.dev/github.com/bube054/go-js-array-methods/v2)
[![CI](https://github.com/bube054/go-js-array-methods/actions/workflows/tests.yml/badge.svg)](https://github.com/bube054/go-js-array-methods/actions/workflows/tests.yml)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](license.md)

A comprehensive toolkit for slice manipulation in Go, inspired by JavaScript's `Array.prototype` methods and built to stay close to hand-written Go loops in speed.

- ✅ **Familiar API** — `Filter`, `Map`, `Reduce`, `Push`, `Pop`, `Includes`, and 30+ more, named to match their JS counterparts.
- ✅ **Allocation-free where it counts** — typed operations such as `ReduceAs`, `ReduceStrict`, `Includes`, and `Slice` make no heap allocations of their own, and `Join` makes exactly one (the result string). [Run the benchmarks](#benchmarks) to see the numbers on your machine.
- ✅ **Go 1.23 Iterators** — Native support for lazy iteration (`iter.Seq`, `iter.Seq2`), allowing massive datasets to be processed with exactly `0 B/op` overhead.
- ✅ **100% test coverage** — plus tests that use the package from outside, the way your code does, to check that inputs are never changed.
- ✅ **Two styles** — use plain functions (`array.Filter(slice, fn)`) or the fluent `Array[T]` type with chainable methods.
- ✅ **Non-mutating** — no function changes the slice you pass in. A few results share memory with it for speed (`Slice`, `Pop`, `Shift`, and the pointer from `Find`/`FindLast`), and their docs say so.
- ✅ **Negative indexes** — `At`, `Slice`, `IndexOf`, etc. accept negative indexes and return clear errors instead of panicking.
- ✅ **Type-safe generics** — including a fully generic `Map[T, V]` that returns `[]V`, not `[]any`.

## Install

Requires **Go 1.23+** to support native iterators.

```bash
go get github.com/bube054/go-js-array-methods/v2@v2.2.0
```

> **v2.2.0 update:** Inputs are never changed (fixes `Push`, `Pop`, `Shift`, and `Slice` sharing memory), `Splice`/`Fill` accept JavaScript edge cases, `Entries` results are readable, and `Unshift` replaces `UnShift`. See [CHANGELOG.md](./CHANGELOG.md) for details.

## Quick start

### Functional style

```go
package main

import (
	"fmt"
	"github.com/bube054/go-js-array-methods/v2/array"
)

func main() {
	nums := []int{1, 2, 3, 4, 5}

	even := array.Filter(nums, func(n, _ int, _ []int) bool { return n%2 == 0 })
	doubled := array.Map(even, func(n, _ int, _ []int) int { return n * 2 })

	fmt.Println(doubled) // [4 8]
}
```

### Object-oriented style with `Array[T]`

```go
package main

import (
	"fmt"
	"github.com/bube054/go-js-array-methods/v2/array"
)

func main() {
	arr := array.Array[int]{1, 2, 3, 4, 5}

	result := arr.
		Filter(func(n, _ int, _ []int) bool { return n%2 == 0 }).
		Push(6).
		Reverse()

	fmt.Println(result) // [6 4 2]
}
```

Both styles are interchangeable — pick the one that fits your code.

## API overview

| #   | Method              | Status | Summary                                                                                                                              |
| --- | ------------------- | ------ | ------------------------------------------------------------------------------------------------------------------------------------ |
| 1   | `At`                | ✔      | Index into a slice (supports negative indexes).                                                                                      |
| 2   | `Concat`            | ✔      | Join two or more slices into a new slice.                                                                                            |
| 3   | `CopyWithin`        | ✔      | Copy a section of a slice to another position within the same slice.                                                                 |
| 4   | `Entries`           | ✔      | Return `[]Entry[T]`, where each entry has `.Index` and `.Value` fields.                                                              |
| 5   | `EntriesIterator`   | ✔      | Go 1.23 Iterator: Yields `(index, value)` pairs as `iter.Seq2`.                                                                      |
| 6   | `Every`             | ✔      | `true` if every element passes the predicate.                                                                                        |
| 7   | `Fill`              | ✔      | Fill a range with a static value.                                                                                                    |
| 8   | `Filter`            | ✔      | Keep elements that pass the predicate.                                                                                               |
| 9   | `Find`              | ✔      | Pointer to the first match inside the slice (writing through it updates the slice), or `nil`.                                                                                     |
| 10  | `FindIndex`         | ✔      | Index of the first match, or `-1`.                                                                                                   |
| 11  | `FindLast`          | ✔      | Pointer to the last match inside the slice (writing through it updates the slice), or `nil`.                                                                                      |
| 12  | `FindLastIndex`     | ✔      | Index of the last match, or `-1`.                                                                                                    |
| 13  | `Flat`              | ✔      | Flatten nested `[]any` to a typed `[]T` (configurable depth).                                                                        |
| 14  | `FlatMap`           | ❌     | Not implemented.                                                                                                                     |
| 15  | `ForEach`           | ✔      | Run a callback for each element.                                                                                                     |
| 16  | `Includes`          | ✔      | `true` if the slice contains the value.                                                                                              |
| 17  | `IndexOf`           | ✔      | First index of a value, or `-1`.                                                                                                     |
| 18  | `Join`              | ✔      | Stringify and join elements with a separator (one allocation, for the result string).                                                |
| 19  | `Keys`              | ✔      | Go 1.23 Iterator: Yields indexes as `iter.Seq[int]`.                                                                                 |
| 20  | `LastIndexOf`       | ✔      | Last index of a value, or `-1`.                                                                                                      |
| 21  | `Map`               | ✔      | Transform each element. Returns `[]V` (generic over output type).                                                                    |
| 22  | `MapStrict`         | ✔      | Transform each element to the same type `T`.                                                                                         |
| 23  | `Pop`               | ✔      | Return slice without last element + pointer to popped value.                                                                         |
| 24  | `Push`              | ✔      | Append elements; returns new slice.                                                                                                  |
| 25  | `Reduce`            | ✔      | Fold to a single value of any type (JS-fidelity dynamic accumulator).                                                                |
| 26  | `ReduceAs`          | ✔      | Reduce to a specific type without interface boxing. Pass an initial value if the slice may be empty.                                 |
| 27  | `ReduceStrict`      | ✔      | Fold to a single value of the same type as the elements.                                                                             |
| 28  | `ReduceRight`       | ✔      | Fold from right to left.                                                                                                             |
| 29  | `ReduceRightAs`     | ✔      | Right-fold to a specific type without interface boxing. Pass an initial value if the slice may be empty.                             |
| 30  | `ReduceRightStrict` | ✔      | Right-to-left fold with strict typing.                                                                                               |
| 31  | `Reverse`           | ✔      | Reverse a slice; returns new slice.                                                                                                  |
| 32  | `Shift`             | ✔      | Return slice without first element + pointer to shifted value.                                                                       |
| 33  | `Slice`             | ✔      | Sub-slice from `start` to `end` (exclusive), supports negative indexes.                                                              |
| 34  | `Some`              | ✔      | `true` if any element passes the predicate.                                                                                          |
| 35  | `Sort`              | ❌     | Use the standard library's `sort` / `slices` packages.                                                                               |
| 36  | `Splice`            | ✔      | Remove and/or insert elements at an index.                                                                                           |
| 37  | `ToString`          | ✔      | Comma-separated string representation.                                                                                               |
| 38  | `Unshift`           | ✔      | Prepend elements; returns new slice. (`UnShift` still works but is deprecated.)                                                      |
| 39  | `ValueOf`           | ✔      | Identity — returns the slice itself.                                                                                                 |
| 40  | `Values`            | ✔      | Go 1.23 Iterator: Yields values as `iter.Seq[T]`.                                                                                    |
| 41  | `With`              | ✔      | Return a copy with the element at `index` replaced.                                                                                  |

For full signatures, runnable examples, and per-method docs, see
[**pkg.go.dev**](https://pkg.go.dev/github.com/bube054/go-js-array-methods/v2/array).

## Design & Performance Philosophy

- **Typed Fast Paths:** When the types are fixed at compile time, the work compiles down to plain loops with no interface boxing. `ReduceAs`, `ReduceStrict`, `Includes`, and `Slice` make no heap allocations of their own (your callback still can, for example by building strings), and `Join` makes exactly one, for the result string.
- **Dynamic JS-Fidelity vs. Static Speed:** Methods like standard `Reduce` are designed for complete JavaScript parity (allowing an accumulator to dynamically change types), which incurs a runtime boxing tax: summing 10,000 ints with `Reduce` makes about 10,000 allocations, while `ReduceStrict` makes none. For pipelines without boxing, use `ReduceAs` or `ReduceStrict`.
- **Lazy Iterators (Go 1.23):** Using `.EntriesIterator()`, `.Keys()`, or `.Values()` returns native Go iterators. Processing a massive array eagerly allocates memory proportionally, while lazy iterators process the same data with exactly `0 Bytes` of heap overhead.
- **No Mutation:** No function modifies the slice you pass in. For speed, `Slice`, `Pop`, and `Shift` return views that share memory with it, and `Find`/`FindLast` return a pointer into it. Appending to a view is safe; assigning to its elements, or writing through the pointer, changes the original. Use `slices.Clone` when you need a fully separate copy.
- **Negative Indexes:** Methods gracefully accept negative indexes (`At(-1)` gets the last element) and return typed errors instead of panicking on out-of-bounds access.

## Examples

A few common patterns. See [`array/example_test.go`](./array/example_test.go)
for many more (also rendered on pkg.go.dev).

**Lazy Iteration (Go 1.23+)**
Process large datasets without building an intermediate slice, using native iterators:

```go
arr := array.Array[string]{"alice", "bob", "carol"}

for index, value := range arr.EntriesIterator() {
	fmt.Printf("Index: %d, Value: %s\n", index, value)
}
```

**Typed Reduction (`ReduceAs`)**
Transform a slice of integers into a completely different type (here a `string`). The accumulator is typed, so nothing is boxed into `any`:

```go
nums := []int{1, 2, 3}
str := array.ReduceAs(nums, func(acc string, curr, _ int, _ []int) string {
	return acc + strconv.Itoa(curr)
}, "")
// str == "123"
```

**Sum a slice with `ReduceStrict`:**

```go
nums := []int{1, 2, 3, 4}
initial := 0
sum, _ := array.ReduceStrict(nums, func(acc, n, _ int, _ []int) int {
	return acc + n
}, &initial)
// sum == 10
```

**Flatten a mixed nested slice:**

```go
nested := []any{1, []int{2, 3}, []int{4, 5}}
flat, _ := array.Flat[int](nested)
// flat == []int{1, 2, 3, 4, 5}
```

**Chain transforms with `Array[T]`:**

```go
arr := array.Array[string]{"alice", "bob", "carol"}
shouts := arr.MapStrict(func(s string, _ int, _ []string) string {
	return strings.ToUpper(s) + "!"
})
// shouts == [ALICE! BOB! CAROL!]
```

## Benchmarks

A representative benchmark suite lives in
[`array/benchmark_test.go`](./array/benchmark_test.go). Run it to see the speed
and allocations of each method on your machine:

```bash
go test -bench=. -benchmem ./array/
```

## Contributing

Issues and PRs welcome! See [CONTRIBUTING.md](./CONTRIBUTING.md) for the
dev loop, coding conventions, and what we look for in a PR.

## Changelog

See [CHANGELOG.md](./CHANGELOG.md).

## License

[MIT](./license.md)
