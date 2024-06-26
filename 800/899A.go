package main

import "fmt"

func main() {
	var n, x int
	fmt.Scan(&n)
	one, two := 0, 0
	for i := 0; i < n; i++ {
		fmt.Scan(&x)
		if x == 1 {
			one++
		} else {
			two++
		}
	}
	if two >= one {
		fmt.Println(one)
	} else {
		fmt.Println(two + (one-two)/3)
	}

}
