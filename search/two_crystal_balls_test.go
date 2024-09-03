package search

import (
	"testing"
)

func TestItFindsAtFirstPosition(t *testing.T) {
	breaks := []bool{true, true, true, true, true, true}
	position := TwoCrystalBalls(breaks)
	if position != 0 {
		t.Fatal(position)
	}
}

func TestItFindsAtLastPosition(t *testing.T) {
	breaks := []bool{false, false, false, false, false, true}
	position := TwoCrystalBalls(breaks)
	if position != 5 {
		t.Fatal(position)
	}
}

func TestItFindsInMiddlePosition(t *testing.T) {
	breaks := []bool{false, false, false, true, true, true, true}
	position := TwoCrystalBalls(breaks)
	if position != 3 {
		t.Fatal(position)
	}
}

func TestItReturnsNegativeOneWhenNotFound(t *testing.T) {
	breaks := []bool{false, false, false, false}
	position := TwoCrystalBalls(breaks)
	if position != -1 {
		t.Fatal(position)
	}
}
