package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	dummyNode := &ListNode{
		Next: head,
	}
	preNode := dummyNode
	for head != nil && head.Next != nil {
		preNode.Next = head.Next
		next := head.Next.Next
		head.Next.Next = head
		head.Next = next
		preNode = head
		head = next
	}

	return dummyNode.Next
}

func main() {

}
