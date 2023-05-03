package main

// 方式二：利用双指针同时遍历两次链表
// fast : A -> B
// slow : B -> A
// 如果存在交点，必定会相交，如果不存在交点，则都等于nil
func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	slowP, fastP := headA, headB
	for slowP != fastP {
		if slowP == nil {
			slowP = headB
		} else {
			slowP = slowP.Next
		}

		if fastP == nil {
			fastP = headA
		} else {
			fastP = fastP.Next
		}
	}

	if slowP == nil {
		return nil
	} else {
		return slowP
	}
}
