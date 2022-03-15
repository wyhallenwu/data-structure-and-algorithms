package linklistpkg

import (
	"fmt"
	"log"
)

type Any interface{}

type Node struct {
	Info Any
	Next *Node
	Pre  *Node
}

func NewNode(info Any) *Node {
	n := &Node{Info: info}
	return n
}

func (n *Node) PrintNode() {
	fmt.Print(n.Info)
}

type Linklist struct {
	head   *Node
	length int
}

func NewList() *Linklist {
	head := &Linklist{NewNode("head"), 0}
	return head
}

// Insert inserts a new node at the index of ix
func (l *Linklist) Insert(ix int, info Any) {
	in := NewNode(info)
	// if ix is out of length or length is 0 then insert to 1st position
	if ix > l.length || l.length == 0 {
		ix = 0
	}
	count := 0
	flag := l.head
	for {
		if count < ix {
			flag = flag.Next
			count += 1
		} else {
			// insert position is not tail
			if flag.Next != nil {
				flag.Next.Pre = in
			}
			flag.Next, in.Next, in.Pre = in, flag.Next, flag
			break
		}
	}
	l.length += 1

}

func (l *Linklist) PrintList() {
	n := l.head.Next
	fmt.Println("length is: ", l.length)
	for {
		if n != nil {
			n.PrintNode()
			fmt.Printf("->")
			n = n.Next
		} else {
			fmt.Println("END")
			break
		}
	}
}

func (l *Linklist) Delete(ix int) {
	if l.length == 0 {
		log.Fatal("empty can't delete")
	} else if ix >= l.length {
		ix = 0
		log.Println("out of length. Delete 1st position")
	}
	count := 0
	flag := l.head.Next
	for {
		if count < ix {
			flag = flag.Next
			count += 1
		} else {
			//if delete head
			if count == 0 {
				flag.Pre.Next = flag.Next
				flag = nil
			} else if flag.Next == nil { //if delete tail
				flag.Pre.Next = nil
				flag = nil
			} else {
				flag.Pre.Next, flag.Next.Pre = flag.Next, flag.Pre
				flag = nil
			}
			break
		}
	}
	l.length -= 1
}
