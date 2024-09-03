package search

import (
	"testing"
)

func TestItFindsNeedleBySearchingForward(t *testing.T) {

	haystack := getArrayOfIntegers(9, 77)

	found := BinarySearch(haystack, 76)

	if !found {
		t.Fatal(found)
	}
}

func TestBinarySearchFindsNeedleBySearchingBackward(t *testing.T) {

	haystack := getArrayOfIntegers(4, 105)

	found := BinarySearch(haystack, 6)

	if !found {
		t.Fatal(found)
	}
}

func TestBinarySearchFindsNeedleInMiddleElement(t *testing.T) {

	haystack := getArrayOfIntegers(1, 100)

	found := BinarySearch(haystack, 50)

	if !found {
		t.Fatal(found)
	}
}

func TestBinarySearchFindsNeedleAtLastElement(t *testing.T) {

	haystack := getArrayOfIntegers(1, 100)

	found := BinarySearch(haystack, 100)

	if !found {
		t.Fatal(found)
	}
}

func TestBinarySearchFindsNeedleAtFirstElement(t *testing.T) {

	haystack := getArrayOfIntegers(10, 100)

	found := BinarySearch(haystack, 10)

	if !found {
		t.Fatal(found)
	}
}

func TestBinarySearchCannotFindNeedle(t *testing.T) {

	haystack := getArrayOfIntegers(1, 100)

	found := BinarySearch(haystack, 101)

	if found {
		t.Fatal(found)
	}
}

func TestBinarySearchFindsNeedleWhenElementsNotSequential(t *testing.T) {

	haystack := []int{10, 11, 12, 14, 19, 20}

	found := BinarySearch(haystack, 19)

	if !found {
		t.Fatal(found)
	}
}

func getArrayOfIntegers(first int, last int) []int {
	N := last - first + 1
	haystack := make([]int, N)
	insertNumber := first
	for i := range haystack {
		haystack[i] = insertNumber
		insertNumber++
	}
	return haystack
}
