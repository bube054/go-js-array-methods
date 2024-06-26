# go-js-array-methods
The Array package offers a comprehensive suite of functions designed to manipulate slices in Go efficiently. Drawing inspiration from JavaScript array methods, our package implements familiar functionalities such as Push, Pop, and Reverse, all while ensuring that the original slice remains immutable. This design choices streamlines code complexity.

One notable feature of this package is its robust handling of indexes. Functions like At, Slice, and Includes seamlessly accommodate negative indexes, safeguarding against crashes while ensuring operations remain within range. Any attempt to access indexes beyond the boundaries of the slice results in clear error messages.

Harnessing the power of generics, our implementation minimizes redundancy and maximizes flexibility, contributing to a more versatile and adaptable codebase.

While some methods like Flat and FlatMap are not feasible due to language limitations, others, such as Sort, are deliberately omitted due to ambiguity, considering Go's built-in Sort function.

Furthermore, we introduce specialized variants of common functions like MapStrict, ReduceStrict, and ReduceRightStrict. These variants guarantee explicit return types ([]T), in contrast to their counterparts (Map, Reduce, ReduceRight), which allow for arbitrary return types ([]any).

Our Array package aims to empower Go developers with a robust toolkit for slice manipulation, prioritizing efficiency, clarity, and reliability.

| s/n | Function   | Description                                                    |
| --- | -------- | -------------------------------------------------------------- |
| 1   | At ✔     | The At() function returns an indexed element from a slice and returns a possible error relating to out of range indexes.     |
| 2   | Concat ✔ | The Concat() function Concatenates (joins) two or more slices. | The Concat() function returns a new slice, containing the joined slices. The Concat() function does not change the existing slices. |
| 3 | CopyWithin ✔ | The CopyWithin() function returns a copy of slice elements to another position in an slice and a possible error relating to out of range indexes.. The CopyWithin() function overwrites the existing values in the new slice. |
| 4 | Entries ✔ | The entries() function returns an Slice with structs with field/value: {index 0, element "Banana"}. The entries() function does not change the original slice. |
| 5 | Every ✔ | The Every() function executes a function for each slice element. The Every() function returns true if the function returns true for all elements. The Every() function returns false if the function returns false for one element. The Every() function does not execute the function for empty elements. |
| 6 | Fill ✔ | The Fill() function returns fills of specified elements in an slice with a value and a possible error due to indexes out of range.|
| 7 | Filter ✔ | The filter() function creates a new slice filled with elements that pass a test provided by a function. The filter() function does not execute the function for empty elements. The filter() function does not change the original slice. |
| 8 | Find ✔ | The Find() function returns the value of the first element that passes a test. The Find() function executes a function for each slice element. The Find() function returns undefined if no elements are found. The Find() function does not execute the function for empty elements. The Find() function does not change the original slice. |
| 9 | FindIndex ✔ |The FindIndex() function executes a function for each slice element. The FindIndex() function returns the index (position) of the first element that passes a test. The FindIndex() function returns -1 if no match is found. The FindIndex() function does not execute the function for empty slice elements. The FindIndex() function does not change the original slice. |
| 10 | FindLast ✔ | The FindLast() function returns the value of the last element that passes a test. The FindLast() function executes a function for each slice element. The FindLast() function returns -1 if no elements are found. The FindLast() function does not execute the function for empty elements. The FindLast() function does not change the original slice. |
| 11 | FindLastIndex ✔ | The FindLastIndex() function executes a function for each array element. The FindLastIndex() function returns the index (position) of the last element that passes a test. The FindLastIndex() function returns -1 if no match is found. The FindLastIndex() function does not execute the function for empty slice elements.The FindLastIndex() function does not change the original slice. |
| 12 | Flat ❌ | ————————— |
| 13 | FlatMap ❌ | ————————— |
| 14 | ForEach ✔ | The ForEach() function calls a function for each element in an slice. The ForEach() function is not executed for empty elements. |
| 15 | Includes ✔ | The Includes() function returns true if an slice contains a specified value. The Includes() function returns false if the value is not found. The Includes() function is case sensitive. |
| 16 | IndexOf ✔ | The IndexOf() function returns the first index (position) of a specified value. The IndexOf() function returns -1 if the value is not found. The IndexOf() function starts at a specified index and searches from left to right (from the given start postion to the end of the array). By default the search starts at the first element and ends at the last. Negative start values counts from the last element (but still searches from left to right). |
| 17 | Join ✔ | The Join() function returns an slice as a string. The Join() function does not change the original slice. Any separator can be specified. The default is comma (,). Join function only works for strings. |
| 18 | Keys ❌ | ————————— |
| 19 | LastIndexOf ✔ | The LastIndexOf() function returns the last index (position) of a specified value and a possible error due to indexes out of range. The LastIndexOf() function returns -1 if the value is not found. The LastIndexOf() starts at a specified index and searches from right to left (from the given position to the beginning of the slice). By default the search starts at the last element and ends at the first. Negative start values counts from the last element (but still searches from right to left). |
| 20 | Map ✔ | Map() creates a new slice from calling a function for every slice element. Map creates a new slice of any type. Map() does not execute the function for empty elements. Map() does not change the original slice. |
| 21 | MapStrict ✔ | Map() creates a new slice from calling a function for every slice element. Map creates a new slice of the original slice type. Map() does not execute the function for empty elements. Map() does not change the original slice. |
| 22 | Pop ✔ | The Pop() function removes (pops) the last element of an slice. The Pop() function does not change the original slice. The Pop() function returns the new slice without the last element and a pointer of removed element. |
| 23 | Push ✔ | The Push() function adds new items to the end of an slice. The Push() function returns the new slice. |
| 24 | Reduce ✔ | The Reduce() function executes a Reducer function for slice element. The Reduce() function returns the function's accumulated result and an an error. The Reduce() function does not execute the function for empty slice elements. The Reduce() function does not change the original slice. Note at the first callback, there is no return value from the previous callback. Normally, alice element 0 is used as initial value, and the iteration starts from slice element 1. If an initial value is supplied, this is used, and the iteration starts from slice element 0. |
| 25 | ReduceStrict ✔ | The ReduceStrict() function executes a Reducer function for slice element. The ReduceStrict() function returns the function's accumulated result(typed with the slices type) and an an error. The ReduceStrict() function does not execute the function for empty slice elements. The ReduceStrict() function does not change the original slice. Note at the first callback, there is no return value from the previous callback. Normally, alice element 0 is used as initial value, and the iteration starts from slice element 1. If an initial value is supplied, this is used, and the iteration starts from slice element 0. |
| 26 | ReduceRight ✔ | The Reduce() function executes a Reducer function for slice element. The ReduceRight() function works from right to left. The Reduce() function returns the function's accumulated result and an an error. The Reduce() function does not execute the function for empty slice elements. The Reduce() function does not change the original slice. Note at the first callback, there is no return value from the previous callback. Normally, alice element 0 is used as initial value, and the iteration starts from slice element 1. If an initial value is supplied, this is used, and the iteration starts from slice element 0. |
| 27 | ReduceRightStrict ✔ | The ReduceStrict() function executes a Reducer function for slice element. The ReduceRightStrict() function works from right to left . The ReduceStrict() function returns the function's accumulated result(typed with the slices type) and an an error. The ReduceStrict() function does not execute the function for empty slice elements. The ReduceStrict() function does not change the original slice. Note at the first callback, there is no return value from the previous callback. Normally, alice element 0 is used as initial value, and the iteration starts from slice element 1. If an initial value is supplied, this is used, and the iteration starts from slice element 0. |
| 28 | Reverse ✔ | The Reverse() function Reverses the order of the elements in an slice. The Reverse() function does not overwrites the original slice. |
| 29 | Shift ✔ | The Shift() function removes the first item of an slice. The Shift() function changes the original slice. The Shift() function returns the new slice and a pointer to the  shifted element. |
| 30 | Slice ✔ | The Slice() function returns selected elements in an slice, as a new slice and and error. The Slice() function selects from a given start, up to a (not inclusive) given end. The Slice() function does not change the original slice. |
| 31 | Some ✔ | The Some() function checks if any slice elements pass a test (provided as a callback function). The Some() function executes the callback function once for each slice element.The Some() function returns true (and stops) if the function returns true for one of the slice elements. The Some() function returns false if the function returns false for all of the slice elements. The Some() function does not execute the function for empty slice elements. The Some() function does not change the original slice.  |
| 32 | Sort ❌ | ————————— |
| 33 | Splice ✔ | The Splice() function adds and/or removes slice elements. The Splice() function does not overwrites the original slice. |
| 34 | ToString ✔ | The ToString() function returns a string with slice values separated by commas. The ToString() function does not change the original slice. The ToString function only works for a slice of strings. |
| 35 | Unshift ✔ | The Unshift() function adds new elements to the beginning of an slice. The Unshift() function does not overwrite the original slice. |
| 36 | ValueOf ✔ | The ValueOf() function returns the slice itself. The ValueOf() function does not change the original slice. fruits. ValueOf() returns the same as fruits. |
| 37 | With ✔ | The With() function updates a specified slice element. The With() function returns a new slice. The With() function does not change the original slice. |

## Downloading The Package
```
go get github.com/bube054/go-js-array-methods
```

## Importing And Using The Package
```go
package main

import (
  "github.com/bube054/go-js-array-methods/array"
)

func main() {
	isekais := []string{"re:zero", "mushoku tensei", "tbate"}
	res := array.Reverse(isekais)
	fmt.Println(res) // [tbate, mushoku tensei, re:zero]
}
```

### 1. At()
The At() function returns an indexed element from a slice and returns a possible error relating to out of range indexes.
 ```go
fruits := []string{"Banana", "Orange", "Apple", "Mango"}
index := 3 // negative values also allowed
result, err := array.At(fruits, index)

if err != nil {
	fmt.Println("Err", err) // nil
}

fmt.Println("Result:", result) // nil
 ```

### 2. Concat()
The Concat() function Concatenates (joins) two or more slices. | The Concat() function returns a new slice, containing the joined slices. The Concat() function does not change the existing slices.
 ```go
names1 := []string{"Cecilie", "Lone"}
names2 := []string{"Emil", "Tobias", "Linus"}
result := []string{"Cecilie", "Lone", "Emil", "Tobias", "Linus"}

children := array.Concat[string](names1, names2)

if !reflect.DeepEqual(result, children) {
	fmt.Println("err") // "err"
	return
}

fmt.Println(children) // ["Cecilie", "Lone", "Emil", "Tobias", "Linus"]
 ```

### 3. CopyWithin()
The CopyWithin() function returns a copy of slice elements to another position in an slice and a possible error relating to out of range indexes.. The CopyWithin() function overwrites the existing values in the new slice.
```go
fruits := []string{"Banana", "Orange", "Apple", "Mango", "Kiwi", "Papaya"}

newFruits, err := array.CopyWithin(fruits, -6, -4, 2)

fmt.Println("New:", newFruits) // [Banana Orange Apple Mango Kiwi Papaya]
fmt.Println("New:", err) // error occurs because of out of range index.
 ```


### 4. Entries()
The entries() function returns an Slice with structs with field/value: {index 0, element "Banana"}. The entries() function does not change the original slice.
 ```go
fruits := []string{"Banana", "Orange", "Apple", "Mango", "Kiwi", "Papaya"}

newFruits := array.Entries(fruits)

fmt.Println("New:", newFruits) // [{0 Banana} {1 Orange} {2 Apple} {3 Mango} {4 Kiwi} {5 Papaya}]
 ```

### 5. Every()
The Every() function executes a function for each slice element. The Every() function returns true if the function returns true for all elements. The Every() function returns false if the function returns false for one element. The Every() function does not execute the function for empty elements.
 ```go
points := []int{10, 25, 65, 40}

allPointsGreaterThan50 := func(point int, ind int, list []int) bool {
	if point > 50 {
		return true
	} else {
		return false
	}
}

result := array.Every(points, allPointsGreaterThan50)

fmt.Println("Result:", result) // false
 ```

### 6. Fill()
The Fill() function returns fills of specified elements in an slice with a value and a possible error due to indexes out of range.
 ```go
fruits := []string{"Banana", "Orange", "Apple", "Mango"}

newFruits, err := array.Fill(fruits, "Kiwi", 1, 3)

fmt.Println("newFruits:", newFruits) // [Banana Kiwi Kiwi Mango]
fmt.Println("err:", err) // an error can occur because of out of range index.
 ```

### 7. Filter()
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

newAges := array.Filter(ages, checkAdult)

fmt.Println("newAges:", newAges) // [32 33 40]
 ```

### 8. Find()
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

ans := array.Find(ages, find18YrsOld)

if ans != nil {
	fmt.Println("ans:", *ans) // 20
} else {
	fmt.Println("ans:", ans) // nil
}
 ```

### 9. FindIndex()
The FindIndex() function executes a function for each slice element. The FindIndex() function returns the index (position) of the first element that passes a test. The FindIndex() function returns -1 if no match is found. The FindIndex() function does not execute the function for empty slice elements. The FindIndex() function does not change the original slice.
 ```go
ages := []int{3, 10, 18, 20}

find18YrsOld := func(age int, ind int, list []int) bool {
	if age >= 18 {
		return true
	} else {
		return false
	}
}

ans := array.FindIndex(ages, find18YrsOld)

fmt.Println("ans:", ans) // 2
 ```

### 10. FindLast()
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

ans := array.FindLast(ages, find18YrsOld)
fmt.Println("ans:", *ans) // 25
 ```

### 11. FindLastIndex()
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

ans := array.FindIndex(ages, find18YrsOld)

fmt.Println("ans:", ans) // 3
 ```

### 12. Flat() ❌

### 13. FlatMap() ❌

### 14. ForEach()
The ForEach() function calls a function for each element in an slice. The ForEach() function is not executed for empty elements.
 ```go
numbers := []int{65, 44, 12, 4}

sum := 0
array.ForEach(numbers, func(el int, ind int, list []int) {
	sum += el
})

fmt.Println("Sum:", sum) // 125
 ```

### 15. Includes()
The Includes() function returns true if an slice contains a specified value. The Includes() function returns false if the value is not found. The Includes() function is case sensitive.
 ```go
fruits := []string{"Banana", "Orange", "Apple", "Mango"}

start := 1 // start searching from index 1, negative indexes are also allowed
present := array.Includes(fruits, "Orange", &start)

fmt.Println("Present:", present) // true
 ```

### 16. IndexOf()
The IndexOf() function returns the first index (position) of a specified value. The IndexOf() function returns -1 if the value is not found. The IndexOf() function starts at a specified index and searches from left to right (from the given start postion to the end of the array). By default the search starts at the first element and ends at the last. Negative start values counts from the last element (but still searches from left to right).
 ```go
  // example 1
fruits := []string{"Banana", "Orange", "Apple", "Mango"}

// start := 1
present := array.IndexOf(fruits, "Orange", nil)

fmt.Println("Present:", present) // 1

// example 2 
fruits = []string{"Orange", "Banana", "Orange", "Apple", "Mango"}

start = 1
present = array.IndexOf(fruits, "Orange", &start)

fmt.Println("Present:", present)
 ```

### 17. Join()
The Join() function returns an slice as a string. The Join() function does not change the original slice. Any separator can be specified. The default is comma (,). Join function only works for strings.
 ```go
// example 1
fruits := []string{"Banana", "Orange", "Apple", "Mango"}

delim := " and "
newFruits := array.Join(fruits, &delim)
fmt.Println("newFruits:", newFruits) // Banana and Orange and Apple and Mango
 ```

### 18. Keys() ❌

### 19. LastIndexOf()
The LastIndexOf() function returns the last index (position) of a specified value and a possible error due to indexes out of range. The LastIndexOf() function returns -1 if the value is not found. The LastIndexOf() starts at a specified index and searches from right to left (from the given position to the beginning of the slice). By default the search starts at the last element and ends at the first. Negative start values counts from the last element (but still searches from right to left).
 ```go
fruits := []string{"Orange", "Apple", "Mango", "Apple", "Banana", "Apple"}

start := 3
present := array.LastIndexOf(fruits, "Apple", &start)

fmt.Println("Present:", present) // 5
 ```

### 20. Map()
Map() creates a new slice from calling a function for every slice element. Map creates a new slice of any type. Map() does not execute the function for empty elements. Map() does not change the original slice.
 ```go
numbers := []int{65, 44, 12, 4}

newNums := array.Map(numbers, func(e int, i int, s []int) any {
	return e * 10
})

fmt.Println("newNums:", newNums) // [650 440 120 40]
 ```

### 21. MapSrict()
Map() creates a new slice from calling a function for every slice element. Map creates a new slice of the original slice type. Map() does not execute the function for empty elements. Map() does not change the original slice.
 ```go
numbers := []int{65, 44, 12, 4}

newNums2 := array.MapStrict(numbers, func(e int, i int, s []int) int {
	return e / 10
})

fmt.Println("newNums2:", newNums2) // [6 4 1 0]
 ```

### 22. Pop()
The Pop() function removes (pops) the last element of an slice. The Pop() function does not change the original slice. The Pop() function returns the new slice without the last element and a pointer of removed element.
 ```go
fruits := []string{"Banana", "Orange", "Apple", "Mango"}

newFruits, fruit := array.Pop(fruits)

fmt.Println("newFruits:", newFruits) // [Banana Orange Apple]
fmt.Println("fruit:", *fruit) // Mango
 ```

### 23. Push()
The Push() function adds new items to the end of an slice. The Push() function returns the new slice.
 ```go
fruits := []string{"Banana", "Orange", "Apple", "Mango"}

newFruits := array.Push(fruits, "Kiwi", "Lemon")

fmt.Println("newFruits:", newFruits) // [Banana Orange Apple Mango Kiwi Lemon]
 ```

### 24. Reduce()
The Reduce() function executes a Reducer function for slice element. The Reduce() function returns the function's accumulated result and an an error. The Reduce() function does not execute the function for empty slice elements. The Reduce() function does not change the original slice. Note at the first callback, there is no return value from the previous callback. Normally, alice element 0 is used as initial value, and the iteration starts from slice element 1. If an initial value is supplied, this is used, and the iteration starts from slice element 0.
 ```go
myFunc := func(acc any, el int, ind int, slice []int) any {
	accNum := acc.(int)

	return accNum - el
}

init := 0
res, err := array.Reduce(numbers, myFunc, init)

fmt.Println("RES:", res) // 250
fmt.Println("ERR:", err) // nil
 ```

### 25. ReduceStrict()
The ReduceStrict() function executes a Reducer function for slice element. The ReduceStrict() function returns the function's accumulated result(typed with the slices type) and an an error. The ReduceStrict() function does not execute the function for empty slice elements. The ReduceStrict() function does not change the original slice. Note at the first callback, there is no return value from the previous callback. Normally, alice element 0 is used as initial value, and the iteration starts from slice element 1. If an initial value is supplied, this is used, and the iteration starts from slice element 0.
 ```go
numbers := []int{175, 50, 25}

myFunc := func(acc int, el int, ind int, slice []int) int {
	return acc - el
}

// init := 0
res, err := array.ReduceStrict(numbers, myFunc, nil)

fmt.Println("RES:", res) // 100
fmt.Println("ERR:", err) // nil
 ```

### 26. ReduceRight()
The Reduce() function executes a Reducer function for slice element. The ReduceRight() function works from right to left. The Reduce() function returns the function's accumulated result and an an error. The Reduce() function does not execute the function for empty slice elements. The Reduce() function does not change the original slice. Note at the first callback, there is no return value from the previous callback. Normally, alice element 0 is used as initial value, and the iteration starts from slice element 1. If an initial value is supplied, this is used, and the iteration starts from slice element 0.
 ```go
numbers := []int{100}

myFunc := func(acc any, el int, ind int, slice []int) any {
	accNum := acc.(int)

	return accNum - el
}

// init := 0
res, err := array.ReduceRight(numbers, myFunc, 300)

fmt.Println("RES:", res) // 200
fmt.Println("ERR:", err) // nil
 ```

### 27. ReduceRightStrict()
The ReduceStrict() function executes a Reducer function for slice element. The ReduceRightStrict() function works from right to left . The ReduceStrict() function returns the function's accumulated result(typed with the slices type) and an an error. The ReduceStrict() function does not execute the function for empty slice elements. The ReduceStrict() function does not change the original slice. Note at the first callback, there is no return value from the previous callback. Normally, alice element 0 is used as initial value, and the iteration starts from slice element 1. If an initial value is supplied, this is used, and the iteration starts from slice element 0.
 ```go
numbers := []int{175, 50, 25}

myFunc := func(acc int, el int, ind int, slice []int) int {
	return acc - el
}

// init := 0
res, err := array.ReduceRightStrict(numbers, myFunc, nil)

fmt.Println("RES:", res) // -200
fmt.Println("ERR:", err) // nil
 ```

### 28. Reverse()
The Reverse() function Reverses the order of the elements in an slice. The Reverse() function does not overwrites the original slice.
 ```go
fruits := []string{"Cecilie", "Lone", "Emil", "Tobias", "Linus"}

res := array.Reverse(fruits)
fmt.Println("RES:", res) // [Linus Tobias Emil Lone Cecilie]
 ```

###  29. Shift()
The Shift() function removes the first item of an slice. The Shift() function changes the original slice. The Shift() function returns the new slice and a pointer to the  shifted element.
 ```go
fruits := []string{"Banana", "Orange", "Apple", "Mango"}

newFruits, fruit := array.Shift(fruits)

fmt.Println("newFruits:", newFruits) // [Orange Apple Mango]
fmt.Println("fruit:", *fruit) // Banana
 ```

### 30. Slice()
The Slice() function returns selected elements in an slice, as a new slice and and error. The Slice() function selects from a given start, up to a (not inclusive) given end. The Slice() function does not change the original slice.
 ```go
fruits := []string{"Banana", "Orange", "Lemon", "Apple", "Mango"}
newFruits, err := array.Slice(fruits, -3, -1)

fmt.Println("RES:", newFruits) // [Lemon Apple]
fmt.Println("ERR:", err) // nil
 ```

### 31. Sort() ❌

### 32. Some()
The Some() function checks if any slice elements pass a test (provided as a callback function). The Some() function executes the callback function once for each slice element.The Some() function returns true (and stops) if the function returns true for one of the slice elements. The Some() function returns false if the function returns false for all of the slice elements. The Some() function does not execute the function for empty slice elements. The Some() function does not change the original slice.
 ```go
ages := []int{3, 10, 18, 20}

age := array.Some(ages, func(age int, ind int, list []int) bool {
	return age >  100
})

fmt.Println("RES:", age) // false
 ```

### 33. Splice()
The Splice() function adds and/or removes slice elements. The Splice() function does not overwrites the original slice.
 ```go
fruits := []string{"Banana", "Orange", "Apple", "Mango", "Kiwi"}

index := 2
res, err := array.Splice(fruits, 2, &index)
fmt.Println("RES:", res) // [Banana Orange Mango Kiwi]
fmt.Println("ERR:", err) // nil

index = 2
res, err = array.Splice(fruits, 0, &index, "Mango", "Coconut")
fmt.Println("RES:", res) // [Mango Coconut Apple Mango Kiwi]
fmt.Println("ERR:", err) // nil
 ```

### 34. ToString()
The ToString() function returns a string with slice values separated by commas. The ToString() function does not change the original slice. The ToString function only works for a slice of strings.
 ```go
fruits := []string{"Banana", "Orange", "Apple", "Mango"}

res := array.ToString(fruits)
fmt.Println("RES:", res) // Banana,Orange,Apple,Mango
 ```

### 35. Unshift()
The Unshift() function adds new elements to the beginning of an slice. The Unshift() function does not overwrite the original slice.
 ```go
fruits := []string{"Banana", "Orange", "Apple", "Mango"}

newFruits := array.UnShift(fruits, "Kiwi", "Lemon")

fmt.Println("newFruits:", newFruits) // [Kiwi Lemon Banana Orange Apple Mango]
 ```

### 36. ValueOf()
The Unshift() function adds new elements to the beginning of an slice. The Unshift() function does not overwrite the original slice.
 ```go
fruits := []string{"Banana", "Orange", "Apple", "Mango"}

newFruits := array.ValueOf(fruits)

fmt.Println("newFruits", newFruits) // [Banana Orange Apple Mango]
 ```

### 37. With()
The With() function updates a specified slice element. The With() function returns a new slice. The With() function does not change the original slice.
 ```go
months := []string{"Januar", "Februar", "Mar", "April"}
// _ = months
myMonths, err := array.With(months, 2, "March")

fmt.Println("Mons:", myMonths) // [Januar Februar March April]
fmt.Println("Err:", err) // nil
 ```