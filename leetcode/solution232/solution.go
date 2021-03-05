package solution232

type MyQueue struct {
	inStack  []int
	outStack []int
}

func Constructor() MyQueue {
	return MyQueue{}
}

func (queue *MyQueue) Push(x int) {
	queue.inStack = append(queue.inStack, x)
}

func (queue *MyQueue) inToOut() {
	for len(queue.inStack) > 0 {
		queue.outStack = append(queue.outStack, queue.inStack[len(queue.inStack)-1])
		queue.inStack = queue.inStack[:len(queue.inStack)-1]
	}
}

func (queue *MyQueue) Pop() int {
	if len(queue.outStack) == 0 {
		queue.inToOut()
	}
	x := queue.outStack[len(queue.outStack)-1]
	queue.outStack = queue.outStack[:len(queue.outStack)-1]
	return x
}

func (queue *MyQueue) Peek() int {
	if len(queue.outStack) == 0 {
		queue.inToOut()
	}
	return queue.outStack[len(queue.outStack)-1]
}

func (queue *MyQueue) Empty() bool {
	return len(queue.inStack) == 0 && len(queue.outStack) == 0
}
