package main

import (
	"reflect"
	"testing"
)

func TestAverageOfLevels(t *testing.T) {
	// Define test cases
	tests := []struct {
		name   string
		tree   *TreeNode
		expect []float64
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
			expect: []float64{3, 14.5, 11},
		},
		{
			name: "Example 2",
			tree: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 4,
					},
					Right: &TreeNode{
						Val: 5,
					},
				},
				Right: &TreeNode{
					Val: 3,
					Right: &TreeNode{
						Val: 6,
					},
				},
			},
			expect: []float64{1, 2.5, 5},
		},
		{
			name: "Example 3",
			tree: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 4,
						Left: &TreeNode{
							Val: 8,
						},
					},
					Right: &TreeNode{
						Val: 5,
						Right: &TreeNode{
							Val: 9,
						},
					},
				},
				Right: &TreeNode{
					Val: 3,
					Right: &TreeNode{
						Val: 6,
						Right: &TreeNode{
							Val: 10,
						},
					},
				},
			},
			expect: []float64{1, 2.5, 5, 9},
		},
	}

	// Run tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := averageOfLevels(tt.tree)
			if !reflect.DeepEqual(got, tt.expect) {
				t.Errorf("unexpected result, got=%v, expect=%v", got, tt.expect)
			}
		})
	}
}
