package main

import (
	"reflect"
	"testing"
)

func TestLevelOrder(t *testing.T) {
	// Define test cases
	tests := []struct {
		name   string
		tree   *TreeNode
		expect [][]int
	}{
		{
			name: "Example 1",
			tree: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 9,
				},
				Right: &TreeNode{
					Val: 20,
					Left: &TreeNode{
						Val: 15,
					},
					Right: &TreeNode{
						Val: 7,
					},
				},
			},
			expect: [][]int{
				{3},
				{9, 20},
				{15, 7},
			},
		},
		// Add more test cases here
	}

	// Run tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := levelOrder(tt.tree)
			if !reflect.DeepEqual(got, tt.expect) {
				t.Errorf("unexpected result, got=%v, expect=%v", got, tt.expect)
			}
		})
	}
}
