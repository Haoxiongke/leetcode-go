package main

type MyStack struct {
	queue []int
}

func Constructor() MyStack {
	return MyStack{
		queue: make([]int, 0),
	}
}

func (this *MyStack) Push(x int) {
	this.queue = append(this.queue, x)
}

func (this *MyStack) Pop() int {
	len := len(this.queue)
	if len == 0 {
		return -1
	}
	val := this.queue[len-1]
	this.queue = this.queue[:len-1]
	return val
}

func (this *MyStack) Top() int {
	len := len(this.queue)
	if len == 0 {
		return -1
	}
	val := this.queue[len-1]
	return val
}

func (this *MyStack) Empty() bool {
	return len(this.queue) == 0
}

func main() {

}
