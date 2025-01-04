package main

import "fmt"

func main() {
	var a, b, r int
	fmt.Scan(&a, &b, &r)
	if r*2 > a || r*2 > b {
		fmt.Println("Second")
	} else {
		fmt.Println("First")
	}
}
