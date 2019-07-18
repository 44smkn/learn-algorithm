package main

import (
	"fmt"
)

func main() {
	arr := []int{5, 3, 2, 4, 1}
	count := sort(arr)
	fmt.Println("sorted array:", arr)
	fmt.Println("inversion count:", count)
}

func sort(arr []int) int {
	var count int
	for i := 0; i < len(arr)-1; i++ {
		for j := len(arr) - 1; j > i; j-- {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
				count++
			}
		}
	}
	return count
}
