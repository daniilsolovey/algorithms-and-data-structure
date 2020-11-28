package sort

func swap(array []int, i, j int) {
	tmp := array[i]
	array[i] = array[j]
	array[j] = tmp
}
