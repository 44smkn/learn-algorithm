package main

import (
	"fmt"
)

type node struct {
	key int
	prev, next *node
}

func main() {
	insert(5)
	insert(2)
	insert(3)
	insert(1)
	deleteKey(3)
	insert(6)
	deleteKey(5)

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
		key: key,
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