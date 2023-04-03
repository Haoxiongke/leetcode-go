package main

import (
	"reflect"
	"testing"
)

func TestPostorderTraversal2EmptyTree(t *testing.T) {
	var root *TreeNode
	expected := []int{}
	res := postorderTraversal2(root)
	if !reflect.DeepEqual(res, expected) {
		t.Errorf("Test case 1 failed: expected %v but got %v", expected, res)
	}
}

func TestPostorderTraversal2SingleNode(t *testing.T) {
	root := &TreeNode{Val: 1}
	expected := []int{1}
	res := postorderTraversal2(root)
	if !reflect.DeepEqual(res, expected) {
		t.Errorf("Test case 2 failed: expected %v but got %v", expected, res)
	}
}

func TestPostorderTraversal2MultipleNodes(t *testing.T) {
	root := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}
	expected := []int{2, 3, 1}
	res := postorderTraversal2(root)
	if !reflect.DeepEqual(res, expected) {
		t.Errorf("Test case 3 failed: expected %v but got %v", expected, res)
	}
}
