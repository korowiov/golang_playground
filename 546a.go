package main

import "fmt"

func main() {
	var k, n, w int
	fmt.Scan(&k, &n, &w)

	sum := k * ((w * (1 + w)) / 2)
	if n < sum {
		fmt.Print(sum - n)
	} else {
		fmt.Print(0)
	}
}
