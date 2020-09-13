package main

import (
	"reflect"
	"testing"
)

func TestSort(t *testing.T) {
	ints1 := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	ints2 := []int{9, 8, 7, 6, 5, 0, 1, 2, 3, 4}
	ints3 := []int{1, 9, 8, 2, 7, 3, 6, 4, 5, 0}

	for _, ints := range [][]int{ints1, ints2, ints3} {
		expectedResult := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		actualResult := Sort(ints)
		if !reflect.DeepEqual(expectedResult, actualResult) {
			t.Errorf("Sort was incorrect, got: %v, want: %v.", actualResult, expectedResult)
		}
	}
}
