package main

import "fmt"

func main() {
	var t, n int
	for fmt.Scan(&t); t > 0; t-- {
		fmt.Scan(&n)
		if n == 1 {
			fmt.Println(4)
		} else if n == 2 {
			fmt.Println(4, 6)
		} else if n == 3 {
			fmt.Println(4, 6, 10)
		} else {
			for i := 4*n - 2; n > 0; n-- {
				fmt.Print(i, " ")
				i -= 2
			}
			fmt.Println()
		}

	}

}
