package stack

import (
	"fmt"
	"log"
)

type Any interface{}

type Node struct {
	Info Any
	Next *Node
}

type Stack struct {
	top    *Node
	length int
}

func NewNode(info Any) *Node {
	n := new(Node)
	n.Next = nil
	n.Info = info
	return n
}

func (n *Node) PrintNode() {
	fmt.Print(n.Info)
}

// NewStack creates a new top node for a stack
func NewStack() *Stack {
	stack := &Stack{NewNode("top"), 0}
	return stack

}

func (s *Stack) Push(info Any) {
	n := NewNode(info)
	// pointer copy is just copy the info contained
	s.top.Next, n.Next = n, s.top.Next
	s.length += 1
}

func (s *Stack) Pop() {
	if s.top.Next != nil {
		s.top.Next = s.top.Next.Next
	} else {
		log.Println("Empty. can't pop")
	}
}

func (s *Stack) SetTop(any Any) {
	s.top.Next.Info = any
}

func (s *Stack) PrintStack() {
	n := s.top.Next
	for {
		if n != nil {
			n.PrintNode()
			fmt.Printf("->")
			n = n.Next
		} else {
			fmt.Printf("END\n")
			break
		}
	}
}
