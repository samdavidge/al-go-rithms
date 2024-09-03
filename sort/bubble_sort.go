package sort

func BubbleSort(values []int) []int {

	for i := 0; i < len(values); i++ {
		for j := 0; j < len(values)-1-i; j++ {
			if values[j] > values[j+1] {
				temp := values[j]
				values[j] = values[j+1]
				values[j+1] = temp
			}
		}
	}
	return values
}
