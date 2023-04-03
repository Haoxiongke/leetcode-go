package main

import (
	"reflect"
	"testing"
)

// Test case 1: Test with a single node tree
func TestInorderTraversalSingleNode(t *testing.T) {
	root := &TreeNode{Val: 1}
	expected := []int{1}
	res := inorderTraversal(root)
	if !reflect.DeepEqual(res, expected) {
		t.Errorf("Expected %v but got %v", expected, res)
	}
}

// Test case 2: Test with a tree with only left nodes
func TestInorderTraversalOnlyLeftNodes(t *testing.T) {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Left.Left = &TreeNode{Val: 3}
	root.Left.Left.Left = &TreeNode{Val: 4}
	expected := []int{4, 3, 2, 1}
	res := inorderTraversal(root)
	if !reflect.DeepEqual(res, expected) {
		t.Errorf("Expected %v but got %v", expected, res)
	}
}

// Test case 3: Test with a tree with only right nodes
func TestInorderTraversalOnlyRightNodes(t *testing.T) {
	root := &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 2}
	root.Right.Right = &TreeNode{Val: 3}
	root.Right.Right.Right = &TreeNode{Val: 4}
	expected := []int{1, 2, 3, 4}
	res := inorderTraversal(root)
	if !reflect.DeepEqual(res, expected) {
		t.Errorf("Expected %v but got %v", expected, res)
	}
}
