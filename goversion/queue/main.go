package main

import (
	"queue/queue"
)

func main() {
	q := queue.NewQueue()
	n1 := queue.NewNode()
	n1.Set("n1", nil)
	n2 := queue.NewNode()
	n2.Set("n2", nil)
	n3 := queue.NewNode()
	n3.Set("n3", nil)
	n4 := queue.NewNode()
	n4.Set("n4", nil)
	q.Push(n1)
	q.Push(n2)
	q.Push(n3)
	q.PrintQueue()
	p := q.Pop()
	q.PrintQueue()
	p.PrintNode()
}
