package main

import "testing"

func TestMaxIndexOfMaxElement(t *testing.T) {
	tests := []struct {
		name           string
		numbers        []int
		indexWant     int
		MaxElementWant int
	}{
		{"All of numbers is positive", []int{1, 4, 25, 25, 100}, 4, 100},
		{"All of numbers is nrgative", []int{-1, -40, -10, -1, -5}, 0, -1},
		{"Numbers has pos and neg", []int{-1, 5, -20, 2}, 1, 5},
	}

	for _, test := range tests {
		 index, element := MaxIndexOfMaxElement(test.numbers)
		if index != test.indexWant || element != test.MaxElementWant{
			t.Errorf("IndexOfMaxAndMaxElement test %s gotIndex %d IndexWant %d gotElement %d ElementWant %d",
				test.name, index, test.indexWant, element, test.MaxElementWant)
	}
}
}