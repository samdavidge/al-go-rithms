package search

import (
	"math"
)

func TwoCrystalBalls(breaks []bool) int {

	jumpAmount := int(math.Floor(math.Sqrt(float64(len(breaks)))))
	i := jumpAmount

	for ; i < len(breaks); i += jumpAmount {
		if breaks[i] {
			break
		}
	}

	searchPosition := i - jumpAmount

	for ; searchPosition < i && searchPosition < len(breaks); searchPosition++ {
		if breaks[searchPosition] {
			return searchPosition
		}
	}

	return -1
}
