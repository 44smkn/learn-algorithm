package main

import (
	"fmt"
)

func main() {
	fmt.Println(simplegcd(74, 54))
	fmt.Println(efficientgcd(74, 54))
}

func simplegcd(smaller, bigger int) int {
	for d := smaller;d>0;d-- {
		if smaller % d == 0 && bigger %d == 0 {
			return d
		}
	}
	return 1
}

func efficientgcd(smaller, bigger int) int {
	for smaller > 0 {
		r := bigger % smaller
		bigger = smaller
		smaller = r
	}
	return bigger
}