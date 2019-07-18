package main 

import (
	"fmt"
)



func binarySearch(a []int, key int) bool {
	left := 0
	right := len(a)

	for left < right {
		mid := (left + right) / 2
		if mid == key {
			return true
		} else if key < a[mid] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return false
}

func main() {
	a := []int{1, 2, 3, 4 ,5}
	q := []int{3, 4, 1}

	var sum int
	for _, e := range q {
		if binarySearch(a, e) {
			sum++
		}
	}
	fmt.Println(sum)
}