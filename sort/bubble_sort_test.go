package sort

import (
	"reflect"
	"testing"
)

func TestBubbleSortsArrayWithUniqueValues(t *testing.T) {
	in := []int{44, 12, 53, 9, 4, 2, 89}
	expected := []int{2, 4, 9, 12, 44, 53, 89}
	out := BubbleSort(in)
	if !reflect.DeepEqual(expected, out) {
		t.Fatal(out)
	}
}

func TestBubbleSortsArrayWithDuplicateValues(t *testing.T) {
	in := []int{4, 5, 6, 4, 5, 6, 4, 5, 6}
	expected := []int{4, 4, 4, 5, 5, 5, 6, 6, 6}
	out := BubbleSort(in)
	if !reflect.DeepEqual(expected, out) {
		t.Fatal(out)
	}
}
