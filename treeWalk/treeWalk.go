package main

import (
    "code.google.com/p/go-tour/tree"
    "fmt"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
    if t == nil {
     	return
    }
    Walk(t.Left, ch)
    ch <- t.Value
    Walk(t.Right, ch)
    
}

// Walker launches Walk in a new goroutine,
// and returns a read-only channel of values.
func Walker(t *tree.Tree) <-chan int {
	ch := make(chan int)
	go func() {
		Walk(t, ch)
		close(ch)
	}()
	return ch
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
    c1, c2 := Walker(t1), Walker(t2)
    for {
        v1, ok1 := <- c1
        v2, ok2 := <- c2
        if !ok1 || !ok2 {
        	return ok1 == ok2
        }
        if v1 != v2 {
         	break 
        }
    }
    return false
}

func main() {
  /*  ch := make(chan int)
    go Walk(tree.New(1), ch)
    for i := range ch {
     	fmt.Println(i)   
    } */
    fmt.Println(Same(tree.New(1), tree.New(1)))
    fmt.Println(Same(tree.New(1), tree.New(2)))
}
