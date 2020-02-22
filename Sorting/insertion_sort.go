package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Print("Enter the size of array: ")
	fmt.Scan(&n)
	fmt.Printf("Enter %d elements: ", n)
	arr := inputArray(n)
	fmt.Println("Input array: ", arr)
	insertionSort(arr)
	fmt.Println("Sorted array: ", arr)
}

func inputArray(n int) []int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	return arr
}

func insertionSort(arr []int) {
	for j := 0; j < len(arr); j++ {
		var key = arr[j]
		var i = j - 1
		for i > -1 && arr[i] > key {
			arr[i+1] = arr[i]
			i--
		}
		arr[i+1] = key
	}
}
