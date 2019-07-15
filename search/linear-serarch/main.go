package main

import (
	"fmt"
)


func search(a []int, key, n int) bool {
	a = append(a, key)

	// 番兵を利用して比較演算子をへらす
	i := 0
	for a[i]!=key { i++ } 
	return i != n
}

func main() {
	a := []int{1, 2, 3, 4 ,5}
	q := []int{3, 4, 1}
	n := len(a)

	var sum int
	for _, e := range q {
		if search(a, e, n) {
			sum++
		}
	}
	fmt.Println(sum)
}