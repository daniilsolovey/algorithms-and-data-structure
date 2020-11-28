package sort

func SelectSort(array []int) {
	for i := 0; i < len(array)-1; i++ {
		min := i
		for j := i + 1; j < len(array); j++ {
			if array[min] > array[j] {
				min = j
			}
		}

		if min != i {
			swap(array, i, min)
		}
	}
}
