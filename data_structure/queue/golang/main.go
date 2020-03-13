package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type process struct {
	name string
	time int
}

func check(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func min(num1, num2 int) int {
	if num1 < num2 {
		return num1
	} else {
		return num2
	}
}

func main() {
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Scan()
	firstline := strings.Split(stdin.Text(), " ")
	n, _ := strconv.Atoi(firstline[0])
	quantum, _ := strconv.Atoi(firstline[1])

	processes := make([]process, n)
	for i := 0; i < n; i++ {
		stdin.Scan()
		input := strings.Split(stdin.Text(), " ")
		time, _ := strconv.Atoi(input[1])
		proc := process{
			name: input[0],
			time: time,
		}
		processes[i] = proc
	}

	q := newQueue(make([]process, 8))
	for _, v := range processes {
		err := q.enqueue(&v)
		check(err)
	}

	fmt.Println("\nresult:")
	var elaps int
	for !q.isEmpty() {
		p, err := q.dequeue()
		check(err)
		c := min(quantum, p.time)
		p.time -= c
		elaps += c
		if p.time > 0 {
			err := q.enqueue(p)
			check(err)
		} else {
			fmt.Printf("%v %v\n", p.name, elaps)
		}
	}
}

type queue struct {
	head  int
	tail  int
	slice []process
	max   int
}

func newQueue(slice []process) *queue {
	return &queue{
		head:  0,
		tail:  0,
		max:   len(slice),
		slice: slice,
	}
}

func (q *queue) isEmpty() bool {
	return q.head == q.tail
}

// On array of Ring Buffer, head is next to tail
func (q *queue) isFull() bool {
	return q.head == (q.tail+1)%q.max
}

func (q *queue) enqueue(x *process) error {
	if q.isFull() {
		return errors.New("Overflow")
	}
	q.slice[q.tail] = *x
	if q.tail == q.max-1 {
		q.tail = 0
	} else {
		q.tail++
	}
	return nil
}

func (q *queue) dequeue() (*process, error) {
	if q.isEmpty() {
		return nil, errors.New("Underflow")
	}
	x := &q.slice[q.head]
	if q.head == q.max-1 {
		q.head = 0
	} else {
		q.head++
	}
	return x, nil
}
