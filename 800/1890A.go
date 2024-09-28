package main

import (
	"fmt"
)

func main() {
	var t, n int
	for fmt.Scan(&t); t > 0; t-- {
		fmt.Scan(&n)
		var x, cnt int
		a, b := -1, -1
		f := true
		for i := 0; i < n; i++ {
			fmt.Scan(&x)
			if a == -1 {
				a = x
			} else if a == x {
				cnt++
			} else if b == -1 {
				b = x
			} else if b == x {
				cnt--
			} else {
				f = false
			}
		}
		if !f || (a != -1 && b != -1 && cnt != 1 && cnt != 0 && cnt != -1) {
			fmt.Println("No")
		} else {
			fmt.Println("Yes")
		}
	}
}
