package linklistpkg

import (
	"errors"
	"fmt"
)

type Any interface{}

type Node struct {
	Info Any
	Next *Node
}

func (n *Node) Setter(any Any) {
	n.Info = any
}

func NewNode(info Any) *Node {
	node := new(Node)
	node.Info = info
	node.Next = nil
	return node
}

func Insert(head *Node, index int, insertInfo Any) error {
	ix := 0
	insertNode := NewNode(insertInfo)
	for {
		if ix == index {
			head.Next, insertNode.Next = insertNode, head.Next
			return nil
		}
		if head.Next == nil {
			head.Next, insertNode.Next = insertNode, head.Next
			return errors.New("out of length. Insert to the tail")
		}
		head = head.Next
		ix += 1
	}
}

func Delete(head *Node, index int) error {
	ix := 0
	for {
		if ix == index {
			head.Next = head.Next.Next
			return nil
		}
		if head.Next == nil {
			return errors.New("can't delete. Beyond length")
		}
		head = head.Next
		ix += 1
	}
}

func PrintList(head *Node) {
	for {
		if head != nil {
			fmt.Print(head.Info, "->")
		} else {
			fmt.Print("||END\n")
			break
		}
		head = head.Next
	}
}

func SetInfo(head *Node, index int, any Any) error {
	ix := 0
	for {
		if head == nil {
			return errors.New("out of length")
		}
		if index == ix {
			head.Setter(any)
			return nil
		}
		ix += 1
		head = head.Next
	}

}