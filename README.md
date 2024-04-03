# go-js-array-methods

| s/n | Function   | Description                                                    |
| --- | -------- | -------------------------------------------------------------- |
| 1   | At ✔     | The At() function returns an indexed element from a slice.     |
| 2   | Concat ✔ | The Concat() function Concatenates (joins) two or more slices. | The Concat() function returns a new array, containing the joined slices. The Concat() function does not change the existing slices. |
| 3 | CopyWithin ✔ | The CopyWithin() function copies slice elements to another position in an slice. The CopyWithin() function overwrites the existing values. The CopyWithin() function does not add items to the slice. |
| 4 | Entries ✔ | The entries() function returns an Slice with structs with field/value: {index 0, element "Banana"}. The entries() function does not change the original slice. |
| 5 | Every ✔ | The Every() function executes a function for each slice element. The Every() function returns true if the function returns true for all elements. The Every() function returns false if the function returns false for one element. The Every() function does not execute the function for empty elements. The Every() function does not change the original slice. |
| 6 | Fill ✔ | The Fill() function fills specified elements in an slice with a value. The fill() method does not overwrites the original array.|
| 7 | Filter ✔ | The filter() function creates a new slice filled with elements that pass a test provided by a function. The filter() function does not execute the function for empty elements. The filter() function does not change the original slice. |
| 8 | Find ✔ | The Find() function returns the value of the first element that passes a test. The Find() function executes a function for each slice element. The Find() function returns undefined if no elements are found. The Find() function does not execute the function for empty elements. The Find() function does not change the original slice. |
| 7 | FindIndex ✔ | The FindIndex() function executes a function for each array element. The FindIndex() function returns the index (position) of the first element that passes a test. The FindIndex() function returns -1 if no match is found. The FindIndex() function does not execute the function for empty array elements. The FindIndex() function does not change the original array. |
| 7 | FindLast ✔ | The FindLast() function returns the value of the last element that passes a test. The FindLast() function executes a function for each slice element. The FindLast() function returns -1 if no elements are found. The FindLast() function does not execute the function for empty elements. The FindLast() function does not change the original slice. |
| 7 | FindLastIndex ✔ | The FindLastIndex() function executes a function for each array element. The FindLastIndex() function returns the index (position) of the last element that passes a test. The FindLastIndex() function returns -1 if no match is found. The FindLastIndex() function does not execute the function for empty slice elements.The FindLastIndex() function does not change the original slice. |
| 8 | Flat ❌ | ----- |
| 7 | FlatMap ❌ | ----- |
| 9 | ForEach ✔ | The ForEach() function calls a function for each element in an slice. The ForEach() function is not executed for empty elements. |
| 10 | Includes ✔ | The Includes() function returns true if an slice contains a specified value. The Includes() function returns false if the value is not found. The Includes() function is case sensitive. |
| 11 | IndexOf ✔ | The IndexOf() function returns the first index (position) of a specified value. The IndexOf() function returns -1 if the value is not found. The IndexOf() function starts at a specified index and searches from left to right (from the given start postion to the end of the array). By default the search starts at the first element and ends at the last. Negative start values counts from the last element (but still searches from left to right). |
| 12 | Join ✔ | The Join() function returns an slice as a string. The Join() function does not change the original slice. Any separator can be specified. The default is comma (,). |
| 13 | Keys ❌ | ----- |
| 14 | LastIndexOf ✔ | The LastIndexOf() function returns the last index (position) of a specified value. The LastIndexOf() function returns -1 if the value is not found. The LastIndexOf() starts at a specified index and searches from right to left (from the given position to the beginning of the slice). By default the search starts at the last element and ends at the first. Negative start values counts from the last element (but still searches from right to left). |
| 15 | Map ✔ | Map() creates a new slice from calling a function for every slice element. Map creates a new slice of any type. Map() does not execute the function for empty elements. Map() does not change the original slice. |
| 16 | MapStrict ✔ | Map() creates a new slice from calling a function for every slice element. Map creates a new slice of the original slice type. Map() does not execute the function for empty elements. Map() does not change the original slice. |
| 17 | Pop ✔ | The Pop() function removes (pops) the last element of an slice. The Pop() function does not change the original slice. The Pop() function returns the new slice without the last element and a pointer of removed element. |
| 18 | Push ✔ | The Push() function adds new items to the end of an slice. The Push() function returns the new slice. |
| 19 | Reduce ✔ | The Reduce() function executes a Reducer function for slice element. The Reduce() function returns the function's accumulated result and an an error. The Reduce() function does not execute the function for empty slice elements. The Reduce() function does not change the original slice. Note at the first callback, there is no return value from the previous callback. Normally, alice element 0 is used as initial value, and the iteration starts from slice element 1. If an initial value is supplied, this is used, and the iteration starts from slice element 0. |
| 20 | ReduceStrict ✔ | The ReduceStrict() function executes a Reducer function for slice element. The ReduceStrict() function returns the function's accumulated result(typed with the slices type) and an an error. The ReduceStrict() function does not execute the function for empty slice elements. The ReduceStrict() function does not change the original slice. Note at the first callback, there is no return value from the previous callback. Normally, alice element 0 is used as initial value, and the iteration starts from slice element 1. If an initial value is supplied, this is used, and the iteration starts from slice element 0. |
| 21 | ReduceRight ✔ | The Reduce() function executes a Reducer function for slice element. The ReduceRight() function works from right to left. The Reduce() function returns the function's accumulated result and an an error. The Reduce() function does not execute the function for empty slice elements. The Reduce() function does not change the original slice. Note at the first callback, there is no return value from the previous callback. Normally, alice element 0 is used as initial value, and the iteration starts from slice element 1. If an initial value is supplied, this is used, and the iteration starts from slice element 0. |
| 22 | ReduceRightStrict ✔ | The ReduceStrict() function executes a Reducer function for slice element. The ReduceRightStrict() function works from right to left . The ReduceStrict() function returns the function's accumulated result(typed with the slices type) and an an error. The ReduceStrict() function does not execute the function for empty slice elements. The ReduceStrict() function does not change the original slice. Note at the first callback, there is no return value from the previous callback. Normally, alice element 0 is used as initial value, and the iteration starts from slice element 1. If an initial value is supplied, this is used, and the iteration starts from slice element 0. |
| 23 | Filter ✔ | x |
| 24 | Filter ✔ | x |
| 25 | Filter ✔ | x |
| 26 | Filter ✔ | x |
| 27 | Filter ✔ | x |
| 28 | Filter ✔ | x |
| 29 | Filter ✔ | x |
| 30 | Filter ✔ | x |
| 31 | Filter ✔ | x |
| 32 | Filter ✔ | x |
| 33 | Filter ✔ | x |
| 34 | Filter ✔ | x |
| 35 | Filter ✔ | x |
| 36 | Filter ✔ | x |
| 37 | Filter ✔ | x |
| 38 | Filter ✔ | x |