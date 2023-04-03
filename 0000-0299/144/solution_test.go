package main

import (
	"reflect"
	"testing"
)

// Test case 1: Test when the root is nil
func TestPreorderTraversalCase1(t *testing.T) {
	var root *TreeNode
	expected := []int{}
	res := preorderTraversal(root)
	if !reflect.DeepEqual(res, expected) {
		t.Errorf("Expected %v but got %v", expected, res)
	}
}

// Test case 2: Test when the tree has only one node
func TestPreorderTraversalCase2(t *testing.T) {
	root := &TreeNode{Val: 1}
	expected := []int{1}
	res := preorderTraversal(root)
	if !reflect.DeepEqual(res, expected) {
		t.Errorf("Expected %v but got %v", expected, res)
	}
}

// Test case 3: Test when the tree has multiple nodes
func TestPreorderTraversalCase3(t *testing.T) {
	root := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}
	expected := []int{1, 2, 3}
	res := preorderTraversal(root)
	if !reflect.DeepEqual(res, expected) {
		t.Errorf("Expected %v but got %v", expected, res)
	}
}
