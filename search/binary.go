package search

func BinarySearch(haystack []int, needle int) bool {

	lo := 0
	hi := len(haystack)

	for lo < hi {

		middle := lo + (hi-lo)/2
		current := haystack[middle]

		if current == needle {
			return true
		} else if current > needle {
			hi = middle
		} else {
			lo = middle + 1
		}
	}

	return false
}
