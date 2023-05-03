package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 方式一：先遍历两个链表的长度，长的先多走长度len(A) - len(B), 然后同时走
func getIntersectionNode(headA, headB *ListNode) *ListNode {

	if headA == nil || headB == nil {
		return nil
	}

	curNodeA, curNodeB := headA, headB
	var lenA, lenB, diff int
	for curNodeA != nil {
		curNodeA = curNodeA.Next
		lenA++
	}

	for curNodeB != nil {
		curNodeB = curNodeB.Next
		lenB++
	}

	diff = lenB - lenA
	if lenA > lenB {
		tmp := headA
		headA = headB
		headB = tmp
		diff = lenA - lenB
	}

	for i := 0; headA != headB; i++ {
		headB = headB.Next
		if i >= diff {
			headA = headA.Next
		}
	}

	if headA == nil {
		return nil
	}
	return headA
}

func main() {
	headA := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}

	headB := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val:  5,
			Next: nil,
		},
	}
	getIntersectionNode(headA, headB)

}
