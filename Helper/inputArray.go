package helper

import "fmt"

// InputArray returns a n-element user entered array
func InputArray(n int) []int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	return arr
}
