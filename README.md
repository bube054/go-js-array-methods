# go-js-array-methods

[![Go Reference](https://pkg.go.dev/badge/github.com/bube054/go-js-array-methods/v2.svg)](https://pkg.go.dev/github.com/bube054/go-js-array-methods/v2)
[![Go Report Card](https://goreportcard.com/badge/github.com/bube054/go-js-array-methods/v2)](https://goreportcard.com/report/github.com/bube054/go-js-array-methods/v2)
[![CI](https://github.com/bube054/go-js-array-methods/actions/workflows/tests.yml/badge.svg)](https://github.com/bube054/go-js-array-methods/actions/workflows/tests.yml)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](license.md)

A comprehensive, generics-powered toolkit for slice manipulation in Go,
inspired by JavaScript's `Array.prototype` methods.

- ✅ **Familiar API** — `Filter`, `Map`, `Reduce`, `Push`, `Pop`, `Includes`,
  and 30+ more, named to match their JS counterparts.
- ✅ **Two styles** — use plain functions (`array.Filter(slice, fn)`) or the
  fluent `Array[T]` type with chainable methods.
- ✅ **Immutable by default** — every operation returns a new slice; the
  original is never mutated.
- ✅ **Negative indexes** — `At`, `Slice`, `IndexOf`, etc. accept negative
  indexes and return clear errors instead of panicking.
- ✅ **Type-safe generics** — including a fully generic `Map[T, V]` that
  returns `[]V`, not `[]any`.

## Install

```bash
go get github.com/bube054/go-js-array-methods/v2
```

> **v1 → v2 migration:** the `Map` signature changed from
> `Map[T](slice, fn) []any` to `Map[T, V](slice, fn) []V`. See
> [CHANGELOG.md](./CHANGELOG.md) for details.

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

| #   | Method            | Status | Summary                                                                 |
| --- | ----------------- | ------ | ----------------------------------------------------------------------- |
| 1   | `At`              | ✔      | Index into a slice (supports negative indexes).                         |
| 2   | `Concat`          | ✔      | Join two or more slices into a new slice.                               |
| 3   | `CopyWithin`      | ✔      | Copy a section of a slice to another position within the same slice.    |
| 4   | `Entries`         | ✔      | Return `[]Entry{Index, Value}` pairs.                                   |
| 5   | `Every`           | ✔      | `true` if every element passes the predicate.                           |
| 6   | `Fill`            | ✔      | Fill a range with a static value.                                       |
| 7   | `Filter`          | ✔      | Keep elements that pass the predicate.                                  |
| 8   | `Find`            | ✔      | Pointer to the first matching element, or `nil`.                        |
| 9   | `FindIndex`       | ✔      | Index of the first match, or `-1`.                                      |
| 10  | `FindLast`        | ✔      | Pointer to the last matching element, or `nil`.                         |
| 11  | `FindLastIndex`   | ✔      | Index of the last match, or `-1`.                                       |
| 12  | `Flat`            | ✔      | Flatten nested `[]any` to a typed `[]T` (configurable depth).           |
| 13  | `FlatMap`         | ❌     | Not implemented.                                                        |
| 14  | `ForEach`         | ✔      | Run a callback for each element.                                        |
| 15  | `Includes`        | ✔      | `true` if the slice contains the value.                                 |
| 16  | `IndexOf`         | ✔      | First index of a value, or `-1`.                                        |
| 17  | `Join`            | ✔      | Stringify and join elements with a separator (default `,`).             |
| 18  | `Keys`            | ❌     | Use `for i := range slice` instead.                                     |
| 19  | `LastIndexOf`     | ✔      | Last index of a value, or `-1`.                                         |
| 20  | `Map`             | ✔      | Transform each element. Returns `[]V` (generic over output type).       |
| 21  | `MapStrict`       | ✔      | Transform each element to the same type `T`.                            |
| 22  | `Pop`             | ✔      | Return slice without last element + pointer to popped value.            |
| 23  | `Push`            | ✔      | Append elements; returns new slice.                                     |
| 24  | `Reduce`          | ✔      | Fold to a single value of any type.                                     |
| 25  | `ReduceStrict`    | ✔      | Fold to a single value of the same type as the elements.                |
| 26  | `ReduceRight`     | ✔      | Fold from right to left.                                                |
| 27  | `ReduceRightStrict` | ✔    | Right-to-left fold with strict typing.                                  |
| 28  | `Reverse`         | ✔      | Reverse a slice; returns new slice.                                     |
| 29  | `Shift`           | ✔      | Return slice without first element + pointer to shifted value.          |
| 30  | `Slice`           | ✔      | Sub-slice from `start` to `end` (exclusive), supports negative indexes. |
| 31  | `Some`            | ✔      | `true` if any element passes the predicate.                             |
| 32  | `Sort`            | ❌     | Use the standard library's `sort` / `slices` packages.                  |
| 33  | `Splice`          | ✔      | Remove and/or insert elements at an index.                              |
| 34  | `ToString`        | ✔      | Comma-separated string representation.                                  |
| 35  | `UnShift`         | ✔      | Prepend elements; returns new slice.                                    |
| 36  | `ValueOf`         | ✔      | Identity — returns the slice itself.                                    |
| 37  | `With`            | ✔      | Return a copy with the element at `index` replaced.                     |

For full signatures, runnable examples, and per-method docs, see
[**pkg.go.dev**](https://pkg.go.dev/github.com/bube054/go-js-array-methods/v2/array).

## Design notes

- **Immutability.** Every function returns a new slice. The original input
  is never modified — even for "mutating" JS methods like `Push` or
  `Splice`. This keeps reasoning local and avoids spooky aliasing bugs.
- **Negative indexes.** `At(-1)` returns the last element, `Slice(s, -2, -1)`
  works as in JS. Out-of-range access returns an error rather than
  panicking.
- **Strict variants.** `MapStrict`, `ReduceStrict`, `ReduceRightStrict`
  preserve the input type `T` instead of returning `any`. Use these when
  you don't need to change the element type.
- **Generic `Map`.** As of v2, `Map[T, V]` returns `[]V`. The compiler
  usually infers `V` from your callback's return type, so call sites stay
  clean: `array.Map(nums, func(n, _ int, _ []int) string { ... })`.
- **Method limitation on `Array[T].Map`.** Because Go does not support
  type parameters on methods, the `Map` *method* is locked to
  `Array[any]`. For full type safety, use the `array.Map` function
  directly.

## Examples

A few common patterns. See [`array/example_test.go`](./array/example_test.go)
for many more (also rendered on pkg.go.dev).

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
[`array/benchmark_test.go`](./array/benchmark_test.go). Run it with:

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
