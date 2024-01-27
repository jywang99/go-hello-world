package binaryTree

import (
	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		Walk(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		Walk(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	// create channels, make sure they're closed after calling Walk
	c1 := make(chan int)
	c2 := make(chan int)
	walkAndClose := func(t *tree.Tree, ch chan int) {
		Walk(t, ch)
		close(ch)
	}

	go walkAndClose(t1, c1)
	go walkAndClose(t2, c2)
	for {
		// read channel output
		e1, ok1 := <-c1
		e2, ok2 := <-c2
		// both channels empty
		if !ok1 && !ok2 {
			return true
		}
		// one of the channels empty = not same size
		if !ok1 || !ok2 {
			return false
		}
		// one of the elements different
		if e1 != e2 {
			return false
		}
	}
}

