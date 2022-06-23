package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b float64
	fmt.Scan(&a, &b)
	fmt.Print(math.Floor(math.Log2(b/a)/math.Log2(1.5) + 1))
}
