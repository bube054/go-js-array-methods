package array

import (
	"fmt"
	"testing"
)

func TestFilter(t *testing.T) {
	ages := []int{32, 33, 16, 40}

	checkAdult := func(el int, ind int, list []int) bool {
		if el > 18 {
			return true
		} else {
			return false
		}
	}

	newAges := Filter(ages, checkAdult)

	fmt.Println("newAges:", newAges)
}
