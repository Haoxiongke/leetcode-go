package main

import (
	"reflect"
	"testing"
)

// Function 1
func TestPostorderTraversal1(t *testing.T) {
	root := &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 2}
	root.Right.Left = &TreeNode{Val: 3}
	expected := []int{3, 2, 1}
	res := postorderTraversal(root)
	if !reflect.DeepEqual(res, expected) {
		t.Errorf("expected %v but got %v", expected, res)
	}
}

// Function 2
func TestPostorderTraversal2(t *testing.T) {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}
	root.Right.Left = &TreeNode{Val: 6}
	root.Right.Right = &TreeNode{Val: 7}
	expected := []int{4, 5, 2, 6, 7, 3, 1}
	res := postorderTraversal(root)
	if !reflect.DeepEqual(res, expected) {
		t.Errorf("expected %v but got %v", expected, res)
	}
}

// Function 3
func TestPostorderTraversal3(t *testing.T) {
	root := &TreeNode{Val: 1}
	expected := []int{1}
	res := postorderTraversal(root)
	if !reflect.DeepEqual(res, expected) {
		t.Errorf("expected %v but got %v", expected, res)
	}
}
