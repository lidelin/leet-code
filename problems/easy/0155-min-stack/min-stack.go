package stack

type MinStack struct {
	items    []int
	minimums []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (minStack *MinStack) Push(x int) {
	minStack.items = append(minStack.items, x)

	if len(minStack.minimums) == 0 {
		minStack.minimums = append(minStack.minimums, x)
	} else {
		min := minStack.GetMin()
		if x < min {
			min = x
		}

		minStack.minimums = append(minStack.minimums, min)
	}
}

func (minStack *MinStack) Pop() {
	minStack.items = minStack.items[:len(minStack.items)-1]
	minStack.minimums = minStack.minimums[:len(minStack.minimums)-1]
}

func (minStack *MinStack) Top() int {
	return minStack.items[len(minStack.items)-1]
}

func (minStack *MinStack) GetMin() int {
	length := len(minStack.minimums)

	if length == 0 {
		return 0
	}

	return minStack.minimums[length-1]
}
