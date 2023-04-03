package main

import (
	"reflect"
	"testing"
)

func TestPostorderTraversal2EmptyTree(t *testing.T) {
	root := &TreeNode{Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}}
	expected := [][]int{{3}, {9, 20}, {15, 7}}
	if res := levelOrder(root); !reflect.DeepEqual(res, expected) {
		t.Errorf("levelOrder(%v) = %v, expected %v", root, res, expected)
	}
}
