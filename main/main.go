package main

import (
	"fmt"

    "go-hello-world/interfaces"
    "go-hello-world/concurrency"
)

func main() {
    fmt.Println("Testing stringer...")
    interfaces.TestStringer()
    fmt.Println()
    fmt.Println("Testing reader...")
    interfaces.TestReader()
    fmt.Println()
    fmt.Println("Testing error...")
    interfaces.TestError()
    fmt.Println()

    fmt.Println("Testing binary tree...")
    concurrency.TestBinaryTree()
    fmt.Println()
    fmt.Println("Testing crawler...")
    concurrency.TestCralwer()
    fmt.Println()
}

