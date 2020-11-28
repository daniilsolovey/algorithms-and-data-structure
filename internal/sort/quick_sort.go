package sort

func Quicksort(array []int) {
	if len(array) <= 1 {
		return
	}

	split := splitArray(array)

	Quicksort(array[:split])
	Quicksort(array[split:])
}

func splitArray(array []int) int {
	pivot := array[len(array)/2]

	left := 0
	right := len(array) - 1

	for {
		for ; array[left] < pivot; left++ {
		}

		for ; array[right] > pivot; right-- {
		}

		if left >= right {
			return right
		}

		swap(array, left, right)
	}
}
