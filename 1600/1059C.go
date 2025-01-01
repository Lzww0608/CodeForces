package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	for k := 1; n > 0; k *= 2 {
		if n == 3 {
			fmt.Print(k, k, 3*k)
			return
		}
		for i := 0; i < (n+1)/2; i++ {
			fmt.Print(k, " ")
		}
		n >>= 1
	}

}
