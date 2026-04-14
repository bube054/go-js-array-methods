package array

import (
	"reflect"
	"testing"
)

// TestArrayConcat tests the Array.Concat receiver method
func TestArrayConcat(t *testing.T) {
	arr := Array[int]{1, 2, 3}
	slice2 := Array[int]{4, 5, 6}
	slice3 := Array[int]{7, 8}

	result := arr.Concat(slice2, slice3)
	expected := Array[int]{1, 2, 3, 4, 5, 6, 7, 8}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Array.Concat() = %v, expected %v", result, expected)
	}
}

// TestArrayCopyWithin tests the Array.CopyWithin receiver method
func TestArrayCopyWithin(t *testing.T) {
	arr := Array[int]{1, 2, 3, 4, 5}
	result, err := arr.CopyWithin(0, 3, 5)

	if err != nil {
		t.Errorf("Array.CopyWithin() got unexpected error: %v", err)
	}

	expected := Array[int]{4, 5, 3, 4, 5}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Array.CopyWithin() = %v, expected %v", result, expected)
	}
}

// TestArrayEvery tests the Array.Every receiver method
func TestArrayEvery(t *testing.T) {
	tests := []struct {
		name      string
		arr       Array[int]
		predicate func(int, int, []int) bool
		expected  bool
	}{
		{
			name:      "Array.Every all pass",
			arr:       Array[int]{2, 4, 6, 8},
			predicate: func(el, _ int, _ []int) bool { return el%2 == 0 },
			expected:  true,
		},
		{
			name:      "Array.Every not all pass",
			arr:       Array[int]{2, 4, 5, 8},
			predicate: func(el, _ int, _ []int) bool { return el%2 == 0 },
			expected:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.arr.Every(tt.predicate)
			if result != tt.expected {
				t.Errorf("Array.Every() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

// TestArrayFill tests the Array.Fill receiver method
func TestArrayFill(t *testing.T) {
	arr := Array[int]{1, 2, 3, 4, 5}
	result, err := arr.Fill(0, 1, 3)

	if err != nil {
		t.Errorf("Array.Fill() got unexpected error: %v", err)
	}

	expected := Array[int]{1, 0, 0, 4, 5}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Array.Fill() = %v, expected %v", result, expected)
	}
}

// TestArrayFilter tests the Array.Filter receiver method
func TestArrayFilter(t *testing.T) {
	arr := Array[int]{10, 20, 25, 30}
	predicate := func(el, _ int, _ []int) bool { return el > 18 }

	result := arr.Filter(predicate)
	expected := Array[int]{20, 25, 30}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Array.Filter() = %v, expected %v", result, expected)
	}
}

// TestArrayFind tests the Array.Find receiver method
func TestArrayFind(t *testing.T) {
	arr := Array[int]{10, 20, 30, 40}
	predicate := func(el, _ int, _ []int) bool { return el > 25 }

	result := arr.Find(predicate)
	expected := 30

	if result == nil || *result != expected {
		t.Errorf("Array.Find() = %v, expected %d", result, expected)
	}
}

// TestArrayFindIndex tests the Array.FindIndex receiver method
func TestArrayFindIndex(t *testing.T) {
	arr := Array[int]{10, 20, 30, 40}
	predicate := func(el, _ int, _ []int) bool { return el > 25 }

	result := arr.FindIndex(predicate)
	expected := 2

	if result != expected {
		t.Errorf("Array.FindIndex() = %d, expected %d", result, expected)
	}
}

// TestArrayForEach tests the Array.ForEach receiver method
func TestArrayForEach(t *testing.T) {
	arr := Array[int]{1, 2, 3}
	var results []int

	arr.ForEach(func(el, _ int, _ []int) {
		results = append(results, el*2)
	})

	expected := []int{2, 4, 6}
	if !reflect.DeepEqual(results, expected) {
		t.Errorf("Array.ForEach() = %v, expected %v", results, expected)
	}
}

// TestArrayIncludes tests the Array.Includes receiver method
func TestArrayIncludes(t *testing.T) {
	arr := Array[string]{"apple", "banana", "cherry"}

	result := arr.Includes("banana", nil)
	if !result {
		t.Errorf("Array.Includes() = %v, expected true", result)
	}

	result = arr.Includes("grape", nil)
	if result {
		t.Errorf("Array.Includes() = %v, expected false", result)
	}
}

// TestArrayIndexOf tests the Array.IndexOf receiver method
func TestArrayIndexOf(t *testing.T) {
	arr := Array[string]{"apple", "banana", "cherry", "banana"}

	result := arr.IndexOf("banana", nil)
	expected := 1

	if result != expected {
		t.Errorf("Array.IndexOf() = %d, expected %d", result, expected)
	}

	result = arr.IndexOf("grape", nil)
	if result != -1 {
		t.Errorf("Array.IndexOf() = %d, expected -1", result)
	}
}

// TestArrayLastIndexOf tests the Array.LastIndexOf receiver method
func TestArrayLastIndexOf(t *testing.T) {
	arr := Array[string]{"apple", "banana", "cherry", "banana"}

	result := arr.LastIndexOf("banana", nil)
	expected := 3

	if result != expected {
		t.Errorf("Array.LastIndexOf() = %d, expected %d", result, expected)
	}
}

// TestArrayMap tests the Array.Map receiver method
func TestArrayMap(t *testing.T) {
	arr := Array[int]{1, 2, 3}

	result := arr.Map(func(el, _ int, _ []int) any {
		return el * 10
	})

	expected := Array[any]{10, 20, 30}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Array.Map() = %v, expected %v", result, expected)
	}
}

// TestArrayMapStrict tests the Array.MapStrict receiver method
func TestArrayMapStrict(t *testing.T) {
	arr := Array[int]{1, 2, 3}

	result := arr.MapStrict(func(el, _ int, _ []int) int {
		return el * 10
	})

	expected := Array[int]{10, 20, 30}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Array.MapStrict() = %v, expected %v", result, expected)
	}
}

// TestArrayPop tests the Array.Pop receiver method
func TestArrayPop(t *testing.T) {
	arr := Array[int]{1, 2, 3, 4, 5}
	result, popped := arr.Pop()

	expectedSlice := Array[int]{1, 2, 3, 4}
	expectedPopped := 5

	if !reflect.DeepEqual(result, expectedSlice) {
		t.Errorf("Array.Pop() slice = %v, expected %v", result, expectedSlice)
	}

	if popped == nil || *popped != expectedPopped {
		t.Errorf("Array.Pop() popped = %v, expected %d", popped, expectedPopped)
	}
}

// TestArrayPush tests the Array.Push receiver method
func TestArrayPush(t *testing.T) {
	arr := Array[int]{1, 2, 3}
	result := arr.Push(4, 5)

	expected := Array[int]{1, 2, 3, 4, 5}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Array.Push() = %v, expected %v", result, expected)
	}
}

// TestArrayReduce tests the Array.Reduce receiver method
func TestArrayReduce(t *testing.T) {
	arr := Array[int]{1, 2, 3, 4}

	result, err := arr.Reduce(func(acc any, el int, _ int, _ []int) any {
		accNum := acc.(int)
		return accNum + el
	}, 0)

	if err != nil {
		t.Errorf("Array.Reduce() got unexpected error: %v", err)
	}

	expected := 10
	if result.(int) != expected {
		t.Errorf("Array.Reduce() = %v, expected %d", result, expected)
	}
}

// TestArrayReduceStrict tests the Array.ReduceStrict receiver method
func TestArrayReduceStrict(t *testing.T) {
	arr := Array[int]{1, 2, 3, 4}
	initialVal := 0

	result, err := arr.ReduceStrict(func(acc int, el int, _ int, _ []int) int {
		return acc + el
	}, &initialVal)

	if err != nil {
		t.Errorf("Array.ReduceStrict() got unexpected error: %v", err)
	}

	expected := 10
	if result != expected {
		t.Errorf("Array.ReduceStrict() = %d, expected %d", result, expected)
	}
}

// TestArrayReverse tests the Array.Reverse receiver method
func TestArrayReverse(t *testing.T) {
	arr := Array[int]{1, 2, 3, 4, 5}
	result := arr.Reverse()

	expected := Array[int]{5, 4, 3, 2, 1}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Array.Reverse() = %v, expected %v", result, expected)
	}
}

// TestArrayShift tests the Array.Shift receiver method
func TestArrayShift(t *testing.T) {
	arr := Array[int]{1, 2, 3, 4, 5}
	result, shifted := arr.Shift()

	expectedSlice := Array[int]{2, 3, 4, 5}
	expectedShifted := 1

	if !reflect.DeepEqual(result, expectedSlice) {
		t.Errorf("Array.Shift() slice = %v, expected %v", result, expectedSlice)
	}

	if shifted == nil || *shifted != expectedShifted {
		t.Errorf("Array.Shift() shifted = %v, expected %d", shifted, expectedShifted)
	}
}

// TestArraySlice tests the Array.Slice receiver method
func TestArraySlice(t *testing.T) {
	arr := Array[int]{1, 2, 3, 4, 5}
	result, err := arr.Slice(1, 4)

	if err != nil {
		t.Errorf("Array.Slice() got unexpected error: %v", err)
	}

	expected := Array[int]{2, 3, 4}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Array.Slice() = %v, expected %v", result, expected)
	}
}

// TestArraySome tests the Array.Some receiver method
func TestArraySome(t *testing.T) {
	arr := Array[int]{1, 2, 3, 4, 5}
	predicate := func(el, _ int, _ []int) bool { return el > 3 }

	result := arr.Some(predicate)
	if !result {
		t.Errorf("Array.Some() = %v, expected true", result)
	}
}

// TestArraySplice tests the Array.Splice receiver method
func TestArraySplice(t *testing.T) {
	arr := Array[int]{1, 2, 3, 4, 5}
	howMany := 2
	result, err := arr.Splice(1, &howMany, 10, 11)

	if err != nil {
		t.Errorf("Array.Splice() got unexpected error: %v", err)
	}

	expectedResult := Array[int]{1, 10, 11, 4, 5}

	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("Array.Splice() result = %v, expected %v", result, expectedResult)
	}
}

// TestArrayUnShift tests the Array.UnShift receiver method
func TestArrayUnShift(t *testing.T) {
	arr := Array[int]{3, 4, 5}
	result := arr.UnShift(1, 2)

	expected := Array[int]{1, 2, 3, 4, 5}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Array.UnShift() = %v, expected %v", result, expected)
	}
}

// TestArrayValueOf tests the Array.ValueOf receiver method
func TestArrayValueOf(t *testing.T) {
	arr := Array[int]{1, 2, 3}
	result := arr.ValueOf()

	expected := Array[int]{1, 2, 3}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Array.ValueOf() = %v, expected %v", result, expected)
	}
}

// TestArrayWith tests the Array.With receiver method
func TestArrayWith(t *testing.T) {
	arr := Array[int]{1, 2, 3, 4, 5}
	result, err := arr.With(2, 99)

	if err != nil {
		t.Errorf("Array.With() got unexpected error: %v", err)
	}

	expected := Array[int]{1, 2, 99, 4, 5}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Array.With() = %v, expected %v", result, expected)
	}
}

// TestArrayMethodChaining verifies that methods returning Array[T] can be chained.
func TestArrayMethodChaining(t *testing.T) {
	arr := Array[int]{1, 2, 3, 4, 5}

	chain1 := arr.Filter(func(el, _ int, _ []int) bool {
		return el%2 == 0
	}).Push(6).Reverse()

	expected1 := Array[int]{6, 4, 2}
	if !reflect.DeepEqual(chain1, expected1) {
		t.Errorf("chained Filter->Push->Reverse = %v, expected %v", chain1, expected1)
	}

	chain2 := arr.Concat(Array[int]{6, 7}).Reverse()
	expected2 := Array[int]{7, 6, 5, 4, 3, 2, 1}
	if !reflect.DeepEqual(chain2, expected2) {
		t.Errorf("chained Concat->Reverse = %v, expected %v", chain2, expected2)
	}
}

// TestArrayEntries tests the Array.Entries receiver method
func TestArrayEntries(t *testing.T) {
	arr := Array[string]{"a", "b", "c"}
	result := arr.Entries()

	if len(result) != 3 {
		t.Errorf("Array.Entries() length = %d, expected 3", len(result))
	}

	if result[0].index != 0 || result[0].element != "a" {
		t.Errorf("Array.Entries()[0] = {%d, %s}, expected {0, a}", result[0].index, result[0].element)
	}

	if result[2].index != 2 || result[2].element != "c" {
		t.Errorf("Array.Entries()[2] = {%d, %s}, expected {2, c}", result[2].index, result[2].element)
	}
}

// TestArrayFindLast tests the Array.FindLast receiver method
func TestArrayFindLast(t *testing.T) {
	arr := Array[int]{10, 20, 30, 40, 20}
	predicate := func(el, _ int, _ []int) bool { return el == 20 }

	result := arr.FindLast(predicate)
	expected := 20

	if result == nil || *result != expected {
		t.Errorf("Array.FindLast() = %v, expected %d", result, expected)
	}
}

// TestArrayFindLastIndex tests the Array.FindLastIndex receiver method
func TestArrayFindLastIndex(t *testing.T) {
	arr := Array[int]{10, 20, 30, 40, 20}
	predicate := func(el, _ int, _ []int) bool { return el == 20 }

	result := arr.FindLastIndex(predicate)
	expected := 4

	if result != expected {
		t.Errorf("Array.FindLastIndex() = %d, expected %d", result, expected)
	}
}

// TestArrayJoin tests the Array.Join standalone function with string slices
func TestArrayJoin(t *testing.T) {
	tests := []struct {
		name      string
		slice     Array[string]
		separator string
		expected  string
	}{
		{
			name:      "Join with custom separator",
			slice:     Array[string]{"apple", "banana", "cherry"},
			separator: " - ",
			expected:  "apple - banana - cherry",
		},
		{
			name:      "Join with comma separator",
			slice:     Array[string]{"apple", "banana", "cherry"},
			separator: ",",
			expected:  "apple,banana,cherry",
		},
		{
			name:      "Join with empty separator",
			slice:     Array[string]{"a", "b", "c"},
			separator: "",
			expected:  "abc",
		},
		{
			name:      "Join with single element",
			slice:     Array[string]{"apple"},
			separator: ",",
			expected:  "apple",
		},
		{
			name:      "Join with empty array",
			slice:     Array[string]{},
			separator: ",",
			expected:  "",
		},
		{
			name:      "Join with space separator",
			slice:     Array[string]{"hello", "world", "go"},
			separator: " ",
			expected:  "hello world go",
		},
		{
			name:      "Join with pipe separator",
			slice:     Array[string]{"red", "green", "blue"},
			separator: "|",
			expected:  "red|green|blue",
		},
		{
			name:      "Join with special characters in separator",
			slice:     Array[string]{"one", "two", "three"},
			separator: ":::",
			expected:  "one:::two:::three",
		},
		{
			name:      "Join with numbers as strings",
			slice:     Array[string]{"1", "2", "3", "4", "5"},
			separator: "-",
			expected:  "1-2-3-4-5",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sep := tt.separator
			result := tt.slice.Join(sep)
			if result != tt.expected {
				t.Errorf("Join() = %s, expected %s", result, tt.expected)
			}
		})
	}
}

// TestArrayToString tests the Array.ToString receiver method
func TestArrayToString(t *testing.T) {
	tests := []struct {
		name     string
		arr      Array[string]
		expected string
	}{
		{
			name:     "ToString with multiple elements",
			arr:      Array[string]{"Banana", "Orange", "Apple", "Mango"},
			expected: "Banana,Orange,Apple,Mango",
		},
		{
			name:     "ToString with empty array",
			arr:      Array[string]{},
			expected: "",
		},
		{
			name:     "ToString with single element",
			arr:      Array[string]{"Apple"},
			expected: "Apple",
		},
		{
			name:     "ToString with numbers as strings",
			arr:      Array[string]{"1", "2", "3", "4", "5"},
			expected: "1,2,3,4,5",
		},
		{
			name:     "ToString with special characters",
			arr:      Array[string]{"hello", "world", "go!"},
			expected: "hello,world,go!",
		},
		{
			name:     "ToString with spaces",
			arr:      Array[string]{"hello world", "foo bar"},
			expected: "hello world,foo bar",
		},
		{
			name:     "ToString with empty strings",
			arr:      Array[string]{"", "hello", "", "world", ""},
			expected: ",hello,,world,",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.arr.ToString()
			if result != tt.expected {
				t.Errorf("Array.ToString() = %s, expected %s", result, tt.expected)
			}
		})
	}
}
