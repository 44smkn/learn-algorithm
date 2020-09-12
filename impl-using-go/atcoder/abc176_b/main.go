package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sync"
)

func main() {
	n := scanRowText()
	result := 0
	for _, v := range n {
		d, _ := strconv.Atoi(string(v))
		result += d
		result %= 9
	}
	if result == 0 {
		fmt.Println("Yes")
		return
	}
	fmt.Println("No")
}

var once sync.Once
var scanner *bufio.Scanner

func scanRowText() string {
	once.Do(func() {
		scanner = bufio.NewScanner(os.Stdin)
		scanner.Split(bufio.ScanWords)
	})
	scanner.Scan()
	return scanner.Text()
}
