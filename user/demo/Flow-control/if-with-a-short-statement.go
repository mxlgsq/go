package main

import (
	"fmt"
	"math"
)
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	// fmt.Println(v) // undefined: v
	return lim
}
func main() {
	fmt.Println(
		pow(3, 2, 8),
		pow(3, 2, 20),
	)
}