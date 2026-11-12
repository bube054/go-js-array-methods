// Package array provides a high-performance utility suite for manipulating Go
// slices, inspired by the ergonomics of JavaScript array methods.
//
// It leverages Go generics to deliver type-safe functional programming
// paradigms—such as Map, Filter, Reduce, and Flat—without the boilerplate,
// allowing developers to write clean, declarative code.
//
// # Performance
//
// The typed functions compile down to plain loops with no interface boxing.
// Searches and folds such as Includes, Some, ReduceStrict, and ReduceAs make no
// heap allocations of their own, and functions that build a new slice, such as
// Filter and Map, make exactly one, for the result. Your callbacks can still
// allocate. Run the package benchmarks to see the numbers on your machine.
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
// At that time, the strictly-typed signatures will become the default behavior.
package array
