package main

import "fmt"

func main() {
	var n, a, b int
	fmt.Scan(&n, &a, &b)
	if n > a*b {
		fmt.Println(-1)
		return
	}

	x, y := 1, 2
	for i := 0; i < a; i++ {
		for j := 0; j < b; j++ {
			if i&1 == 0 {
				if j&1 == 0 {
					if x <= n {
						fmt.Print(x, " ")
						x += 2
					} else {
						fmt.Print(0, " ")
					}
				} else {
					if y <= n {
						fmt.Print(y, " ")
						y += 2
					} else {
						fmt.Print(0, " ")
					}
				}
			} else {
				if j&1 == 1 {
					if x <= n {
						fmt.Print(x, " ")
						x += 2
					} else {
						fmt.Print(0, " ")
					}
				} else {
					if y <= n {
						fmt.Print(y, " ")
						y += 2
					} else {
						fmt.Print(0, " ")
					}
				}
			}

		}
		fmt.Println()
	}
}
