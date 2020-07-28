package main

import (
	"fmt"
	"golang.org/x/tour/tree"
	"sort"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	queue := make([]*tree.Tree, 0)

	if t != nil {
		queue = append(queue, t)
	}
	
	for len(queue) != 0 {
		temp := queue[0]
		ch <- temp.Value
		queue = queue[1:]
		if temp.Left != nil {
		queue = append(queue, temp.Left)
		}
		if temp.Right != nil {
		queue = append(queue, temp.Right)
		}
	}
	
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	c1 := make(chan int, 10)
	c2 := make(chan int, 10)
	a1 := make([]int, 0)
	a2 := make([]int, 0)
	go Walk(t2, c1)
	go Walk(t1, c2)
	
	for i := range c1 {
		a1 = append(a1, i)
	}
	
	for j := range c2 {
		a2 = append(a2, j)
	}
	
	sort.Ints(a1)
	sort.Ints(a2)
	
	for index := range a1 {
		fmt.Println(a1[index], a2[index])
		if (a1[index] != a2[index]) {
			return false
		}
	}
	
	return true
}

func main() {
	c := make(chan int, 10)
	
	go Walk(tree.New(2), c)
	for i := range c {
		fmt.Println(i)
	}
	
	fmt.Println(Same(tree.New(2), tree.New(1)))
}
