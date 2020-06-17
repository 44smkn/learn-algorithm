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
	// iは未ソート部分列の先頭インデックス
	for i := 1; i < len(arr); i++ { // 1番目の要素は既にソート済の想定なので、2番目の要素からスタートする
		// jはソート済み部分列からAを挿入するためのループ変数
		j, key := i-1, arr[i]
		for ; j >= 0 && arr[j] > key; j-- {
			arr[j+1] = arr[j]
		}

		// why j "+1" ?
		// 条件をジャッジする前にjのデクリメントが行われているため
		arr[j+1] = key
	}
}
