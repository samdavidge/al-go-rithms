package sort

type Node[T interface{}] struct {
	value    T
	previous *Node[T]
}

type Stack[T interface{}] struct {
	head   *Node[T]
	length int
}

func (stack *Stack[T]) Push(item T) {
	node := Node[T]{value: item}
	stack.length++

	if stack.head == nil {
		stack.head = &node
		return
	}
	node.previous = stack.head
	stack.head = &node
}

func (stack *Stack[T]) Pop() T {
	var value T
	if stack.head == nil {
		return value
	}

	stack.length--
	head := stack.head
	stack.head = head.previous
	return head.value
}

func (stack *Stack[T]) Peek() T {
	var value T
	if stack.head != nil {
		value = stack.head.value
	}

	return value
}
