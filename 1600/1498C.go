package main

import (
	"bufio"
	"fmt"
	"os"
)

const MOD int = 1_000_000_007

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	for fmt.Fscan(in, &t); t > 0; t-- {
		solve(in, out)
	}
}

func solve(in *bufio.Reader, out *bufio.Writer) {
	var n, k int
	fmt.Fscan(in, &n, &k)
	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, k+1)
	}
	for i := 1; i <= n; i++ {
		f[i][1] = 1
	}
	for j := 1; j <= k; j++ {
		f[0][j] = 1
	}

	for j := 1; j <= k; j++ {
		for i := 1; i <= n; i++ {
			f[i][j] = (f[i-1][j] + f[n-i][j-1]) % MOD
		}
	}

	fmt.Fprintln(out, f[n][k])
}
