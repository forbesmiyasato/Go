package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	y := make([][]uint8, dy)
	x := make([]uint8, dx)
	
	for i := range y {
		for j := range x {
			x[j] = uint8(i ^ j)
		}
		y[i] = x
	}
	
	return y
}

func main() {
	pic.Show(Pic)
}
