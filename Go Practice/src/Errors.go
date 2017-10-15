package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64 

func (f ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(f))
}

func Sqrt(x float64) (float64, error) {
	if x < 0.0 {
		return -1, ErrNegativeSqrt(x)
	}
	const sigma float64 = 1.0 / float64(1<<50)
	z := x / 2.0
	znext := z - (z*z-x)/(2*z)
	for math.Abs(z-znext) > sigma {
		z = znext
		znext = z - (z*z-x)/(2*z)
	}
	// fmt.Println(i)
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}