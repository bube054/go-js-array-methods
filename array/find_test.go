package array

import (
	"fmt"
	"testing"
)

func TestFind(t *testing.T) {
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
		fmt.Println("ans:", *ans)
	} else {
		fmt.Println("ans:", ans)
	}
}
