package array

import (
	"fmt"
	"testing"
)

func TestWith(t *testing.T) {
	months := []string{"Januar", "Februar", "Mar", "April"}
	// _ = months
	myMonths, err := With(months, 2, "March")

	fmt.Println("Mons:", myMonths)
	fmt.Println("Err:", err)
}
