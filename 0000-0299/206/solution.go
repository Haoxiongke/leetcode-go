package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	// 当头节点不为空
	var preNode *ListNode
	for head != nil {
		tmp := head.Next
		head.Next = preNode
		preNode = head
		head = tmp
	}

	// 如果头节点为空，直接就返回
	return preNode
}

func main() {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val:  2,
			Next: nil,
		},
	}
	println(reverseList(head))
}
