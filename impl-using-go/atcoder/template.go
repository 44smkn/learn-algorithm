package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	solve(os.Stdin, os.Stdout)
}

func solve(r io.Reader, w io.Writer) {
	scanner := initScanner(r)
	// 解く
	fmt.Fprintln(w /* 答え */)
}

// utilityメソッド群

const maxCapacity = 512 * 1024 // デフォルト値は64*1024

func initScanner(r io.Reader) *bufio.Scanner {
	scanner := bufio.NewScanner(r)
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)
	scanner.Split(bufio.ScanWords) // スペースごとに読み込む
	return scanner
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func scanUint64(scanner *bufio.Scanner) uint64 {
	scanner.Scan()
	val, err := strconv.ParseUint(scanner.Text(), 10, 64)
	check(err)
	return val
}

func scanInt64(scanner *bufio.Scanner) int64 {
	scanner.Scan()
	val, err := strconv.ParseInt(scanner.Text(), 10, 64)
	check(err)
	return val
}

func scanFloat64(scanner *bufio.Scanner) float64 {
	scanner.Scan()
	val, err := strconv.ParseFloat(scanner.Text(), 64)
	check(err)
	return val
}

func scanString(scanner *bufio.Scanner) string {
	scanner.Scan()
	return scanner.Text()
}
