package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64
	fmt.Scan(&a, &b, &c)

	var s = (a + b + c) / 2
    var area float64 = math.Sqrt(s * (s - a) * (s - b) * (s - c))
	fmt.Printf("%.2f\n", area)
}
