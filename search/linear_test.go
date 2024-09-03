package search

import (
	"testing"
)

func TestLinearSearchFindsNeedle(t *testing.T) {
	haystack := []string{"Sam", "Vikki", "Teresa", "Joe"}
	found := LinearSearch(haystack, "Teresa")
	if !found {
		t.Fatal(found)
	}
}

func TestLinearSearchCannotFindNeedle(t *testing.T) {
	haystack := []string{"Sam", "Vikki", "Teresa", "Joe"}
	found := LinearSearch(haystack, "Sa")
	if found {
		t.Fatal(found)
	}
}
