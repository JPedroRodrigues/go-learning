package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}

	Walk(t.Left, ch)
	ch <- t.Value
	Walk(t.Right, ch)
}

func Same(t1, t2 *tree.Tree) bool {
	if t1 == nil && t2 == nil {
		return true
	} else if t1 == nil || t2 == nil {
		return false
	} else if t1.Value != t2.Value {
		return false
	}

	x := Same(t1.Left, t2.Left)
	y := Same(t1.Right, t2.Right)

	return x && y
}

func main() {
	n := 1
	ch := make(chan int)
	t1 := tree.New(n)

	go Walk(t1, ch)

	for range n * 10 {
		fmt.Println(<-ch)
	}

	result := Same(t1, t1)
	fmt.Println(result)
}
