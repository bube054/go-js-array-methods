# go-js-array-methods

| s/n | Function   | Description                                                    |
| --- | -------- | -------------------------------------------------------------- |
| 1   | At ✔     | The At() function returns an indexed element from a slice and an error. The At() function returns the same as []     |
| 2   | Concat ✔ | The Concat() function Concatenates (joins) two or more slices. | The Concat() function returns a new slice, containing the joined slices. The Concat() function does not change the existing slices. |
| 3 | CopyWithin ✔ | The CopyWithin() function copies slice elements to another position in an slice. The CopyWithin() function overwrites the existing values. The CopyWithin() function does not add items to the slice. |
| 4 | Entries ✔ | The entries() function returns an Slice with structs with field/value: {index 0, element "Banana"}. The entries() function does not change the original slice. |
| 5 | Every ✔ | The Every() function executes a function for each slice element. The Every() function returns true if the function returns true for all elements. The Every() function returns false if the function returns false for one element. The Every() function does not execute the function for empty elements. The Every() function does not change the original slice. |
| 6 | Fill ✔ | The Fill() function fills specified elements in an slice with a value. The fill() method does not overwrites the original slice.|
| 7 | Filter ✔ | The filter() function creates a new slice filled with elements that pass a test provided by a function. The filter() function does not execute the function for empty elements. The filter() function does not change the original slice. |
| 8 | Find ✔ | The Find() function returns the value of the first element that passes a test. The Find() function executes a function for each slice element. The Find() function returns undefined if no elements are found. The Find() function does not execute the function for empty elements. The Find() function does not change the original slice. |
| 7 | FindIndex ✔ |The FindIndex() function executes a function for each slice element. The FindIndex() function returns the index (position) of the first element that passes a test. The FindIndex() function returns -1 if no match is found. The FindIndex() function does not execute the function for empty slice elements. The FindIndex() function does not change the original slice. |
| 7 | FindLast ✔ | The FindLast() function returns the value of the last element that passes a test. The FindLast() function executes a function for each slice element. The FindLast() function returns -1 if no elements are found. The FindLast() function does not execute the function for empty elements. The FindLast() function does not change the original slice. |
| 7 | FindLastIndex ✔ | The FindLastIndex() function executes a function for each array element. The FindLastIndex() function returns the index (position) of the last element that passes a test. The FindLastIndex() function returns -1 if no match is found. The FindLastIndex() function does not execute the function for empty slice elements.The FindLastIndex() function does not change the original slice. |
| 8 | Flat ❌ | ————————— |
| 7 | FlatMap ❌ | ————————— |
| 9 | ForEach ✔ | The ForEach() function calls a function for each element in an slice. The ForEach() function is not executed for empty elements. |
| 10 | Includes ✔ | The Includes() function returns true if an slice contains a specified value. The Includes() function returns false if the value is not found. The Includes() function is case sensitive. |
| 11 | IndexOf ✔ | The IndexOf() function returns the first index (position) of a specified value. The IndexOf() function returns -1 if the value is not found. The IndexOf() function starts at a specified index and searches from left to right (from the given start postion to the end of the array). By default the search starts at the first element and ends at the last. Negative start values counts from the last element (but still searches from left to right). |
| 12 | Join ✔ | The Join() function returns an slice as a string. The Join() function does not change the original slice. Any separator can be specified. The default is comma (,). Join function only works for strings. |
| 13 | Keys ❌ | ————————— |
| 14 | LastIndexOf ✔ | The LastIndexOf() function returns the last index (position) of a specified value. The LastIndexOf() function returns -1 if the value is not found. The LastIndexOf() starts at a specified index and searches from right to left (from the given position to the beginning of the slice). By default the search starts at the last element and ends at the first. Negative start values counts from the last element (but still searches from right to left). |
| 15 | Map ✔ | Map() creates a new slice from calling a function for every slice element. Map creates a new slice of any type. Map() does not execute the function for empty elements. Map() does not change the original slice. |
| 16 | MapStrict ✔ | Map() creates a new slice from calling a function for every slice element. Map creates a new slice of the original slice type. Map() does not execute the function for empty elements. Map() does not change the original slice. |
| 17 | Pop ✔ | The Pop() function removes (pops) the last element of an slice. The Pop() function does not change the original slice. The Pop() function returns the new slice without the last element and a pointer of removed element. |
| 18 | Push ✔ | The Push() function adds new items to the end of an slice. The Push() function returns the new slice. |
| 19 | Reduce ✔ | The Reduce() function executes a Reducer function for slice element. The Reduce() function returns the function's accumulated result and an an error. The Reduce() function does not execute the function for empty slice elements. The Reduce() function does not change the original slice. Note at the first callback, there is no return value from the previous callback. Normally, alice element 0 is used as initial value, and the iteration starts from slice element 1. If an initial value is supplied, this is used, and the iteration starts from slice element 0. |
| 20 | ReduceStrict ✔ | The ReduceStrict() function executes a Reducer function for slice element. The ReduceStrict() function returns the function's accumulated result(typed with the slices type) and an an error. The ReduceStrict() function does not execute the function for empty slice elements. The ReduceStrict() function does not change the original slice. Note at the first callback, there is no return value from the previous callback. Normally, alice element 0 is used as initial value, and the iteration starts from slice element 1. If an initial value is supplied, this is used, and the iteration starts from slice element 0. |
| 21 | ReduceRight ✔ | The Reduce() function executes a Reducer function for slice element. The ReduceRight() function works from right to left. The Reduce() function returns the function's accumulated result and an an error. The Reduce() function does not execute the function for empty slice elements. The Reduce() function does not change the original slice. Note at the first callback, there is no return value from the previous callback. Normally, alice element 0 is used as initial value, and the iteration starts from slice element 1. If an initial value is supplied, this is used, and the iteration starts from slice element 0. |
| 22 | ReduceRightStrict ✔ | The ReduceStrict() function executes a Reducer function for slice element. The ReduceRightStrict() function works from right to left . The ReduceStrict() function returns the function's accumulated result(typed with the slices type) and an an error. The ReduceStrict() function does not execute the function for empty slice elements. The ReduceStrict() function does not change the original slice. Note at the first callback, there is no return value from the previous callback. Normally, alice element 0 is used as initial value, and the iteration starts from slice element 1. If an initial value is supplied, this is used, and the iteration starts from slice element 0. |
| 23 | Reverse ✔ | The Reverse() function Reverses the order of the elements in an slice. The Reverse() function does not overwrites the original slice. |
| 24 | Shift ✔ | The Shift() function removes the first item of an slice. The Shift() function changes the original slice. The Shift() function returns the new slice and a pointer to the  shifted element. |
| 25 | Slice ✔ | The Slice() function returns selected elements in an slice, as a new slice and and error. The Slice() function selects from a given start, up to a (not inclusive) given end. The Slice() function does not change the original slice. |
| 26 | Some ✔ | The Some() function checks if any slice elements pass a test (provided as a callback function). The Some() function executes the callback function once for each slice element.The Some() function returns true (and stops) if the function returns true for one of the slice elements. The Some() function returns false if the function returns false for all of the slice elements. The Some() function does not execute the function for empty slice elements. The Some() function does not change the original slice.  |
| 27 | Sort ❌ | ————————— |
| 28 | Splice ✔ | x |
| 29 | ToString ✔ | x |
| 30 | Unshift ✔ | The Unshift() function adds new elements to the beginning of an array. The Unshift() function does not overwrite the original array. |
| 31 | ValueOf ✔ | The ValueOf() function returns the slice itself. The ValueOf() function does not change the original slice. fruits. ValueOf() returns the same as fruits. |
| 32 | With ✔ | The With() function updates a specified slice element. The With() function returns a new slice. The With() function does not change the original slice. |

### Importing The Package
```go
package main

import (
  "github.com/bube054/go-js-array-methods/array"
)

func main() {

}
```

### At()
The At() function returns an indexed element from a slice and an error. The At() function returns the same as [].
 ```go
	fruits := []string{"Banana", "Orange", "Apple", "Mango"}
	want := "Apple"
	index := 2 // negative values also allowed
	result, err := At(fruits, index)

  // error will be present if index is out of range
	if err != nil {
		fmt.Println("Err", err)
	}

  // "Apple"
	fmt.Println("Result:", result)
 ```

### Concat()
The Concat() function Concatenates (joins) two or more slices. | The Concat() function returns a new slice, containing the joined slices. The Concat() function does not change the existing slices.
 ```go
	names1 := []string{"Cecilie", "Lone"}
	names2 := []string{"Emil", "Tobias", "Linus"}
	result := []string{"Cecilie", "Lone", "Emil", "Tobias", "Linus"}

	children := Concat[string](names1, names2) // "Cecilie", "Lone", "Emil", "Tobias", "Linus"

	if !reflect.DeepEqual(result, children) {
		fmt.Println("results and children are not equal.")
	}

	fmt.Println(children)
 ```

### CopyWithin()
The CopyWithin() function copies slice elements to another position in an slice. The CopyWithin() function overwrites the existing values in the new slice. The CopyWithin() function does not add items to the slice.
 ```go
	fruits := []string{"Banana", "Orange", "Apple", "Mango", "Kiwi", "Papaya"}

	newFruits, err := CopyWithin(fruits, -6, -4, 2)

	fmt.Println("New:", newFruits) // [Banana Orange Apple Mango Kiwi Papaya]
	fmt.Println("New:", err) // error occurs because of out of range index.
 ```


### Entries()
The entries() function returns an Slice with structs with field/value: {index 0, element "Banana"}. The entries() function does not change the original slice.
 ```go
	fruits := []string{"Banana", "Orange", "Apple", "Mango", "Kiwi", "Papaya"}

	newFruits := Entries(fruits)

	fmt.Println("New:", newFruits) // [{0 Banana} {1 Orange} {2 Apple} {3 Mango} {4 Kiwi} {5 Papaya}]
 ```

### Every()
The Every() function executes a function for each slice element. The Every() function returns true if the function returns true for all elements. The Every() function returns false if the function returns false for one element. The Every() function does not execute the function for empty elements. The Every() function does not change the original slice.
 ```go
	points := []int{10, 25, 65, 40}

	allPointsGreaterThan50 := func(point int, ind int, list []int) bool {
		if point > 1 {
			return true
		} else {
			return false
		}
	}

	result := Every(points, allPointsGreaterThan50)

	fmt.Println("Result:", result) // some
 ```


### Fill()
The Fill() function fills specified elements in an slice with a value. The fill() method does not overwrites the original slice.
 ```go
	fruits := []string{"Banana", "Orange", "Apple", "Mango"}

	newFruits, err := Fill(fruits, "Kiwi", 1, 3)

	fmt.Println("newFruits:", newFruits) // [Banana Kiwi Kiwi Mango]
	fmt.Println("err:", err) // an error can occur because of out of range index.
 ```

### Filter()
The filter() function creates a new slice filled with elements that pass a test provided by a function. The filter() function does not execute the function for empty elements. The filter() function does not change the original slice.
 ```go
	ages := []int{32, 33, 16, 40}

	checkAdult := func(age int, ind int, list []int) bool {
		if age > 18 {
			return true
		} else {
			return false
		}
	}

	newAges := Filter(ages, checkAdult)

	fmt.Println("newAges:", newAges) // [32 33 40]
 ```

### Find()
The Find() function returns a pointer value of the first element that passes a test. The Find() function executes a function for each slice element. The Find() function returns undefined if no elements are found. The Find() function does not execute the function for empty elements. The Find() function does not change the original slice.
 ```go
	ages := []int{3, 10, 18, 20}

	find18YrsOld := func(el int, ind int, list []int) bool {
		if el > 18 {
			return true
		} else {
			return false
		}
	}

	ans := Find(ages, find18YrsOld)

	if ans != nil {
		fmt.Println("ans:", *ans) // 20
	} else {
		fmt.Println("ans:", ans) // nil
	}
 ```


### FindIndex()
The FindIndex() function executes a function for each slice element. The FindIndex() function returns the index (position) of the first element that passes a test. The FindIndex() function returns -1 if no match is found. The FindIndex() function does not execute the function for empty slice elements. The FindIndex() function does not change the original slice.
 ```go
	ages := []int{3, 10, 18, 20}

	find18YrsOld := func(el int, ind int, list []int) bool {
		if el >= 18 {
			return true
		} else {
			return false
		}
	}

	ans := FindIndex(ages, find18YrsOld)

	fmt.Println("ans:", ans) // 2
 ```

### FindLast()
The FindLast() function returns the value of the last element that passes a test. The FindLast() function executes a function for each slice element. The FindLast() function returns -1 if no elements are found. The FindLast() function does not execute the function for empty elements. The FindLast() function does not change the original slice.
 ```go

	ages := []int{3, 7, 10, 15, 18, 20, 25}

	find18YrsOld := func(el int, ind int, list []int) bool {
		if el >= 18 {
			return true
		} else {
			return false
		}
	}

	ans := FindLast(ages, find18YrsOld)
	fmt.Println("ans:", *ans) // 25
 ```

### FindLastIndex()
The FindLastIndex() function executes a function for each array element. The FindLastIndex() function returns the index (position) of the last element that passes a test. The FindLastIndex() function returns -1 if no match is found. The FindLastIndex() function does not execute the function for empty slice elements.The FindLastIndex() function does not change the original slice.
 ```go
	ages := []int{3, 10, 18, 20}

	find18YrsOld := func(el int, ind int, list []int) bool {
		if el > 18 {
			return true
		} else {
			return false
		}
	}

	ans := FindIndex(ages, find18YrsOld)

	fmt.Println("ans:", ans) // 3
 ```

### ForEach()
The ForEach() function calls a function for each element in an slice. The ForEach() function is not executed for empty elements.
 ```go
	numbers := []int{65, 44, 12, 4}

	sum := 0
	ForEach(numbers, func(el int, ind int, list []int) {
		sum += el
	})

	fmt.Println("Sum:", sum) // 125
 ```

### Includes()
The Includes() function returns true if an slice contains a specified value. The Includes() function returns false if the value is not found. The Includes() function is case sensitive.
 ```go
	fruits := []string{"Banana", "Orange", "Apple", "Mango"}

	start := 1 // start searching from index 1, negative indexes are also allowed
	present := Includes(fruits, "Orange", &start)

	fmt.Println("Present:", present) // true
 ```

### IndexOf()
The IndexOf() function returns the first index (position) of a specified value. The IndexOf() function returns -1 if the value is not found. The IndexOf() function starts at a specified index and searches from left to right (from the given start postion to the end of the array). By default the search starts at the first element and ends at the last. Negative start values counts from the last element (but still searches from left to right).
 ```go
  // example 1
	fruits := []string{"Banana", "Orange", "Apple", "Mango"}

	// start := 1
	present := IndexOf(fruits, "Orange", nil)

	fmt.Println("Present:", present) // 1

  // example 2 
  fruits = []string{"Orange", "Banana", "Orange", "Apple", "Mango"}

	start = 1
	present = IndexOf(fruits, "Orange", &start)

	fmt.Println("Present:", present)
 ```

### Join()
The Join() function returns an slice as a string. The Join() function does not change the original slice. Any separator can be specified. The default is comma (,). Join function only works for strings.
 ```go
  // example 1
	fruits := []string{"Banana", "Orange", "Apple", "Mango"}

	delim := " and "
	newFruits := Join(fruits, &delim)
	fmt.Println("newFruits:", newFruits) // Banana and Orange and Apple and Mango
 ```

### LastIndexOf()
The LastIndexOf() function returns the last index (position) of a specified value. The LastIndexOf() function returns -1 if the value is not found. The LastIndexOf() starts at a specified index and searches from right to left (from the given position to the beginning of the slice). By default the search starts at the last element and ends at the first. Negative start values counts from the last element (but still searches from right to left).
 ```go
	fruits := []string{"Orange", "Apple", "Mango", "Apple", "Banana", "Apple"}

	start := 3
	present := LastIndexOf(fruits, "Apple", &start)

	fmt.Println("Present:", present) // 5
 ```

### Map()
Map() creates a new slice from calling a function for every slice element. Map creates a new slice of any type. Map() does not execute the function for empty elements. Map() does not change the original slice.
 ```go
	numbers := []int{65, 44, 12, 4}

	newNums := Map(numbers, func(e int, i int, s []int) any {
		return e * 10
	})

	fmt.Println("newNums:", newNums) // [650 440 120 40]
 ```

### MapSrict()
Map() creates a new slice from calling a function for every slice element. Map creates a new slice of the original slice type. Map() does not execute the function for empty elements. Map() does not change the original slice.
 ```go
	numbers := []int{65, 44, 12, 4}

	newNums2 := MapStrict(numbers, func(e int, i int, s []int) int {
		return e / 10
	})

	fmt.Println("newNums2:", newNums2) // [6 4 1 0]
 ```

### Pop()
The Pop() function removes (pops) the last element of an slice. The Pop() function does not change the original slice. The Pop() function returns the new slice without the last element and a pointer of removed element.
 ```go
	fruits := []string{"Banana", "Orange", "Apple", "Mango"}

	newFruits, fruit := Pop(fruits)

	fmt.Println("newFruits:", newFruits) // [Banana Orange Apple]
	fmt.Println("fruit:", *fruit) // Mango
 ```

### Push()
The Push() function adds new items to the end of an slice. The Push() function returns the new slice.
 ```go
	fruits := []string{"Banana", "Orange", "Apple", "Mango"}

	newFruits := Push(fruits, "Kiwi", "Lemon")

	fmt.Println("newFruits:", newFruits) // [Banana Orange Apple Mango Kiwi Lemon]
 ```

### Reduce()
The LastIndexOf() function returns the last index (position) of a specified value. The LastIndexOf() function returns -1 if the value is not found. The LastIndexOf() starts at a specified index and searches from right to left (from the given position to the beginning of the slice). By default the search starts at the last element and ends at the first. Negative start values counts from the last element (but still searches from right to left).
 ```go
	fruits := []string{"Orange", "Apple", "Mango", "Apple", "Banana", "Apple"}

	start := 3
	present := LastIndexOf(fruits, "Apple", &start)

	fmt.Println("Present:", present) // 5
 ```

### ReduceStrict()
The LastIndexOf() function returns the last index (position) of a specified value. The LastIndexOf() function returns -1 if the value is not found. The LastIndexOf() starts at a specified index and searches from right to left (from the given position to the beginning of the slice). By default the search starts at the last element and ends at the first. Negative start values counts from the last element (but still searches from right to left).
 ```go
	fruits := []string{"Orange", "Apple", "Mango", "Apple", "Banana", "Apple"}

	start := 3
	present := LastIndexOf(fruits, "Apple", &start)

	fmt.Println("Present:", present) // 5
 ```

### ReduceRight()
The LastIndexOf() function returns the last index (position) of a specified value. The LastIndexOf() function returns -1 if the value is not found. The LastIndexOf() starts at a specified index and searches from right to left (from the given position to the beginning of the slice). By default the search starts at the last element and ends at the first. Negative start values counts from the last element (but still searches from right to left).
 ```go
	fruits := []string{"Orange", "Apple", "Mango", "Apple", "Banana", "Apple"}

	start := 3
	present := LastIndexOf(fruits, "Apple", &start)

	fmt.Println("Present:", present) // 5
 ```

### ReduceRightStrict()
The LastIndexOf() function returns the last index (position) of a specified value. The LastIndexOf() function returns -1 if the value is not found. The LastIndexOf() starts at a specified index and searches from right to left (from the given position to the beginning of the slice). By default the search starts at the last element and ends at the first. Negative start values counts from the last element (but still searches from right to left).
 ```go
	fruits := []string{"Orange", "Apple", "Mango", "Apple", "Banana", "Apple"}

	start := 3
	present := LastIndexOf(fruits, "Apple", &start)

	fmt.Println("Present:", present) // 5
 ```

### Reduce()
The LastIndexOf() function returns the last index (position) of a specified value. The LastIndexOf() function returns -1 if the value is not found. The LastIndexOf() starts at a specified index and searches from right to left (from the given position to the beginning of the slice). By default the search starts at the last element and ends at the first. Negative start values counts from the last element (but still searches from right to left).
 ```go
	fruits := []string{"Orange", "Apple", "Mango", "Apple", "Banana", "Apple"}

	start := 3
	present := LastIndexOf(fruits, "Apple", &start)

	fmt.Println("Present:", present) // 5
 ```
