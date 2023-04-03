package main

import (
	"reflect"
	"testing"
)

// Test case 1: root is nil
func Test_preorderTraversal2_case1(t *testing.T) {
	var root *TreeNode
	expected := []int{}
	res := preorderTraversal2(root)
	if !reflect.DeepEqual(res, expected) {
		t.Errorf("Expected %v but got %v", expected, res)
	}
}

// Test case 2: root has only one node
func Test_preorderTraversal2_case2(t *testing.T) {
	root := &TreeNode{Val: 1}
	expected := []int{1}
	res := preorderTraversal2(root)
	if !reflect.DeepEqual(res, expected) {
		t.Errorf("Expected %v but got %v", expected, res)
	}
}

// Test case 3: root has multiple nodes
func Test_preorderTraversal2_case3(t *testing.T) {
	root := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}
	expected := []int{1, 2, 3}
	res := preorderTraversal2(root)
	if !reflect.DeepEqual(res, expected) {
		t.Errorf("Expected %v but got %v", expected, res)
	}
}
