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

	var t, n, k int

	for fmt.Fscan(in, &t); t > 0; t-- {
		fmt.Fscan(in, &n, &k)
		a := make([][]int, n)
		for i := range a {
			a[i] = make([]int, n)
			for j := range a[i] {
				fmt.Fscan(in, &a[i][j])
			}
		}

		cnt := 0
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if a[i][j] != a[n-i-1][n-j-1] {
					cnt++
				}
			}
		}
		cnt >>= 1
		if cnt <= k && (n&1 == 1 || (k-cnt)&1 == 0) {
			fmt.Fprintln(out, "YES")
		} else {
			fmt.Fprintln(out, "NO")
		}

	}
}
