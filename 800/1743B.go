package main

import "fmt"

func main() {
	var t, n int
	for fmt.Scan(&t); t > 0; t-- {
		fmt.Scan(&n)
		fmt.Print(1, " ")
		for i := 3; i <= n; i++ {
			fmt.Print(i, " ")
		}
		fmt.Println(2)
	}

}
