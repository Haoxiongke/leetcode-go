package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	slowPt, fastPt := head, head
	for fastPt != nil && fastPt.Next != nil {
		fastPt = fastPt.Next.Next
		slowPt = slowPt.Next
		if fastPt == slowPt {
			slowPt = head
			for fastPt != slowPt {
				fastPt = fastPt.Next
				slowPt = slowPt.Next
			}
			return fastPt
		}
	}
	return nil
}

func main() {

}
