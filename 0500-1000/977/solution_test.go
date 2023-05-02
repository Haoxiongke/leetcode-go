package main

import (
	"reflect"
	"testing"
)

func TestSortedSquaresMixed(t *testing.T) {
	nums := []int{-4, -1, 0, 3, 10}
	expected := []int{0, 1, 9, 16, 100}
	result := sortedSquares(nums)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("sortedSquares(%v) = %v; want %v", nums, result, expected)
	}
}
func TestSortedSquaresNegative(t *testing.T) {
	nums := []int{-7, -3, -1}
	expected := []int{1, 9, 49}
	result := sortedSquares(nums)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("sortedSquares(%v) = %v; want %v", nums, result, expected)
	}
}
func TestSortedSquaresPositive(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	expected := []int{1, 4, 9, 16, 25}
	result := sortedSquares(nums)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("sortedSquares(%v) = %v; want %v", nums, result, expected)
	}
}
