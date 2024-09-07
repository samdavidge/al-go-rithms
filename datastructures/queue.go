package datastructures

type QueueNode[T interface{}] struct {
	value T
	next  *QueueNode[T]
}

type Queue[T interface{}] struct {
	head   *QueueNode[T]
	tail   *QueueNode[T]
	length int
}

func (queue *Queue[T]) Peek() T {
	// @todo: When peaking/dequeueing when head is empty I am returning
	// an empty value. This complicates queueing an empty value. I imagine
	// the Go approach would be to return multiple values (value, success),
	// find out if it is.
	var value T
	if queue.head != nil {
		value = queue.head.value
	}

	return value
}

func (queue *Queue[T]) Enqueue(item T) {
	queue.length++
	node := &QueueNode[T]{value: item}

	if queue.tail == nil {
		queue.head = node
		queue.tail = node
		return
	}

	queue.tail.next = node
	queue.tail = node
}

func (queue *Queue[T]) Dequeue() T {
	if queue.head == nil {
		var empty T
		return empty
	}

	queue.length--

	head := queue.head
	queue.head = queue.head.next

	return head.value
}
