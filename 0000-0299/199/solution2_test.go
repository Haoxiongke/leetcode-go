package main

import (
	"reflect"
	"testing"
)

func Test_rightSideView2(t *testing.T) {
	tests := []struct {
		name   string
		tree   *TreeNode
		expect []int
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
			expect: []int{3, 20, 7},
		},
		// Add more test cases here
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rightSideView2(tt.tree); !reflect.DeepEqual(got, tt.expect) {
				t.Errorf("rightSideView2() = %v, expect %v", got, tt.expect)
			}
		})
	}
}
