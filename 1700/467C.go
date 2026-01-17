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

	var n, m, k int
	fmt.Fscan(in, &n, &m, &k)
	a := make([]int, n)
	sum := make([]int, n+1)
	for i := range a {
		fmt.Fscan(in, &a[i])
		sum[i+1] = sum[i] + a[i]
	}

	f := make([][]int, k+1)
	for i := range f {
		f[i] = make([]int, n+1)
	}

	for i := 1; i <= k; i++ {
		for j := m; j <= n; j++ {
			f[i][j] = max(f[i][j-1], f[i-1][j-m]+sum[j]-sum[j-m])
		}
	}

	fmt.Fprintln(out, f[k][n])
}
