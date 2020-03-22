package main

/*
	2.3-7
	Describe a nlgn time algorithm that, given a set S of n integers and another
	integer x, determines whether or not there exist two elements in S whose sum is
	exactly x.
*/
import (
	"fmt"

	"github.com/DharmikOO7/alGorithms/Sorting/sorting"
	"github.com/DharmikOO7/alGorithms/helper"
)

func main() {
	var n, x int
	fmt.Print("Enter the size of set: ")
	fmt.Scan(&n)
	fmt.Printf("Enter %d elements of set: ", n)
	S := helper.InputArray(n)
	fmt.Print("Enter the sum: ")
	fmt.Scan(&x)
	fmt.Println("Input array: ", S)
	sorting.MergeSort(0, len(S)-1, S)
	fmt.Println("Sorted array: ", S)
	for i := 0; i < len(S); i++ {
		diffElIdx := binarySearch(S, i+1, len(S)-1, x-S[i])
		if diffElIdx != -1 {
			fmt.Printf("%d and %d equals %d: ", S[i], S[diffElIdx], x)
			break
		}
	}
}

func binarySearch(S []int, low, high, val int) int {
	for low <= high {
		mid := (low + high) / 2
		if S[mid] == val {
			return mid
		} else if S[mid] < val {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}
