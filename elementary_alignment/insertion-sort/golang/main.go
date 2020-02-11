package main

import (
	"fmt"
)

func main() {
	// 数列の長さを表す整数Nを受け取る
	var n int
	fmt.Scan(&n)

	// ソート対象の数列を読み込み
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	fmt.Printf("ソート前： %v\n", arr)
	sort(arr)
	fmt.Printf("ソート後： %v\n", arr)
}

func sort(arr []int) {
	for i := 1; i < len(arr); i++ { // start at 2nd factor
		j, key := i-1, arr[i] // key is the factor that is concerned
		for ; j >= 0 && arr[j] > key; j-- {
			arr[j+1] = arr[j]
		}

		// why j "+1" ?
		// It is because of j-- calculattion handled before judge conditions
		arr[j+1] = key
	}
}
