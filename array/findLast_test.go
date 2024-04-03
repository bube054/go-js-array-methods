package array

import (
	"fmt"
	"testing"
)

func TestFindLast(t *testing.T) {
	ages := []int{3, 7, 10, 15, 18, 20}

	find18YrsOld := func(el int, ind int, list []int) bool {
		if el > 9 {
			return true
		} else {
			return false
		}
	}

	// _, _ = ages, find18YrsOld

	ans := FindLast(ages, find18YrsOld)
	fmt.Println("ans:", *ans)
}
