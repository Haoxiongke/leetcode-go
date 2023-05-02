package main

import "testing"

func TestRemoveElementSuccess(t *testing.T) {
	nums := []int{3, 2, 2, 3}
	val := 3
	expected := 2
	result := removeElement(nums, val)
	if result != expected {
		t.Errorf("removeElement(%v, %d) = %d; want %d", nums, val, result, expected)
	}
	for i := 0; i < result; i++ {
		if nums[i] == val {
			t.Errorf("removeElement(%v, %d) = %v; want %v", nums, val, nums, []int{2, 2})
		}
	}
}
func TestRemoveElementFail(t *testing.T) {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val := 5
	expected := 8
	result := removeElement(nums, val)
	if result != expected {
		t.Errorf("removeElement(%v, %d) = %d; want %d", nums, val, result, expected)
	}
	for i := 0; i < result; i++ {
		if nums[i] == val {
			t.Errorf("removeElement(%v, %d) = %v; want %v", nums, val, nums, []int{0, 1, 2, 2, 3, 0, 4, 2})
		}
	}
}
func TestRemoveElementEmpty(t *testing.T) {
	nums := []int{}
	val := 5
	expected := 0
	result := removeElement(nums, val)
	if result != expected {
		t.Errorf("removeElement(%v, %d) = %d; want %d", nums, val, result, expected)
	}
}
