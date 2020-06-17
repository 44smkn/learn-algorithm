package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type node struct {
	key        int
	prev, next *node
}

func main() {
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Scan()
	n, _ := strconv.Atoi(stdin.Text())

	for i := 0; i < n; i++ {
		stdin.Scan()
		switch stdin.Text() 
	}

	cur := null.next
	for cur != null {
		fmt.Println(cur.key)
		cur = cur.next
	}
}

// 番兵
var null = &node{}

func init() {
	null.prev = null
	null.next = null
}

func insert(key int) {
	x := &node{
		key:  key,
		next: null.next,
		prev: null,
	}
	null.next.prev = x
	null.next = x
}

func listSearch(key int) *node {
	cur := null.next
	for cur != null && cur.key != key {
		cur = cur.next
	}
	return cur
}

func deleteNode(t *node) {
	if t == null {
		return
	}
	t.next.prev = t.prev
	t.prev.next = t.next
}

func deleteFirst() {
	deleteNode(null.next)
}

func deleteLast() {
	deleteNode(null.prev)
}

func deleteKey(key int) {
	deleteNode(listSearch(key))
}
