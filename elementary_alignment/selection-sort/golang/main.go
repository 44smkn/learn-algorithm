package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	fmt.Printf("ソート前： %v\n", arr)
	count := sort(arr)
	fmt.Printf("ソート後： %v\n", arr)
	fmt.Println("count:", count)
}

func sort(arr []int) int {
	var count int
	for i := 0; i < len(arr)-1; i++ {
		minj := i
		for j := i + 1; j < len(arr); j++ {
			if arr[minj] > arr[j] {
				minj = j
				count++
			}
		}
		arr[i], arr[minj] = arr[minj], arr[i]
	}
	return count
}
