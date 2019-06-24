package main

import (
	"fmt"
)

func main() {
	arr := []int{5, 2, 4, 6, 1, 3}
	sort(arr)
	fmt.Println(arr)
}

func sort(arr []int) {
	for i:=1;i<len(arr);i++ {  // start at 2nd factor
		j, key := i-1, arr[i]  // key is the factor that is concerned
		for ;j >= 0 && arr[j] > key; j-- {
			arr[j+1] = arr[j]
		}

		// why j "+1" ?
		// It is because of j-- calculattion handled before judge conditions
		arr[j+1] = key  
	}
}