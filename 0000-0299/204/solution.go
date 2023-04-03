package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	dummyHead := &ListNode{}
	dummyHead.Next = head
	currentNode := dummyHead
	for currentNode.Next != nil {
		if currentNode.Next.Val == val {
			currentNode.Next = currentNode.Next.Next
		} else {
			currentNode = currentNode.Next
		}
	}

	return dummyHead.Next
}

func main() {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 6,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val: 5,
							Next: &ListNode{
								Val:  6,
								Next: nil,
							},
						},
					},
				},
			},
		},
	}
	println(removeElements(head, 6))
}
