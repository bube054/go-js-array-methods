package main

import (
	"fmt"
	"github.com/bube054/go-js-array-methods/array"
)

func main() {
	// a := []int{1,2,3}
	// b := []int{4,5,6}
	// c := []int{7,8,9}
	d := []string{"ant", "bison", "camel", "duck", "elephant"}

	r := array.Slice[string](d, 7, 2)

	fmt.Println(r)

	// r := array.Concat[int](a, b, c)

	// fmt.Println(a)
	// array.ConcatMut[int](&a, b, c)
	// x, err := array.ShiftMut[int](&a, 0)
	// x,y,err := array.Shift[int](a, 0)
	// x, err := array.PopMut[int](&a, 0)
	// x, y, err := array.Pop[int](a, 0)
	// array.PushMut[int](&a, 4)
	// fmt.Println(r)
	// fmt.Println(a)
}

// [1, 2, 3].push(4); // [1, 2, 3, 4]
// [1, 2, 3].pop(); // [1, 2]
// [1, 2, 3].shift(); // [2, 3]
// [1, 2, 3].unshift(0); // [0, 1, 2, 3]
// ['a', 'b'].concat('c'); // ['a', 'b', 'c']
// ['a', 'b', 'c'].join('-'); // a-b-c
// ['a', 'b', 'c'].slice(1); // ['b', 'c']
// ['a', 'b', 'c'].indexOf('b'); // 1
// ['a', 'b', 'c'].includes('c'); // true
// [3, 5, 6, 8].find((n) => n % 2 === 0); // 6
// [2, 4, 3, 5].findIndex((n) => n % 2 !== 0); // 2
// [3, 4, 8, 6].map((n) => n * 2); // [6, 8, 16, 12]
// [1, 4, 7, 8].filter((n) => n % 2 === 0); // [4, 8]
// [2, 4, 3, 7].reduce((acc, cur) => acc + cur); // 16
// [2, 3, 4, 5].every((x) => x < 6); // true
// [3, 5, 6, 8].some((n) => n > 6); // true
// [1, 2, 3, 4].reverse(); // [4, 3, 2, 1]
// [3, 5, 7, 8].at(-2); // 7
