package linklistpkg

import (
	"fmt"
	"testing"
)

func TestAll(t *testing.T) {
	head := NewNode("head")
	err := Insert(head, 1, "test1")
	if err != nil {
		fmt.Println(err)
	}
	err = Insert(head, 2, "test2")
	err = Insert(head, 3, "test3")
	err = Insert(head, 1, "test insert")
	PrintList(head)
	err = Delete(head, 1)
	PrintList(head)
}
