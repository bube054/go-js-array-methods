# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [2.0.0] - 2026-04-15

### Added

- **`Array[T]` type** — a slice-backed type with receiver methods for every
  package-level function, enabling a fluent, chainable, JavaScript-like API:
  ```go
  arr := array.Array[int]{1, 2, 3, 4, 5}
  result := arr.Filter(isEven).Push(6).Reverse()
  ```
- **`Flat[T any]` function** — flattens nested `[]any` slices to a typed `[]T`
  with configurable depth (default 1). Mirrors JavaScript's `Array.prototype.flat`.
- **`Array[T].Flat(depths ...int)` method** — method form returning `Array[any]`
  (locked to `any` because Go does not support generic type parameters on methods).
- **`OptionalParam[T]` helper** in `utils.go` — small utility for treating a
  trailing variadic as a single optional parameter with a default.
- **GitHub Actions CI workflow** (`.github/workflows/tests.yml`) — runs
  `go test -v ./...` on every push to `main` and on every pull request.
- **`example_test.go`** files with runnable examples that surface on
  [pkg.go.dev](https://pkg.go.dev/).
- **`benchmark_test.go`** with benchmarks for the most common operations
  (`Filter`, `Map`, `Reduce`, `Reverse`, `Flat`, the `Array[T]` chained pipeline).
- **`CONTRIBUTING.md`** explaining the dev loop, test/lint expectations,
  and PR conventions.

### Changed

- **BREAKING — `Map` is now generic over its return type.** The signature
  changed from:
  ```go
  func Map[T comparable](slice []T, fn MapFunc[T]) []any
  ```
  to:
  ```go
  func Map[T comparable, V any](slice []T, fn MapFunc[T, V]) []V
  ```
  Callers no longer need to type-assert `[]any` results — the returned slice
  is strongly typed. Update existing call sites by adding the second type
  parameter, e.g. `array.Map[int, string](nums, fn)`, or rely on inference
  from the callback's return type.
- **BREAKING — `MapFunc` is now `MapFunc[T comparable, V any]`** (was
  `MapFunc[T comparable]` returning `any`).
- **`Join` is now generic** — accepts `[]T` of any element type and converts
  each element with `fmt.Sprint` before joining. Internally uses
  `strings.Join` for efficiency. Previously restricted to `[]string`.
- **`ToString` is now generic** — accepts `[]T` of any element type. Previously
  restricted to `[]string`.
- **`Array[T].Concat` accepts `...Array[T]`** instead of `...[]T` for
  consistency with the receiver type. The package-level `Concat` still
  takes `...[]T`.
- **README rewritten** as UTF-8 (previously UTF-16 LE) with badges, the new
  `Array[T]` API documented, and the function table updated.

### Removed

- **`main.go` deleted** — this is a library package; the demo `main.go` no
  longer ships with the module.

### Notes on Migration to v2

Because of the breaking `Map` signature, the module path is bumped to
`github.com/bube054/go-js-array-methods/v2`. Update your imports:

```go
// before
import "github.com/bube054/go-js-array-methods/array"

// after
import "github.com/bube054/go-js-array-methods/v2/array"
```

If you only use `Filter`, `Reduce`, `Push`, `Pop`, etc., the migration is
import-path-only. `Map` callers need the type-parameter update above.

## [1.x] - prior to 2026-04-15

Initial functional API: `At`, `Concat`, `CopyWithin`, `Entries`, `Every`,
`Fill`, `Filter`, `Find`, `FindIndex`, `FindLast`, `FindLastIndex`,
`ForEach`, `Includes`, `IndexOf`, `Join`, `LastIndexOf`, `Map`, `MapStrict`,
`Pop`, `Push`, `Reduce`, `ReduceRight`, `ReduceStrict`, `ReduceRightStrict`,
`Reverse`, `Shift`, `Slice`, `Some`, `Splice`, `ToString`, `UnShift`,
`ValueOf`, `With`.

[Unreleased]: https://github.com/bube054/go-js-array-methods/compare/v2.0.0...HEAD
[2.0.0]: https://github.com/bube054/go-js-array-methods/releases/tag/v2.0.0
