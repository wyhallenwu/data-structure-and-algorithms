package main

import (
	"fmt"
	"stack/stack"
)

func main() {
	top := stack.NewStack("top")
	top = stack.Push(top, "wuyuheng")
	stack.PrintStack(top)
	fmt.Printf("%p\n", top)
	top = stack.Push(top, "yuhengwu")
	fmt.Printf("%p\n", top)
	stack.PrintStack(top)
	top = stack.Push(top, "test")
	fmt.Printf("%p\n", top)
	stack.PrintStack(top)
	stack.SetTop(top, "newtop")
	stack.PrintStack(top)
}
