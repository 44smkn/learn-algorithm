package main

import (
	"testing"
)

func initVar() (a, q []int, n int) {
	a = makeRange(0, 10000000)
	q = []int{3890000, 4900005, 12890000, 9808009}
	n = len(a)
	return
}

func makeRange(min, max int) []int {
    a := make([]int, max-min+1)
    for i := range a {
        a[i] = min + i
    }
    return a
}

func BenchmarkSearchBySentinel(b *testing.B) {
	a, q, n := initVar()

	b.ResetTimer()
	var sum int
	for _, e := range q {
		if SearchBySentinel(a, e, n) {
			sum++
		}
	}
}

func BenchmarkSearchByRange(b *testing.B) {
	a, q, n := initVar()

	b.ResetTimer()
	var sum int
	for _, e := range q {
		if SearchByRange(a, e, n) {
			sum++
		}
	}
}