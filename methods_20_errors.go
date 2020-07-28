package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprint("cannot Sqrt negative number:", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if (x < 0) {
		return x, ErrNegativeSqrt(x)
	}
	z := x / 2
	prev := float64(0)
	counter := 0
	/* value stops changing when difference is between 1e-16 to 2e-16 */
	for (math.Abs(z - prev) / z) > 1e-8 {
		prev = z
		z -= (z*z - x) / (2*z)
		counter++
	}
	return z, nil;
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
