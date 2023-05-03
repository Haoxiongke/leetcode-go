package main

import "testing"

func TestGetIntersectionNodeOriginal(t *testing.T) {
	// Original test case provided
	listA := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}
	listB := &ListNode{Val: 4, Next: listA.Next}
	expected := listA.Next
	result := getIntersectionNode2(listA, listB)
	if result != expected {
		t.Errorf("Test case failed: expected %v but got %v", expected, result)
	}
}

func TestGetIntersectionNodeFirstNode(t *testing.T) {
	// Test case: intersect at the first node
	listA1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}
	listB1 := &ListNode{Val: 4, Next: listA1}
	expected1 := listA1
	result1 := getIntersectionNode2(listA1, listB1)
	if result1 != expected1 {
		t.Errorf("Test case failed: expected %v but got %v", expected1, result1)
	}
}

func TestGetIntersectionNodeSecondNode(t *testing.T) {
	// Test case: intersect at the second node
	listA2 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}
	listB2 := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: listA2.Next}}
	expected2 := listA2.Next
	result2 := getIntersectionNode2(listA2, listB2)
	if result2 != expected2 {
		t.Errorf("Test case failed: expected %v but got %v", expected2, result2)
	}
}

func TestGetIntersectionNodeNoIntersection(t *testing.T) {
	// Test case: no intersection
	listA3 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}
	listB3 := &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 6}}}
	expected3 := (*ListNode)(nil)
	result3 := getIntersectionNode2(listA3, listB3)
	if result3 != expected3 {
		t.Errorf("Test case failed: expected %v but got %v", expected3, result3)
	}
}
