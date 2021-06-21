package main

import (
	"fmt"
	"math"
)

const threshold float64 = 1e-10

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("Can not Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
    z := 1.0
    for i := 0; i < 10; i++ {
        if math.Abs((z*z - x) / (2*z)) < threshold {
            return z, nil
        } else {
            z -= (z*z - x) / (2*z)
        }
    }
    return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}

