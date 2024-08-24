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
		a := make([][]int, 2)
		for i := range a {
			a[i] = make([]int, n)
		}
		l, r := n, n*2-1
		for j := 0; j < n; j++ {
			for i := 0; i < 2; i++ {
				if (i+j)&1 == 0 {
					a[i][j] = r
					r--
				} else {
					a[i][j] = l
					l--
				}

			}
		}
		a[1][n-1] = n * 2
		for i := range a {
			for _, x := range a[i] {
				fmt.Fprint(out, x, " ")
			}
			fmt.Fprintln(out)
		}
	}
}
