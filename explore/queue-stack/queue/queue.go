package queue

type Queue struct {
	data     []int
	position int
}

func NewQueue() *Queue {
	return &Queue{position: 0}
}

func (queue *Queue) EnQueue(value int) bool {
	queue.data = append(queue.data, value)
	return true
}

func (queue *Queue) DeQueue() bool {
	queue.position++
	return true
}

func (queue *Queue) Front() int {
	if queue.IsEmpty() {
		return -1
	}
	return queue.data[queue.position]
}

func (queue *Queue) IsEmpty() bool {
	return queue.position == len(queue.data)
}
