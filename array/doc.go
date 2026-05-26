// Package array provides a high-performance, zero-allocation utility suite for
// manipulating Go slices, inspired by the ergonomics of JavaScript array methods.
//
// It leverages Go generics to deliver type-safe functional programming
// paradigms—such as Map, Filter, Reduce, and Flat—without the boilerplate,
// allowing developers to write clean, declarative code.
//
// # Performance
//
// This package is engineered for bare-metal execution. The core iterative
// functions are designed to bypass Go's interface boxing and dynamic heap
// allocations entirely. Hot-path operations, including Includes, Filter, and
// the strictly-typed ReduceAs, operate at 0 allocs/op, matching the speed
// of native for-loops.
//
// # Type Safety & Migration Roadmap
//
// To preserve backward compatibility, this package currently exposes both
// legacy loosely-typed methods (which accept and return 'any') and modern,
// strictly-typed methods.
//
// For maximum performance and compile-time safety, always prefer the typed variants:
//   - For type-preserving operations: Filter, MapStrict, ReduceStrict
//   - For type-transforming operations: ReduceAs, ReduceRightAs
//
// Deprecation Notice: The legacy methods (such as Reduce) incur minor
// interface boxing overhead and will be deprecated in an upcoming major release.
// At that time, the strict zero-allocation signatures will become the default behavior.
package array
