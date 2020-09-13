package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
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

func scanRowText(r io.Reader) string {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	return scanner.Text()
}
