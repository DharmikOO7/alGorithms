package sorting

// InsertionSort Implementation of insertion sort alogrithm
func InsertionSort(a []int) {
	for j := 0; j < len(a); j++ {
		var key = a[j]
		var i = j - 1
		for i > -1 && a[i] > key {
			a[i+1] = a[i]
			i--
		}
	}
}

// InsertionSortRecursive recursviely sorts A[0..n-2] and inserts a[n-1] into sorted array
func InsertionSortRecursive(arr []int, n int) {
	if n == 0 {
		return
	}
	InsertionSortRecursive(arr, n-1)
	key := arr[n]
	i := n - 1
	for i > -1 && arr[i] > key {
		arr[i+1] = arr[i]
		i--
	}
	arr[i+1] = key
}
