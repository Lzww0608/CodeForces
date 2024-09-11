package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t, n int
	for fmt.Fscan(in, &t); t > 0; t-- {
		fmt.Fscan(in, &n)
		l, r := 1, n*n
		a := make([][]int, n)
		for i := range a {
			a[i] = make([]int, n)
		}
		for k := 1; k <= n*n; k++ {
			i := k - 1
			if k&1 == 1 {
				if (i/n)&1 == 0 {
					a[i/n][i%n] = r
				} else {
					a[i/n][n-1-i%n] = r
				}
				r--
			} else {
				if (i/n)&1 == 0 {
					a[i/n][i%n] = l
				} else {
					a[i/n][n-1-i%n] = l
				}
				l++
			}

		}

		for i := range a {
			for j := range a[i] {
				fmt.Fprint(out, a[i][j], " ")
			}
			fmt.Fprintln(out)
		}
	}
}
