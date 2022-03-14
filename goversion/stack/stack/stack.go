package stack

import "fmt"

type Any interface{}

type Node struct {
	Info Any
	Next *Node
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

func Push(top *Node, info Any) *Node {
	n := NewNode(info)
	// pointer copy is just copy the info contained
	n.Next, top = top, n
	return top
}

func NewStack(info Any) *Node {
	top := NewNode(info)
	return top
}

func PrintStack(top *Node) {
	for {
		if top != nil {
			top.PrintNode()
			fmt.Printf("->")
			top = top.Next
		} else {
			fmt.Printf("->||END\n")
			break
		}
	}
}
