package sorting

// SelectionSort Implementation of selection sort alogrithm
func SelectionSort(arr []int) {
	for j := 0; j < len(arr)-1; j++ {
		var minIdx = j
		//find minimum element of rest of array
		for i := j + 1; i < len(arr); i++ {
			if arr[i] < arr[minIdx] {
				minIdx = i
			}
		}
		//swap if min element si not same as current element
		if minIdx != j {
			arr[j], arr[minIdx] = arr[minIdx], arr[j]
		}
	}
}
