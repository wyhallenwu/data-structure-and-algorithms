package linklistpkg

import (
	"testing"
)

func TestAll(t *testing.T) {
	l := NewList()
	l.Insert(0, 10)
	l.Insert(0, 20)
	l.Insert(1, 30)
	l.Insert(2, "test")
	l.PrintList()
}
