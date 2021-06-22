package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	WalkR(t, ch)
	close(ch)
}

func WalkR(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}

	WalkR(t.Left, ch)
	ch <- t.Value
	WalkR(t.Right, ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for {
        v1, ok1 := <-ch1
        v2, ok2 := <-ch2

		if (v1 != v2 || ok1 != ok2) {
			return false
		}

        if !ok1 {
            break
        }
	}
	return true
}

func main() {
	ch := make(chan int)
	go Walk(tree.New(1), ch)
	for {
        v, ok := <-ch
		fmt.Printf("%v ", v)
        
        if !ok {
            break
        }
	}

	fmt.Println()
    fmt.Println("1 and 1 are same:", Same(tree.New(1), tree.New(1)))
    fmt.Println("1 and 2 are same:", Same(tree.New(1), tree.New(2)))
}
