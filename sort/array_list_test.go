package sort

import (
	"testing"
)

func TestItCanPushToArrayList(t *testing.T) {
	var arrayList ArrayList[string]

	arrayList.Push("Some string")

	if arrayList.values[0] != "Some string" {
		t.Fatal()
	}
}

func TestItCanIncreaseCapacityWhenPushedTo(t *testing.T) {
	var arrayList ArrayList[string]

	capacity := BaseCapacity * 10

	for i := 0; i < capacity; i++ {
		arrayList.Push("s")
	}

	if len(arrayList.values) != capacity {
		t.Fatal(len(arrayList.values))
	}
}

func TestItCanPopAValue(t *testing.T) {
	var arrayList ArrayList[string]

	for _, value := range []string{"A", "B", "C", "D"} {
		arrayList.Push(value)
	}

	popped, err := arrayList.Pop()

	if popped != "D" {
		t.Fatal(popped)
	}

	if err != nil {
		t.Fatal(err)
	}
}

func TestItReturnsAnErrorWhenPoppingEmptyValue(t *testing.T) {
	var arrayList ArrayList[string]

	_, err := arrayList.Pop()

	if err == nil {
		t.Fatal(err)
	}
}

func TestItCanGetValueAtIndex(t *testing.T) {
	var arrayList ArrayList[string]

	for _, value := range []string{"A", "B", "C", "D"} {
		arrayList.Push(value)
	}

	value, _ := arrayList.Get(2)

	if value != "C" {
		t.Fatal(value)
	}
}

func TestItReturnsAnErrorWhenGettingAnEmptyIndex(t *testing.T) {
	var arrayList ArrayList[string]

	for _, value := range []string{"A", "B", "C", "D"} {
		arrayList.Push(value)
	}

	_, err := arrayList.Get(4)

	if err == nil {
		t.Fatal(err)
	}
}
