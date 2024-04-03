package array

import (
	"fmt"
	"reflect"
	"testing"
)

func TestConcat(t *testing.T) {
	names1 := []string{"Cecilie", "Lone"}
	names2 := []string{"Emil", "Tobias", "Linus"}
	result := []string{"Cecilie", "Lone", "Emil", "Tobias", "Linus"}

	children := Concat[string](names1, names2)

	if !reflect.DeepEqual(result, children) {
		t.Fatalf("results and children are not equal.")
	}

	fmt.Println(children)
}
