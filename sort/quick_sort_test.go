package sort

import (
	"reflect"
	"testing"
)

func TestItCanQuickSortReverseSortedArray(t *testing.T) {
	in := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	out := QuickSort(in)

	if !reflect.DeepEqual(expected, out) {
		t.Fatal(out)
	}
}

func TestItCanQuickSortSortedArray(t *testing.T) {
	in := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	out := QuickSort(in)

	if !reflect.DeepEqual(expected, out) {
		t.Fatal(out)
	}
}

func TestItCanQuickSortRandomArray(t *testing.T) {
	in := []int{2, 3, 8, 6, 4, 1, 5, 9, 7}
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	out := QuickSort(in)

	if !reflect.DeepEqual(expected, out) {
		t.Fatal(out)
	}
}
