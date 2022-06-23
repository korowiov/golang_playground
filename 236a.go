package main

import (
	"fmt"
)

func main() {
	var s string
	fmt.Scan(&s)
	var set = map[rune]bool{}
	for _, chr := range s {
		set[chr] = true
	}
	if len(set)%2 == 0 {
		fmt.Println("CHAT WITH HER!")
	} else {
		fmt.Println("IGNORE HIM!")
	}
}
