package main

/*
	2.3-6
	Observe that the while loop of lines 5–7 of the INSERTION-SORT procedure in
	Section 2.1 uses a linear search to scan (backward) through the sorted subarray
	A[1..j -1]. Can we use a binary search (see Exercise 2.3-5) instead to improve
	the overall worst-case running time of insertion sort to nlgn ?
*/
import (
	"fmt"

	"github.com/DharmikOO7/alGorithms/helper"
)

func main() {
	var n int
	fmt.Print("Enter the size of array: ")
	fmt.Scan(&n)
	fmt.Printf("Enter %d elements: ", n)
	arr := helper.InputArray(n)
	fmt.Println("Input array: ", arr)
	BinaryInsertionSort(arr)
	fmt.Println("Sorted array: ", arr)
}

// BinaryInsertionSort Insertion sort which uses binary search
func BinaryInsertionSort(a []int) {
	for j := 0; j < len(a); j++ {
		var key = a[j]
		var i = j - 1
		//binary search
		loc := binaryInsertion(a, key, 0, j-1)
		for i >= loc {
			a[i+1] = a[i]
			i--
		}
		a[i+1] = key
	}
}

func binaryInsertion(a []int, key, low, high int) int {
	loc := 0
	for low <= high {
		mid := (low + high) / 2
		if a[mid] == key {
			loc = mid + 1
			break
		}
		if key > a[mid] {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	if low >= high {
		if key > a[low] {
			loc = low + 1
		} else {
			loc = low
		}
	}
	return loc
}

func inputArray(n int) []int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	return arr
}
