# Contributing to go-js-array-methods

Thanks for your interest in contributing! This guide covers the workflow,
expectations, and conventions for the project.

## Quick start

```bash
git clone https://github.com/bube054/go-js-array-methods.git
cd go-js-array-methods
go mod download
go test ./...
```

That's it — no extra tooling is required to run the test suite.

## Project layout

```
.
├── array/                 # the library package
│   ├── array.go           # Array[T] type and its receiver methods
│   ├── types.go           # shared callback type definitions
│   ├── utils.go           # shared helpers (ConvertIndex, OptionalParam)
│   ├── <method>.go        # one file per JS-style method (filter.go, map.go, ...)
│   ├── <method>_test.go   # unit tests for that method
│   ├── example_test.go    # runnable examples (also rendered on pkg.go.dev)
│   └── benchmark_test.go  # benchmarks
├── .github/workflows/     # CI configuration
├── CHANGELOG.md
├── CONTRIBUTING.md
└── README.md
```

When you add a new JS-style method, add a `<method>.go` and a matching
`<method>_test.go` in the same style as the existing files.

## Development workflow

1. **Open an issue first** for non-trivial changes so we can agree on the
   approach before you write code.
2. **Fork and branch** off `main`. Use a descriptive branch name like
   `feat/add-flatmap` or `fix/concat-empty-slice`.
3. **Make your change** with tests and (where applicable) an example.
4. **Run the full check locally** — see "Before submitting" below.
5. **Open a pull request** against `main`. Reference the issue it closes.

## Coding conventions

- **Format with `gofmt`** — `go fmt ./...` before committing.
- **Names mirror JavaScript.** `Filter`, `Map`, `ForEach`, etc. Keep the
  Go-idiomatic capitalization (no `forEach`).
- **Original slice is never mutated.** Every method returns a new slice.
  This is the core design promise of the package.
- **Errors over panics.** Out-of-range indexes return `(zero, error)`; they
  do not panic. See `ConvertIndex` in `utils.go` for the standard pattern.
- **Doc comments on every exported symbol.** Start with the symbol's name
  (`// The Filter() function ...`) so it surfaces correctly on
  [pkg.go.dev](https://pkg.go.dev/).
- **Generics:** prefer `comparable` only when the function actually needs
  equality (e.g. `Includes`, `IndexOf`). Otherwise `any` is fine and gives
  callers more flexibility.
- **Optional parameters** are modeled as a trailing variadic plus
  `OptionalParam` from `utils.go`. Don't introduce a different convention.

## Tests

Every new function or method needs tests. We use the standard `testing`
package with table-driven subtests via `t.Run`:

```go
func TestFoo(t *testing.T) {
    tests := []struct {
        name     string
        input    []int
        expected []int
    }{
        {"empty input", []int{}, []int{}},
        {"happy path", []int{1, 2, 3}, []int{1, 2, 3}},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := Foo(tt.input)
            if !reflect.DeepEqual(got, tt.expected) {
                t.Errorf("Foo() = %v, want %v", got, tt.expected)
            }
        })
    }
}
```

Cover at minimum: the happy path, empty input, boundary values, and any
error paths.

## Examples

If your function is part of the public API, add an `Example` function in
`array/example_test.go`. It will be both compiled (so it can't drift) and
rendered on pkg.go.dev:

```go
func ExampleFilter() {
    nums := []int{1, 2, 3, 4, 5}
    even := array.Filter(nums, func(n, _ int, _ []int) bool { return n%2 == 0 })
    fmt.Println(even)
    // Output: [2 4]
}
```

The `// Output:` comment is what the test harness checks.

## Benchmarks

For anything in the hot path or anything that uses `reflect`, add a
benchmark in `array/benchmark_test.go`:

```go
func BenchmarkFilter(b *testing.B) {
    nums := makeInts(10_000)
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        _ = array.Filter(nums, func(n, _ int, _ []int) bool { return n%2 == 0 })
    }
}
```

Run with `go test -bench=. -benchmem ./array/`.

## Before submitting

```bash
go fmt ./...
go vet ./...
go test -race -cover ./...
```

All three should pass with no output beyond the test summary. CI runs
`go test -v ./...` on every PR, so a green local run usually means a
green CI run.

## Commit messages

Conventional-Commits-style prefixes are appreciated but not required:

- `feat: add FlatMap function`
- `fix: handle empty slice in Reduce`
- `refactor: simplify ConvertIndex`
- `docs: clarify Map type parameter`
- `test: add edge cases for Splice`

Keep the subject line under ~70 characters; put detail in the body.

## Pull request expectations

- Tests pass, including new tests for your change.
- New exported symbols have doc comments and (ideally) an example.
- `CHANGELOG.md` updated under `[Unreleased]` if your change is
  user-visible.
- One logical change per PR. If you're tempted to write "and also..." in
  the description, split it.

## Reporting bugs

Open an issue with:

- Go version (`go version`)
- A minimal reproduction (the smaller the better)
- What you expected to happen vs. what actually happened

## Questions?

Open a discussion or ping the maintainer on the issue tracker. We're happy
to help.
