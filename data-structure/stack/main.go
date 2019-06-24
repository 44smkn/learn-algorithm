package main

import (
	"errors"
)

func main() {

}

type stack struct {
	top int
	max int
	slice []int
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

func (s *stack) pop() (*int, error) {
	if s.isEmpty(){
		return nil, errors.New("Underflow")
	}
	s.top--
	return &s.slice[s.top+1], nil
}