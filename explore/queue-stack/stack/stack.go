package stack

type Stack struct {
	items []interface{}
}

func NewStack() *Stack {
	return &Stack{}
}

func (stack *Stack) Push(item interface{}) {
	stack.items = append(stack.items, item)
}

func (stack *Stack) Pop() interface{} {
	length := len(stack.items)
	if length == 0 {
		return nil
	}

	last := stack.items[length-1]
	stack.items = stack.items[:length-1]

	return last
}

func (stack *Stack) Top() interface{} {
	length := len(stack.items)
	if length == 0 {
		return nil
	}

	return stack.items[len(stack.items)-1]
}
