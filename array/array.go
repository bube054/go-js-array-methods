package array

// The At() method returns an indexed element from the array and returns a possible error relating to out of range indexes.
func (a Array[T]) At(index int) (T, error) {
	return At([]T(a), index)
}

// The Concat() method concatenates the array with other slices and returns a new concatenated array.
func (a Array[T]) Concat(slices ...Array[T]) Array[T] {
	b := make([][]T, len(slices))
	for i, s := range slices {
		b[i] = []T(s)
	}
	return Array[T](Concat([]T(a), b...))
}

// The CopyWithin() method copies array elements to another position in the array and returns the modified array.
func (a Array[T]) CopyWithin(target, start, end int) (Array[T], error) {
	result, err := CopyWithin([]T(a), target, start, end)
	return Array[T](result), err
}

// The Entries() method returns an array of [index, element] pairs for each element in the array.
func (a Array[T]) Entries() []Entry[T] {
	return Entries([]T(a))
}

// The Every() method tests whether all elements in the array pass the provided predicate function.
func (a Array[T]) Every(fn Predicate[T]) bool {
	return Every([]T(a), fn)
}

// The Fill() method fills array elements from a start index to an end index with a static value.
func (a Array[T]) Fill(value T, start, end int) (Array[T], error) {
	result, err := Fill([]T(a), value, start, end)
	return Array[T](result), err
}

// The Flat() method flattens nested arrays up to the specified depth.
// The Flat() method returns a new array with the sub-array elements concatenated into it. The depth parameter specifies how deep a nested array structure should be flattened. The default is 1.
// Until Generic Methods are implemented in Go, the Flat() method can only return Array[any] since it needs to handle slices of any type. Once Generic Methods are available, we can update the Flat() method to return Array[T] and handle type assertions accordingly.
func (a Array[T]) Flat(depths ...int) (Array[any], error) {
	// Convert Array[T] to []any for flattening
	anySlice := make([]any, len(a))
	for i, v := range a {
		anySlice[i] = v
	}
	result, err := Flat[any](anySlice, depths...)
	return Array[any](result), err
}

// The Filter() method creates a new array with elements that pass the provided predicate function.
func (a Array[T]) Filter(fn Predicate[T]) Array[T] {
	return Array[T](Filter([]T(a), fn))
}

// The Find() method returns the first element that passes the provided predicate function.
func (a Array[T]) Find(fn Predicate[T]) *T {
	return Find([]T(a), fn)
}

// The FindIndex() method returns the index of the first element that passes the provided predicate function.
func (a Array[T]) FindIndex(fn Predicate[T]) int {
	return FindIndex([]T(a), fn)
}

// The FindLast() method returns the last element that passes the provided predicate function.
func (a Array[T]) FindLast(fn Predicate[T]) *T {
	return FindLast([]T(a), fn)
}

// The FindLastIndex() method returns the index of the last element that passes the provided predicate function.
func (a Array[T]) FindLastIndex(fn Predicate[T]) int {
	return FindLastIndex([]T(a), fn)
}

// The ForEach() method executes a provided function once for each array element.
func (a Array[T]) ForEach(fn ForEachFunc[T]) {
	ForEach([]T(a), fn)
}

// The Includes() method determines whether the array includes a certain element.
func (a Array[T]) Includes(element T, start *int) bool {
	return Includes([]T(a), element, start)
}

// The IndexOf() method returns the first index at which a given element can be found in the array.
func (a Array[T]) IndexOf(element T, start *int) int {
	return IndexOf([]T(a), element, start)
}

// The Join() method joins all elements of the array into a single string using the provided separator.
// Elements are converted to their string form, so this works for any Array[T].
// If no separator is provided, a comma is used, matching JavaScript Array.join().
func (a Array[T]) Join(separator ...string) string {
	sep := OptionalParam(separator, ",")

	return Join([]T(a), &sep)
}

// The ToString() method returns a string representation of the array.
// It mirrors JavaScript Array.toString() by joining elements with commas.
func (a Array[T]) ToString() string {
	return a.Join()
}

// The LastIndexOf() method returns the last index at which a given element can be found in the array.
func (a Array[T]) LastIndexOf(element T, start *int) int {
	return LastIndexOf([]T(a), element, start)
}

// The Map() method creates a new array with the results of calling a function for every array element.
func (a Array[T]) Map(fn MapFunc[T, any]) Array[any] {
	return Array[any](Map([]T(a), fn))
}

// The MapStrict() method creates a new array of the same type with the results of calling a function for every array element.
func (a Array[T]) MapStrict(fn MapFuncStrict[T]) Array[T] {
	return Array[T](MapStrict([]T(a), fn))
}

// The Pop() method removes the last element from the array and returns the modified array and the removed element.
func (a Array[T]) Pop() (Array[T], *T) {
	result, popped := Pop([]T(a))
	return Array[T](result), popped
}

// The Push() method adds one or more elements to the end of the array and returns the new array.
func (a Array[T]) Push(elements ...T) Array[T] {
	return Array[T](Push([]T(a), elements...))
}

// The Reduce() method executes a reducer function for each array element and returns a single value.
func (a Array[T]) Reduce(fn ReduceFunc[T], initialValue any) (any, error) {
	return Reduce([]T(a), fn, initialValue)
}

// The ReduceRight() method executes a reducer function for each array element (from right to left) and returns a single value.
func (a Array[T]) ReduceRight(fn ReduceFunc[T], initialValue any) (any, error) {
	return ReduceRight([]T(a), fn, initialValue)
}

// The ReduceStrict() method executes a reducer function for each array element and returns a value of the same type.
func (a Array[T]) ReduceStrict(fn ReduceStrictFunc[T], initialValue *T) (T, error) {
	return ReduceStrict([]T(a), fn, initialValue)
}

// The ReduceRightStrict() method executes a reducer function for each array element (from right to left) and returns a value of the same type.
func (a Array[T]) ReduceRightStrict(fn ReduceStrictFunc[T], initialValue *T) (T, error) {
	return ReduceRightStrict([]T(a), fn, initialValue)
}

// The Reverse() method reverses the array and returns a new reversed array.
func (a Array[T]) Reverse() Array[T] {
	return Array[T](Reverse([]T(a)))
}

// The Shift() method removes the first element from the array and returns the modified array and the removed element.
func (a Array[T]) Shift() (Array[T], *T) {
	result, shifted := Shift([]T(a))
	return Array[T](result), shifted
}

// The Slice() method returns a shallow copy of a portion of the array.
func (a Array[T]) Slice(start, end int) (Array[T], error) {
	result, err := Slice([]T(a), start, end)
	return Array[T](result), err
}

// The Some() method tests whether at least one element in the array passes the provided predicate function.
func (a Array[T]) Some(fn Predicate[T]) bool {
	return Some([]T(a), fn)
}

// The Splice() method changes the contents of an array by removing or replacing existing elements and/or adding new elements.
func (a Array[T]) Splice(index int, howMany *int, items ...T) (Array[T], error) {
	result, err := Splice([]T(a), index, howMany, items...)
	return Array[T](result), err
}

// The UnShift() method adds one or more elements to the beginning of the array and returns the new array.
func (a Array[T]) UnShift(elements ...T) Array[T] {
	return Array[T](UnShift([]T(a), elements...))
}

// The ValueOf() method returns the array itself (identity function).
func (a Array[T]) ValueOf() Array[T] {
	return Array[T](ValueOf([]T(a)))
}

// The With() method returns a new array with the element at the given index replaced with the given value.
func (a Array[T]) With(index int, value T) (Array[T], error) {
	result, err := With([]T(a), index, value)
	return Array[T](result), err
}
