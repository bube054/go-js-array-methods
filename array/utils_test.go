package array

import (
	"fmt"
	"testing"
)

func TestUtils(t *testing.T) {
	fruits := []string{"Banana", "Orange", "Apple", "Mango"}
	index := -4

	newIndex, err := ConvertIndex(fruits, index, "index")

	fmt.Println(newIndex, err)
}