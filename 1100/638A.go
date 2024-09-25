package main

import "fmt"

func main() {
	var n, a int
	fmt.Scan(&n, &a)
	if a&1 == 1 {
		fmt.Println(a/2 + 1)
	} else {
		fmt.Println((n-a)/2 + 1)
	}
}
