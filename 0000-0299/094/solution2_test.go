package main

import (
	"reflect"
	"testing"
)

func TestInorderTraversal2EmptyTree(t *testing.T) {
	var root *TreeNode
	expected := []int{}
	res := inorderTraversal2(root)
	if !reflect.DeepEqual(res, expected) {
		t.Errorf("Test case 1 failed: expected %v but got %v", expected, res)
	}
}

func TestInorderTraversal2SingleNode(t *testing.T) {
	root := &TreeNode{Val: 1}
	expected := []int{1}
	res := inorderTraversal2(root)
	if !reflect.DeepEqual(res, expected) {
		t.Errorf("Test case 2 failed: expected %v but got %v", expected, res)
	}
}

func TestInorderTraversal2MultipleNodes(t *testing.T) {
	root := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}
	expected := []int{2, 1, 3}
	res := inorderTraversal2(root)
	if !reflect.DeepEqual(res, expected) {
		t.Errorf("Test case 3 failed: expected %v but got %v", expected, res)
	}
}
