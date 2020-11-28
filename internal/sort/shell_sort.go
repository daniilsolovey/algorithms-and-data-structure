package sort

func ShellSort(array []int) {
	for gap := len(array) / 2; gap > 0; gap /= 2 {
		for i := gap; i < len(array); i++ {
			x := array[i]
			j := i
			for ; j >= gap && array[j-gap] > x; j -= gap {
				array[j] = array[j-gap]
			}

			array[j] = x
		}
	}
}
