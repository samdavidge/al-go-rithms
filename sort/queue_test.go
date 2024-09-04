package sort

import (
	"testing"
)

func TestItCanEnqueueFirstValue(t *testing.T) {
	var queue Queue[int]
	queue.Enqueue(123)
	queuedValue := queue.head.value

	if queuedValue != 123 {
		t.Fatal(queuedValue)
	}
}

func TestItCanDequeueFirstValue(t *testing.T) {
	var queue Queue[int]
	queue.Enqueue(123)

	dequeued := queue.Dequeue()

	if dequeued != 123 {
		t.Fatal(dequeued)
	}

	if queue.head != nil {
		t.Fatal()
	}
}

func TestItCanEnqueueThenDequeueMultipleValues(t *testing.T) {
	var queue Queue[string]

	items := []string{"Following", "Primes", "Algorithm", "Course"}

	for _, item := range items {
		queue.Enqueue(item)
	}

	for _, item := range items {
		deqeued := queue.Dequeue()
		if deqeued != item {
			t.Fatal(deqeued)
		}
	}

	if queue.head != nil {
		t.Fatal()
	}
}

func TestItCanPeakEmptyQueue(t *testing.T) {
	var queue Queue[string]

	item := queue.Peek()

	if item != "" {
		t.Fatal(item)
	}
}

func TestItCanPeakQueueWithOneItem(t *testing.T) {
	var queue Queue[string]
	node := QueueNode[string]{value: "sam"}
	queue.head = &node
	value := queue.Peek()

	if node.value != value {
		t.Fatal(value)
	}
}

func TestItCanPeakQueueWithMultipleItems(t *testing.T) {
	var queue Queue[string]

	names := []string{"Sam", "Bob", "John"}

	for _, name := range names {
		queue.Enqueue(name)
	}

	value := queue.Peek()

	if value != "Sam" {
		t.Fatal(value)
	}
}
