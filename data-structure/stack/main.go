package main

import (
	"fmt"
	"strconv"
	"log"
	"errors"
)

func check(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	input := []string{"1", "2", "+", "3", "4", "-", "*"}

	s := newStack(make([]int,10))
	for _, v := range input {
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
			err = s.push(b - a);
			check(err)
		case "*":
			a, err := s.pop()
			check(err)
			b, err := s.pop()
			check(err)
			err = s.push(b * a);
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
	top int
	max int
	slice []int
}

func newStack(slice []int) *stack {
	return  &stack{
		top: 0,
		max: len(slice),
		slice: slice,
	}
}

func (s *stack) isEmpty() bool {
	return s.top == 0
}

func (s *stack) isFull() bool{
	return s.top == s.max - 1
}

func (s *stack) push(x int)error {
	if s.isFull() {
		return errors.New("Overflow")
	}
	s.top++
	s.slice[s.top] = x
	return nil
}

func (s *stack) pop() (int, error) {
	if s.isEmpty(){
		return 0, errors.New("Underflow")
	}
	s.top--
	return s.slice[s.top+1], nil
}