package sort

func QuickSort(values []int) []int {
	sort(&values, 0, len(values)-1)
	return values
}

func sort(values *[]int, lo int, hi int) {
	if lo >= hi {
		return
	}

	pivotIndex := partition(values, lo, hi)

	sort(values, lo, pivotIndex-1)
	sort(values, pivotIndex+1, hi)
}

func partition(values *[]int, lo int, hi int) int {
	pivot := (*values)[hi]

	index := lo - 1

	for i := lo; i < hi; i++ {
		if (*values)[i] <= pivot {
			index++
			temp := (*values)[i]
			(*values)[i] = (*values)[index]
			(*values)[index] = temp
		}
	}

	index++

	(*values)[hi] = (*values)[index]
	(*values)[index] = pivot

	return index
}
