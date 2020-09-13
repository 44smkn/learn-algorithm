package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"sync"
)

func main() {
	judge(os.Stdin, os.Stdout)
}

func judge(r io.Reader, w io.Writer) {
	n := scanRowText(r)
	result := 0
	for _, v := range n {
		d, _ := strconv.Atoi(string(v))
		result += d
		result %= 9
	}
	if result == 0 {
		fmt.Fprintln(w, "Yes")
		return
	}
	fmt.Fprintln(w, "No")
}

var once sync.Once
var scanner *bufio.Scanner

func scanRowText(r io.Reader) string {
	once.Do(func() {
		scanner = bufio.NewScanner(r)
		scanner.Split(bufio.ScanWords)
	})
	scanner.Scan()
	return scanner.Text()
}
