package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	fmt.Println(t)
	walkTree(t, ch)
	close(ch)
}

func walkTree(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		walkTree(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		walkTree(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for i := range ch1 {
		if i != <-ch2 {
			return false
		}
	}
	return true
}

func main() {
	myTree := tree.New(1)
	fmt.Println(myTree)
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-ch)
		}
	}()
	Walk(myTree, ch)
	if Same(myTree, myTree) {
		fmt.Println("Trees are identical")
	} else {
		fmt.Println("Trees are NOT identical")
	}
	if Same(tree.New(1), tree.New(1)) {
		fmt.Println("Trees are identical")
	} else {
		fmt.Println("Trees are NOT identical")
	}
	if Same(tree.New(2), tree.New(2)) {
		fmt.Println("Trees are identical")
	} else {
		fmt.Println("Trees are NOT identical")
	}
	if Same(tree.New(2), tree.New(10)) {
		fmt.Println("Trees are identical")
	} else {
		fmt.Println("Trees are NOT identical")
	}
}
