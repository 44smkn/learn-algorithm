package main

import (
	"fmt"
)

func main() {
	fmt.Println(getSum(1, 2))
	fmt.Println(getSum(-2, 3))

	fmt.Println(getSum(1, 3))
}

func getSum(a int, b int) int {
	for b != 0 {
		c := (a & b) << 1
		a ^= b
		b = c
	}
	return a
}
