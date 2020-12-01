package stack

func BinarySearch(a []int, search int) (result int, searchCount int) {
	mid := len(a) / 2
	switch {
	case len(a) == 0:
		result = -1 // not found
	case a[mid] > search:
		result, searchCount = BinarySearch(a[:mid], search)
	case a[mid] < search:
		result, searchCount = BinarySearch(a[mid+1:], search)
		result += mid + 1
	default: // a[mid] == search
		result = mid // found
	}

	searchCount++
	return
}
