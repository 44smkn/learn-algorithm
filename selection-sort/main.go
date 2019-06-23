package main

import (
	"fmt"
)

func main() {
	arr := []int{5, 6, 4, 2, 1, 3}
	count := sort(arr)
	fmt.Println("sorted array:", arr)
	fmt.Println("count:", count)
}

func sort(arr []int) int {
	var count int
	for i := 0; i < len(arr)-1; i++ {
		minj := i
		for j := i+1; j < len(arr); j++ {
			if arr[minj] > arr[j] {
				minj=j
				count++
			}
		}
		arr[i], arr[minj] = arr[minj], arr[i]
	}
	return count
}
