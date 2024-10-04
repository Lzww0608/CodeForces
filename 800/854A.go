package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	if n&1 == 1 {
		fmt.Println(n>>1, n-n/2)
	} else if (n/2)&1 == 1 {
		fmt.Println(n/2-2, n/2+2)
	} else {
		fmt.Println(n/2-1, n/2+1)
	}

}
