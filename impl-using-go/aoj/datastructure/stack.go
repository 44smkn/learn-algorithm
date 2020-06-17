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

func check(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Scan()

	s := newStack(make([]int, 10))
	for _, v := range strings.Split(stdin.Text(), " ") {
		switch v {
		case "+":
			a, err := s.pop()
			check(err)
			b, err := s.pop()
			check(err)
			err = s.push(b + a)
			check(err)
		case "-":
			a, err := s.pop()
			check(err)
			b, err := s.pop()
			check(err)
			err = s.push(b - a)
			check(err)
		case "*":
			a, err := s.pop()
			check(err)
			b, err := s.pop()
			check(err)
			err = s.push(b * a)
			check(err)
		default:
			num, err := strconv.Atoi(v)
			check(err)
			err = s.push(num)
			check(err)
		}
	}

	fmt.Println(s.slice[s.top])
}

type stack struct {
	top   int
	max   int
	slice []int
}

func newStack(slice []int) *stack {
	return &stack{
		top:   0,
		max:   len(slice),
		slice: slice,
	}
}

func (s *stack) isEmpty() bool {
	return s.top == 0
}

func (s *stack) isFull() bool {
	return s.top == s.max-1
}

func (s *stack) push(x int) error {
	if s.isFull() {
		return errors.New("Overflow")
	}
	s.top++
	s.slice[s.top] = x
	return nil
}

func (s *stack) pop() (int, error) {
	if s.isEmpty() {
		return 0, errors.New("Underflow")
	}
	s.top--
	return s.slice[s.top+1], nil
}
