package main

import (
	"fmt"

	"github.com/algorithms-and-data-structure/internal/sort"
	"github.com/algorithms-and-data-structure/internal/stack"
)

func main() {
	var data []int

	data = []int{3, 4, 1, 2, 7, 8, 9, 10}
	fmt.Println("data before quick sort: ", data)
	sort.Quicksort(data)
	fmt.Println("data after quick sort: ", data)

	data = []int{3, 4, 1, 2, 7, 8, 9, 10}
	fmt.Println("data before bubble sort: ", data)
	sort.BubbleSort(data)
	fmt.Println("data after bubble sort: ", data)

	data = []int{3, 4, 1, 2, 7, 8, 9, 10}
	fmt.Println("data before select sort: ", data)
	sort.SelectSort(data)
	fmt.Println("data after select sort: ", data)

	data = []int{3, 4, 1, 2, 7, 8, 9, 10, 100, 200, 95, 340}
	fmt.Println("data before insertion sort: ", data)
	sort.InsertionSort(data)
	fmt.Println("data after insertion sort: ", data)

	data = []int{3, 4, 1, 2, 7, 8, 9, 10, 100, 200, 95, 340}
	fmt.Println("data before shell sort: ", data)
	sort.ShellSort(data)
	fmt.Println("data after shell sort: ", data)

	slice := stack.NewSlice(data)
	HandleSlice(slice)
}

func HandleSlice(slice stack.Stack) {
	fmt.Println(slice)
	slice.Clear()
	fmt.Println(slice)
}
