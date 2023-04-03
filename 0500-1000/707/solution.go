package main

type MyLinkedList struct {
	Next *MyLinkedList
	Val  int
}

func Constructor() MyLinkedList {
	return MyLinkedList{}
}

func (this *MyLinkedList) Get(index int) int {
	if this == nil {
		return -1
	}
	curNode := this
	curIndex := 0
	for curNode.Next != nil {
		if curIndex == index {
			return curNode.Val
		}
		curNode = curNode.Next
	}
	return -1
}

func (this *MyLinkedList) AddAtHead(val int) {
	tmp := this
	this.Val = val
	this.Next = tmp
}

func (this *MyLinkedList) AddAtTail(val int) {
	if this == nil {
		this = &MyLinkedList{
			Val:  val,
			Next: nil,
		}
	} else {
		curNode := this
		for curNode.Next != nil {
			curNode = curNode.Next
		}
		curNode.Next = &MyLinkedList{
			Val:  val,
			Next: nil,
		}
	}
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index <= 0 {
		this.AddAtHead(val)
	} else {
		dummyHead := &MyLinkedList{}
		dummyHead.Next = this
		curNode := dummyHead
		curIndex := 0
		for curNode.Next != nil {
			if curIndex == index {
				curNode.Next = &MyLinkedList{
					Val:  val,
					Next: curNode.Next.Next,
				}
				break
			} else {
				curNode = curNode.Next
				curIndex++
			}
		}
	}
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	dummyHead := &MyLinkedList{}
	dummyHead.Next = this
	curNode := dummyHead
	curIndex := 0
	for curNode.Next != nil {
		if curIndex == index {
			curNode.Next = curNode.Next.Next
			break
		} else {
			curNode = curNode.Next
			curIndex++
		}

	}
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */

func main() {

}
