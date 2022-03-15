package linklistpkg

import "fmt"

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
