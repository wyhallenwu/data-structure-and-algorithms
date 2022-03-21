package queue

import "fmt"

type Any interface {
}

type node struct {
	info Any
	next *node
}

func NewNode() *node {
	return &node{}
}

func CopyNode(cp *node) *node {
	return &node{cp.info, cp.next}
}

func (n *node) Set(info Any, next *node) {
	n.info = info
	n.next = next
}

func (n *node) PrintNode() {
	fmt.Println(n.info)
}

type queue struct {
	front  *node
	rear   *node
	length int
}

func NewQueue() *queue {
	return &queue{nil, nil, 0}
}

func (q *queue) Front() Any {
	return q.front.info
}

func (q *queue) Rear() Any {
	return q.rear.info
}

func (q *queue) Push(n *node) {
	if q.IsEmpty() {
		q.front, q.rear = n, n
		q.front.next = q.rear
		q.length += 1
		return
	}
	q.rear.next = n
	q.rear = n
	q.length += 1
}

func (q *queue) Pop() *node {
	if q.IsEmpty() {
		return nil
	}
	p := CopyNode(q.front)
	q.front = q.front.next
	q.length -= 1
	return p
}

func (q *queue) Len() int {
	return q.length
}

func (q *queue) IsEmpty() bool {
	if q.length == 0 {
		return true
	}
	return false
}

func (q *queue) PrintQueue() {
	n := q.front
	for {
		if n != nil {
			n.PrintNode()
			n = n.next
		} else {
			break
		}
	}
}
