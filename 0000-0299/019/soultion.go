package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummyNode := &ListNode{
		Next: head,
	}
	slowP, fastP := dummyNode, dummyNode

	for i := 0; fastP != nil; i++ {
		fastP = fastP.Next
		if i > n {
			slowP = slowP.Next
		}
	}

	slowP.Next = slowP.Next.Next

	return dummyNode.Next
}

func main() {

}
