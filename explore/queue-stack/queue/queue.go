package queue

type Queue struct {
	items    []interface{}
	position int
}

func NewQueue() *Queue {
	return &Queue{position: 0}
}

func (queue *Queue) EnQueue(item interface{}) bool {
	queue.items = append(queue.items, item)
	return true
}

func (queue *Queue) DeQueue() bool {
	if queue.IsEmpty() {
		return false
	}

	queue.position++
	return true
}

func (queue *Queue) Front() interface{} {
	if queue.IsEmpty() {
		return -1
	}
	return queue.items[queue.position]
}

func (queue *Queue) IsEmpty() bool {
	return queue.position == len(queue.items)
}
