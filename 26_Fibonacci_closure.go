package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func(int) int {
	return func(x int) int {
		if (x == 0) {
			return 0
		} else if (x == 1) {
			return 1
		}
		n1 := 0
		n2 := 1
		n2Prev := 0
		for i := 2; i <= x; i++ {
			n2Prev = n2
			n2+=n1
			n1 = n2Prev
		}
		return n2
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f(i))
	}
}
