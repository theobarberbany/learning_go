package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	fmt.Printf("Calculating the sqrt of %v \n", x)
	z:= x/2
	for i:=0; (z*z - x) / (2 * z) != 0; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Println(z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(9245125))
	fmt.Println("Comparing to math.Sqrt")
	fmt.Println(math.Sqrt(9245125))
}

