package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

func walkImpl(t *tree.Tree, ch, quit chan int) {
	if t == nil {
		return
	}

	walkImpl(t.Left, ch, quit) // take the leftmost derivation.
	select {
	case ch <- t.Value:
		// Success!
	case <-quit:
		return // Stop leaking of goroutines.
	}
	ch <- t.Value
	walkImpl(t.Right, ch, quit)
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch, quit chan int) {
	walkImpl(t, ch, quit)
	// must close the channel like this, otherwise don't know when to close.
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	w1, w2 := make(chan int), make(chan int)
	quit := make(chan int)
	defer close(quit)

	go Walk(t1, w1, quit)
	go Walk(t2, w2, quit)

	for {
		v1, ok1 := <-w1 // okn returns false when channel empty, else true
		v2, ok2 := <-w2
		if !ok1 || !ok2 {
			return ok1 == ok2 // check the two trees are of the same length (channels end at same time)
		}
		if v1 != v2 {
			return false // If at any point the values do not match, break and return false.
		}
	}
}

func main() {
	fmt.Print("tree.New(1) == tree.New(1): ")
	if Same(tree.New(1), tree.New(1)) {
		fmt.Println("PASSED")
	} else {
		fmt.Println("FAILED")
	}

	fmt.Print("tree.New(1) != tree.New(2): ")
	if !Same(tree.New(1), tree.New(2)) {
		fmt.Println("PASSED")
	} else {
		fmt.Println("FAILED")
	}
}

