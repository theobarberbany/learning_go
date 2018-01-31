package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("Error: %v is negative", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	fmt.Printf("Calculating the sqrt of %v \n", x)
	z := x / 2
	for i := 0; (z*z-x)/(2*z) != 0; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Println(z)
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(245125))
	fmt.Println("Handling error")
	fmt.Println(Sqrt(-2))
}

