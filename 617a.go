package main

import "fmt"

func main() {
	var a int
	var res int

	fmt.Scan(&a)
	res = a / 5
	if a%5 > 0 {
		res += 1
	}
	fmt.Println(res)
}
