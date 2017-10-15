package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	const sigma float64 = 1.0 / float64(1<<50)
	z := x / 2.0
	znext:= z-(z*z-x)/(2*z)
	i:= 0
	
	for math.Abs(z-znext)>sigma {
		z = znext
		znext = z-(z*z-x)/(2*z)
		i++
	}
	fmt.Println(i)
	return z
}

func main() {
	s := 2.0
	fmt.Println(Sqrt(s))
	fmt.Println(math.Sqrt(s))
}
