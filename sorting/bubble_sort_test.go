package main

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{input: []int{5, 3, 8, 4, 2}, expected: []int{2, 3, 4, 5, 8}},
		{input: []int{1, 2, 3, 4, 5}, expected: []int{1, 2, 3, 4, 5}},
		{input: []int{5, 4, 3, 2, 1}, expected: []int{1, 2, 3, 4, 5}},
		{input: []int{3, 3, 3, 3, 3}, expected: []int{3, 3, 3, 3, 3}},
		{input: []int{}, expected: []int{}},
	}

	for _, test := range tests {
		result := BubbleSort(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("BubbleSort(%v) = %v; want %v", test.input, result, test.expected)
		}
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	input := []int{5, 3, 8, 4, 2, 7, 6, 1, 9, 0}
	for i := 0; i < b.N; i++ {
		BubbleSort(input)
	}
}
