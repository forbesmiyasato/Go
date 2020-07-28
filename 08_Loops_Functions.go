package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := x / 2
	prev := float64(0)
	counter := 0
	/* value stops changing when difference is between 1e-16 to 2e-16 */
	for (math.Abs(z - prev) / z) > 1e-8 {
		prev = z
		z -= (z*z - x) / (2*z)
		fmt.Println((math.Abs(z - prev) / z))
		counter++
	}
	
	fmt.Println(counter)
	return z;
}

func main() {
	test := float64(5)
	fmt.Println(Sqrt(test))
	fmt.Println(math.Sqrt(test))
}
