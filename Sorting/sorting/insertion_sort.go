package sorting

// InsertionSort Implementation of insertion sort alogrithm
func InsertionSort(arr []int) {
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
