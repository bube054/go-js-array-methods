package array

import (
	"fmt"
	"testing"
)

func TestEvery(t *testing.T) {
	points := []int{10, 25, 65, 40}

	allIntsGreaterThan100 := func(point int, ind int, list []int) bool {
		if point > 1 {
			return true
		} else {
			return false
		}
	}

	result := Every(points, allIntsGreaterThan100)

	fmt.Println("Result:", result)
}
