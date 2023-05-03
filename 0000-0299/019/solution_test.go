package main

import "testing"

func TestRemoveNthFromEnd1(t *testing.T) {
	// Test case 1: remove the second node from the end
	list1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}
	expected1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 5}}}}
	result1 := removeNthFromEnd(list1, 2)
	if !isEqual(result1, expected1) {
		t.Errorf("Test case 1 failed: expected %v but got %v", expected1, result1)
	}
}

func TestRemoveNthFromEnd2(t *testing.T) {
	// Test case 2: remove the first node from the end
	list2 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}
	expected2 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}}
	result2 := removeNthFromEnd(list2, 1)
	if !isEqual(result2, expected2) {
		t.Errorf("Test case 2 failed: expected %v but got %v", expected2, result2)
	}
}

func TestRemoveNthFromEnd3(t *testing.T) {
	// Test case 3: remove the fifth node from the end
	list3 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}
	expected3 := &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}
	result3 := removeNthFromEnd(list3, 5)
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
