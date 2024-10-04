package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var n, k, x, t int
	for fmt.Fscan(in, &t); t > 0; t-- {
		fmt.Fscan(in, &n, &k, &x)
		if x != 1 {
			fmt.Println("YES")
			fmt.Println(n)
			for i := 0; i < n; i++ {
				fmt.Print(1, " ")
			}
			fmt.Println()
		} else {
			if k == 1 || (k == 2 && n&1 == 1) {
				fmt.Println("NO")
				continue
			}
			fmt.Println("YES")
			fmt.Println(n / 2)
			if n&1 == 1 {
				fmt.Print(3, " ")
				for i := 0; i < n/2-1; i++ {
					fmt.Print(2, " ")
				}
			} else {
				for i := 0; i < n/2; i++ {
					fmt.Print(2, " ")
				}
			}
			fmt.Println()
		}

	}
}
