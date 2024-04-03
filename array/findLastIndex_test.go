package array

import (
	"fmt"
	"testing"
)

func TestFindLastIndex(t *testing.T) {
	ages := []int{3, 10, 18, 20}

	find18YrsOld := func(el int, ind int, list []int) bool {
		if el > 18 {
			return true
		} else {
			return false
		}
	}

	ans := FindIndex(ages, find18YrsOld)

	fmt.Println("ans:", ans)
}
