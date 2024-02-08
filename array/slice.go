package array

import (
	_"fmt"
	"math"
)

func Slice[T comparable](s []T, begin, end int) []T{
	newBegin, newEnd := begin, end
	length := len(s)

	if (newBegin >= 0 && newEnd >= 0){
		if(newBegin < length && newEnd < length){
			
		}
	}

	return s
}

// e.g d := []string{"ant", "bison", "camel", "duck", "elephant"}
// 1) begin and end are both pos
// i) both in range 1,2 [bison]
// ii) first out of range 7,2 [camel duck elephant]
// iii) second out of range 2,7 [camel duck elephant]

// 2) begin is pos and end is neg
// i) both in range 1,-3 [bison]
// ii) first out of range 7,-2 [duck elephant]
// iii) second out of range 3,-7 [camel duck elephant]

// 3) begin is neg and end is pos
// i) both in range 1,-3 [bison]
// ii) first out of range 7,-2 [duck elephant]
// iii) second out of range 3,-7 [camel duck elephant]

// 2) begin and end are neg
// i) both in range 1,-3 [bison]
// ii) first out of range 7,-2 [duck elephant]
// iii) second out of range 3,-7 [camel duck elephant]