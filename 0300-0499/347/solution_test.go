package main

import (
	"reflect"
	"testing"
)

// Function 1: Test for empty input
func TestTopKFrequent_EmptyInput(t *testing.T) {
	nums := []int{1}
	k := 1
	result := topKFrequent(nums, k)
	expected := []int{1}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v but got %v", expected, result)
	}
}

// Function 3: Test for normal input
func TestTopKFrequent_NormalInput(t *testing.T) {
	nums := []int{1, 1, 1, 2, 2, 3}
	k := 2
	result := topKFrequent(nums, k)
	expected := []int{1, 2}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v but got %v", expected, result)
	}
}
