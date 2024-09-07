package sort

import (
	"slices"
	"testing"
)

func TestItCanPushToStack(t *testing.T) {
	var stack Stack[string]

	stack.Push("Sam")
	stack.Push("Davidge")
	stack.Push("Go")
	stack.Push("Lang")

	if stack.length != 4 {
		t.Fatal(stack.length)
	}
}

func TestItCanPopFromEmptyStack(t *testing.T) {
	var stack Stack[string]

	popped := stack.Pop()

	if popped != "" {
		t.Fatal(popped)
	}
}

func TestItCanPopFromPopulatedStack(t *testing.T) {
	var stack Stack[string]

	stack.Push("Sam")
	stack.Push("Davidge")
	stack.Push("Go")
	stack.Push("Lang")

	popped := stack.Pop()

	if popped != "Lang" {
		t.Fatal(popped)
	}
}

func TestItCanKeepPushPeekingAndPopping(t *testing.T) {
	var stack Stack[string]

	values := []string{"Go", "PHP", "Python", "JS"}

	for _, value := range values {
		stack.Push(value)

		peeked := stack.Peek()

		if peeked != value {
			t.Fatal(peeked)
		}
	}

	slices.Reverse(values)

	for _, value := range values {
		popped := stack.Pop()

		if popped != value {
			t.Fatal(popped)
		}
	}

	if stack.Pop() != "" {
		t.Fatal()
	}

}

func TestItCanPeakEmptyStack(t *testing.T) {
	var stack Stack[string]
	peaked := stack.Peek()

	if peaked != "" {
		t.Fatal(peaked)
	}
}

func TestItCanPeakPopulatedStack(t *testing.T) {
	var stack Stack[string]
	stack.Push("Sam")
	stack.Push("Davidge")

	peaked := stack.Peek()

	if peaked != "Davidge" {
		t.Fatal(peaked)
	}
}
