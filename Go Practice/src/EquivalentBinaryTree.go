package main

import(
	"golang.org/x/tour/tree"
	"fmt"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	defer close(ch) //the Golang way of close
	
	// define a function in Walk() and recur that function,
	// in order to properly close the channel
	var recurWalk func(t *tree.Tree) 
	recurWalk = func(t *tree.Tree) {
		if (t==nil) { return }
		recurWalk(t.Left)
		ch <- t.Value
		recurWalk(t.Right)
	}
	
	recurWalk(t)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	
	for {
		v1, ok1 := <- ch1
		v2, ok2 := <- ch2
		
		if ok1 != ok2 || v1 != v2  {
			return false
		}
		if !ok1 {
			break
		}
	}
	return true
}

func main() {
	ch := make(chan int, 10);
	go Walk(tree.New(1), ch)
	for i := range ch {
		fmt.Printf("%v ",i)
	}
	fmt.Printf("\n%v", Same(tree.New(1), tree.New(2)))
}
