# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [2.2.0] - 2026-11-12

### 🐛 Fixed

- **Inputs are never changed:** `Push` could write into spare capacity of the slice you passed in, so two `Push` calls on the same slice could overwrite each other, and pushing onto a result of `Pop`, `Shift`, or `Slice` could change the original. `Push` now always builds a new slice, and `Pop`, `Shift`, and `Slice` cap their result's capacity so appending to it allocates instead.
- **`Splice` and `Fill` accept JavaScript's valid edge cases:** inserting at the end (`index == len(slice)`), working on empty or `nil` slices, and empty ranges used to return an error.
- **`Array.Includes`, `Array.IndexOf`, and `Array.LastIndexOf` no longer panic** on values that `==` cannot compare, such as arrays of slices or structs holding a slice in an interface field. They fall back to `reflect.DeepEqual`, and plain types like `int` still make no allocations.
- **README `ReduceAs` example** now compiles (`ReduceAs` returns one value).

### 🚀 Added

- **`Entry` fields are readable:** `Entry` now has exported `Index` and `Value` fields, so code outside the package can read what `Entries` returns.
- **`Unshift`**, matching JavaScript's spelling, as a function and an `Array` method.

### 🔄 Changed

- **Error and panic messages follow Go style:** `Splice` and `With` say `index: N out of range` (like `At`), and `ReduceAs`/`ReduceRightAs` panics start with `array:` instead of `TypeError:`.
- **Docs match the benchmarks:** removed the "sub-nanosecond" and blanket "zero-allocation" claims, corrected `Join` (one allocation, for the result string), and documented which results share memory with the input (`Slice`, `Pop`, `Shift`, and the pointer from `Find`/`FindLast`). See [#6](https://github.com/bube054/go-js-array-methods/issues/6) for why `Find` keeps returning a pointer into the slice.

### ⚠️ Deprecated

- **`UnShift`** (function and method). Use `Unshift`; the old name keeps working.

### 🛡️ Chore & Testing

- **CI** now runs on Go 1.23, oldstable, and stable, with a `gofmt` check, `go vet`, and the race detector.
- **Tests from outside the package** check that inputs are never changed and that JavaScript edge cases work. Coverage stays at 100%.
- Removed the retired Go Report Card badge.

## [2.1.0] - 2026-05-26

### 🚀 Added

- **Go 1.23 Native Iterators:** Introduced `EntriesIterator()`, `Keys()`, and `Values()` methods returning `iter.Seq2` and `iter.Seq`. This enables zero-allocation, lazy iteration over massive datasets (bypassing eager matrix construction).
- **Zero-Allocation Reducers:** Added `ReduceAs` and `ReduceRightAs` primitives. These provide explicitly typed fast-paths that bypass Go's `any` interface boxing entirely, enabling 0-alloc transformations.
- **Executable Examples:** Added comprehensive `ExampleFunc_name` test wrappers for all new methods to ensure `pkg.go.dev` documentation provides dynamically verified usage patterns.

### ⚡ Performance

- **Sub-Nanosecond Fast Paths:** Validated and locked in core operations (like `ConvertIndex` and primitive `Join`) running at hardware instruction speeds (< 2 ns/op).
  _Correction, 2026-11-12: this holds for `ConvertIndex` (about 0.4 ns/op), but not for `Join`, which makes one allocation for the result string and takes about 30 µs for 1,000 ints. See [2.2.0]._
- **Iterator Memory Scaling:** Replaced eager matrix generation on massive loops, dropping memory footprints from ~1.6MB to exactly `0 B/op` at $N=100,000$ scales.

### 🔄 Changed

- **Toolchain Requirement:** Bumped the `go` directive in `go.mod` to `1.23` to natively support the new range-over-func mechanics. _(Note: This is a toolchain minimum bump, not a breaking API change for the module itself)._

### 🛡️ Chore & Testing

- **100% Test Coverage:** Hardened the test suite to achieve and strictly enforce 100% statement coverage across the entire package.

### 📝 Documentation

- **README Overhaul:** Added the "Design & Performance Philosophy" section to explicitly document the architectural trade-offs between dynamic JS-fidelity methods (like standard `Reduce`) and static, bare-metal speed pipelines (`ReduceAs`).
- Added advanced pattern examples demonstrating memory-free lazy iteration and strict-type reduction.

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

## [1.0.0] - 2024-04-08

First stable release. Module path `github.com/bube054/go-js-array-methods`,
requires Go 1.22.

Initial functional API: `At`, `Concat`, `ConvertIndex`, `CopyWithin`, `Entries`, `Every`,
`Fill`, `Filter`, `Find`, `FindIndex`, `FindLast`, `FindLastIndex`,
`ForEach`, `Includes`, `IndexOf`, `Join`, `LastIndexOf`, `Map`, `MapStrict`,
`Pop`, `Push`, `Reduce`, `ReduceRight`, `ReduceStrict`, `ReduceRightStrict`,
`Reverse`, `Shift`, `Slice`, `Some`, `Splice`, `ToString`, `UnShift`,
`ValueOf`, `With`.

Tags `v0.0.0` to `v0.0.5` (2024-02-10 to 2024-04-04) were early pre-releases
leading up to 1.0.0.

[Unreleased]: https://github.com/bube054/go-js-array-methods/compare/v2.2.0...HEAD
[2.2.0]: https://github.com/bube054/go-js-array-methods/compare/v2.1.0...v2.2.0
[2.1.0]: https://github.com/bube054/go-js-array-methods/compare/v2.0.0...v2.1.0
[2.0.0]: https://github.com/bube054/go-js-array-methods/compare/v1.0.0...v2.0.0
[1.0.0]: https://github.com/bube054/go-js-array-methods/releases/tag/v1.0.0
