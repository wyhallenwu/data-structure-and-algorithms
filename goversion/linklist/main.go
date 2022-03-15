package main

import "linklist/linklistpkg"

func main() {
	l := linklistpkg.NewList()
	l.Insert(0, 10)
	l.Insert(0, 20)
	l.Insert(1, 30)
	l.Insert(3, "wuyuheng")
	l.PrintList()

}
