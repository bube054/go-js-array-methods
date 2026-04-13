package array

import (
	"fmt"
	"testing"
)

func TestAt(t *testing.T) {
	fruits := []string{"Banana", "Orange", "Apple", "Mango"}
	index := -2
	result, err := At(fruits, index)

	if err != nil {
		fmt.Println("Err", err)
	}

	fmt.Println("Result:", result)
}

func TestArrayAt(t *testing.T) {
	fruits := Array[string]{"Banana", "Orange", "Apple", "Mango"}
	index := -2
	result, err := fruits.At(index)
	if err != nil{
		t.Fatal("Error", err)
	}
	if result != "Apple" {
		t.Fatalf("Expected 'Apple', got '%s'", result)
	}
}