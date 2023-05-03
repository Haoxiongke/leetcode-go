package main

import "testing"

func TestGetIntersectionNode1(t *testing.T) {
	// Test case 1: intersect at node with value 8
	listA1 := &ListNode{Val: 4, Next: &ListNode{Val: 1, Next: &ListNode{Val: 8, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}
	listB1 := &ListNode{Val: 5, Next: &ListNode{Val: 0, Next: &ListNode{Val: 1, Next: listA1.Next.Next}}}
	expected1 := listA1.Next.Next
	result1 := getIntersectionNode(listA1, listB1)
	if result1 != expected1 {
		t.Errorf("Test case 1 failed: expected %v but got %v", expected1, result1)
	}
}

func TestGetIntersectionNode2(t *testing.T) {
	// Test case 2: intersect at node with value 2
	listA2 := &ListNode{Val: 0, Next: &ListNode{Val: 9, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4}}}}}
	listB2 := &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: listA2.Next.Next.Next}}
	expected2 := listA2.Next.Next.Next
	result2 := getIntersectionNode(listA2, listB2)
	if result2 != expected2 {
		t.Errorf("Test case 2 failed: expected %v but got %v", expected2, result2)
	}
}

func TestGetIntersectionNode3(t *testing.T) {
	// Test case 3: no intersection
	listA3 := &ListNode{Val: 2, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}}
	listB3 := &ListNode{Val: 1, Next: &ListNode{Val: 5}}
	expected3 := (*ListNode)(nil)
	result3 := getIntersectionNode(listA3, listB3)
	if result3 != expected3 {
		t.Errorf("Test case 3 failed: expected %v but got %v", expected3, result3)
	}
}
