package main

import "testing"

func TestSearchSuccess(t *testing.T) {
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 9
	expected := 4
	result := search(nums, target)
	if result != expected {
		t.Errorf("search(%v, %d) = %d; want %d", nums, target, result, expected)
	}
}
func TestSearchFail(t *testing.T) {
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 2
	expected := -1
	result := search(nums, target)
	if result != expected {
		t.Errorf("search(%v, %d) = %d; want %d", nums, target, result, expected)
	}
}
func TestSearchEmpty(t *testing.T) {
	nums := []int{}
	target := 5
	expected := -1
	result := search(nums, target)
	if result != expected {
		t.Errorf("search(%v, %d) = %d; want %d", nums, target, result, expected)
	}
}
