package main

import (
    "fmt"
    "math"
)

const threshold float64 = 1e-10

// Sqrt - Approximates math.Sqrt function
func Sqrt(x float64) float64 {
    z := 1.0
    for i := 0; i < 10; i++ {
        if math.Abs((z*z - x) / (2*z)) < threshold {
            return z
        } else {
            z -= (z*z - x) / (2*z)
        }
    }
    return z
}

func Compare(x float64) {
    fmt.Printf("Newton Sqrt(%v): %v\n", x, Sqrt(x))
    fmt.Printf("math.Sqrt(%v): %v\n", x, math.Sqrt(x))
}

func main() {
    fmt.Println("x = 2.0")
    Compare(2.0)
    fmt.Println("x = 20.0")
    Compare(20.0)
    fmt.Println("x = 200.0")
    Compare(200.0)
}
