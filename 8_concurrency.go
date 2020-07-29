package main

import (
	"fmt"
	"golang.org/x/tour/tree"
	"sync"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	if (t != nil) {
	wg.Add(1)
	Walk(t.Left, ch, wg)
	ch <- t.Value
	wg.Add(1)
	Walk(t.Right, ch, wg)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree, wg *sync.WaitGroup) bool {
	c1 := make(chan int, 10)
	c2 := make(chan int, 10)
	a1 := make([]int, 0)
	a2 := make([]int, 0)
	wg.Add(1)
	go Walk(t2, c1, wg)
	wg.Add(1)
	go Walk(t1, c2, wg)
	
	wg.Wait()
	close(c1)
	close(c2)
	for i := range c1 {
		a1 = append(a1, i)
	}
	
	for j := range c2 {
		a2 = append(a2, j)
	}

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
	var wg sync.WaitGroup

	wg.Add(1)

	go Walk(tree.New(2), c, &wg)
	wg.Wait()
	close(c)
	for i := range c {
		fmt.Println(i)
	}
	fmt.Println(Same(tree.New(3), tree.New(4), &wg))
}