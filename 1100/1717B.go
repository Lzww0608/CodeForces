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

	var t, n, k, r, c int
	for fmt.Fscan(in, &t); t > 0; t-- {
		fmt.Fscan(in, &n, &k, &r, &c)
		a := make([][]byte, n)
		for i := range a {
			a[i] = make([]byte, n)
			for j := range a[i] {
				a[i][j] = '.'
			}
		}
		a[r-1][c-1] = 'X'
		for i := 0; i < n; i++ {
			t := ((i - (r - 1)) + k*n + (c-1)%k) % k
			for j := 0; j < n; j++ {
				if j%k == t {
					a[i][j] = 'X'
				}
			}
		}

		for i := range a {
			fmt.Fprintln(out, string(a[i]))
		}
	}
}
