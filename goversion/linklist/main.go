package main

import (
	"fmt"
	"linklist/linklistpkg"
)

func main() {
	head := linklistpkg.NewNode("head")
	err := linklistpkg.Insert(head, 1, "test1")
	if err != nil {
		fmt.Println(err)
	}
	linklistpkg.Insert(head, 2, "test2")
	err = linklistpkg.Insert(head, 0, "test3")
	fmt.Println(err)
	linklistpkg.PrintList(head)
	linklistpkg.Delete(head, 0)
	linklistpkg.PrintList(head)
	linklistpkg.SetInfo(head, 1, "test4")
	linklistpkg.PrintList(head)
}
