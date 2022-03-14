package main

import "stack/stack"

func main() {
	s := stack.NewStack()
	s.Push("1")
	s.Push("2")
	s.Push("3")
	s.Push("4")
	s.PrintStack()
	s.Pop()
	s.PrintStack()
}
