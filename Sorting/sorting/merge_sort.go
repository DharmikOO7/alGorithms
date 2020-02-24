package sorting

import (
	"math"
)

// MergeSort Implementation of merge sort alogrithm
func MergeSort(st int, end int, arr []int) {
	if end == st+1 {
		if arr[st] > arr[end] {
			arr[st], arr[end] = arr[end], arr[st]
		}
	} else if end != st {
		mid := st + (end-st)/2
		MergeSort(st, mid, arr)
		MergeSort(mid+1, end, arr)
		merge(arr, st, mid, end)
	}
}

func merge(arr []int, st, mid, end int) {
	i := 0
	j := 0
	n1 := mid - st + 1
	n2 := end - mid
	leftSub := make([]int, n1+1)
	rightSub := make([]int, n2+1)
	for i := 0; i < n1; i++ {
		leftSub[i] = arr[st+i]
	}
	leftSub[n1] = math.MaxInt64
	for j := 0; j < n2; j++ {
		rightSub[j] = arr[mid+1+j]
	}
	rightSub[n2] = math.MaxInt64
	for k := st; k < end+1; k++ {
		if leftSub[i] < rightSub[j] {
			arr[k] = leftSub[i]
			i++
		} else {
			arr[k] = rightSub[j]
			j++
		}
	}
}
