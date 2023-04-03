package main

import (
	"reflect"
	"testing"
)

func Test_topKFrequent2_case1(t *testing.T) {
	nums := []int{1, 1, 1, 2, 2, 3}
	k := 2
	expected := []int{2, 1}
	result := topKFrequent2(nums, k)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected: %v, but got: %v", expected, result)
	}
}

func Test_topKFrequent2_case2(t *testing.T) {
	nums := []int{1}
	k := 1
	expected := []int{1}
	result := topKFrequent2(nums, k)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected: %v, but got: %v", expected, result)
	}
}
