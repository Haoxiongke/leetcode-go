package main

import "testing"

func TestRemoveElements1(t *testing.T) {
	// Test case 1: remove all instances of 1
	list1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1, Next: &ListNode{Val: 3}}}}
	expected1 := &ListNode{Val: 2, Next: &ListNode{Val: 3}}
	result1 := removeElements(list1, 1)
	if !isEqual(result1, expected1) {
		t.Errorf("Test case 1 failed: expected %v but got %v", expected1, result1)
	}
}

func TestRemoveElements2(t *testing.T) {
	// Test case 2: remove all instances of 2
	list2 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1, Next: &ListNode{Val: 3}}}}
	expected2 := &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 3}}}
	result2 := removeElements(list2, 2)
	if !isEqual(result2, expected2) {
		t.Errorf("Test case 2 failed: expected %v but got %v", expected2, result2)
	}
}

func TestRemoveElements3(t *testing.T) {
	// Test case 3: remove all instances of 3
	list3 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1, Next: &ListNode{Val: 3}}}}
	expected3 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1}}}
	result3 := removeElements(list3, 3)
	if !isEqual(result3, expected3) {
		t.Errorf("Test case 3 failed: expected %v but got %v", expected3, result3)
	}
}

func isEqual(l1 *ListNode, l2 *ListNode) bool {
	for l1 != nil && l2 != nil {
		if l1.Val != l2.Val {
			return false
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	return l1 == nil && l2 == nil
}
