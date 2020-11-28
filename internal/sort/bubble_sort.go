package sort

func BubbleSort(array []int) {
	for i := 0; i < len(array); i++ {
		for j := i; j < len(array); j++ {
			if array[i] > array[j] {
				swap(array, i, j)
			}
		}
	}
}
