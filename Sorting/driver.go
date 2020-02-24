package main

import (
	"fmt"

	"github.com/DharmikOO7/alGorithms/Sorting/sorting"
)

func main() {
	var n int
	fmt.Print("Enter the size of array: ")
	fmt.Scan(&n)
	fmt.Printf("Enter %d elements: ", n)
	arr := inputArray(n)
	fmt.Println("Input array: ", arr)
	fmt.Println("Please select sorting algorithm: \n 1. Insertion sort \n 2. Selection sort \n 3. Merge Sort")
	var choice int
	fmt.Scan(&choice)
	switch choice {
	case 1:
		sorting.InsertionSort(arr)
	case 2:
		sorting.SelectionSort(arr)
	case 3:
		sorting.MergeSort(0, len(arr)-1, arr)
	}
	fmt.Println("Sorted array: ", arr)
}

func inputArray(n int) []int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	return arr
}
