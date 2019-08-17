package stack

type Stack struct {
	items    []interface{}
	position int
}

const CAPACITY = 1024

func NewStack() *Stack {
	return &Stack{
		items:    make([]interface{}, CAPACITY),
		position: -1}
}

func (stack *Stack) Push(item interface{}) bool {
	if stack.isFull() {
		newSpace := make([]interface{}, 2*len(stack.items))
		copy(newSpace, stack.items)
		stack.items = newSpace
	}

	stack.position++
	stack.items[stack.position] = item

	return true
}

func (stack *Stack) Pop() interface{} {
	if stack.isEmpty() {
		return nil
	}

	last := stack.items[stack.position]

	stack.items[stack.position] = nil
	stack.position--

	return last
}

func (stack *Stack) Peek() interface{} {
	if stack.isEmpty() {
		return nil
	}

	return stack.items[stack.position]
}

func (stack *Stack) isEmpty() bool {
	return stack.position == -1
}

func (stack *Stack) isFull() bool {
	return stack.position == len(stack.items)-1
}
