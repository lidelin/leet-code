package queue

type CircularQueue struct {
	Capacity int
	values   []int
	head     int
	tail     int
	length   int
}

/** Initialize your data structure here. Set the size of the queue to be k. */
func NewCircularQueue(k int) CircularQueue {
	return CircularQueue{
		Capacity: k,
		values:   make([]int, k),
		head:     0,
		tail:     0,
		length:   0}
}

/** Insert an element into the circular queue. Return true if the operation is successful. */
func (queue *CircularQueue) EnQueue(value int) bool {
	if queue.IsFull() {
		return false
	}

	if queue.length == 0 {
		queue.values[queue.tail] = value
	} else {
		queue.tail = (queue.tail + 1) % queue.Capacity
		queue.values[queue.tail] = value
	}

	queue.length++

	return true
}

/** Delete an element from the circular queue. Return true if the operation is successful. */
func (queue *CircularQueue) DeQueue() bool {
	if queue.IsEmpty() {
		return false
	}

	queue.head = (queue.head + 1) % queue.Capacity
	queue.length--

	return true
}

/** Get the front item from the queue. */
func (queue *CircularQueue) Front() int {
	if queue.IsEmpty() {
		return -1
	}
	return queue.values[queue.head]
}

/** Get the last item from the queue. */
func (queue *CircularQueue) Rear() int {
	if queue.IsEmpty() {
		return -1
	}
	return queue.values[queue.tail]
}

/** Checks whether the circular queue is empty or not. */
func (queue *CircularQueue) IsEmpty() bool {
	return queue.length == 0
}

/** Checks whether the circular queue is full or not. */
func (queue *CircularQueue) IsFull() bool {
	return queue.length == queue.Capacity
}
