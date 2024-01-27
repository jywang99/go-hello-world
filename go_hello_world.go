package main

import (
	"fmt"

	"golang.org/x/tour/tree"
    "github.com/jyking99/go_hello_world/packs/binaryTree"
)

func main() {
	tree1 := tree.New(1)
	tree2 := tree.New(1)
	fmt.Println("tree1 and tree2 are the same: ", Same(tree1, tree2))
}

