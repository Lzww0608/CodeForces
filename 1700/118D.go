package main

import (
	"bufio"
	"fmt"
	"os"
)

const MOD = 1e8

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n1, n2, k1, k2 int
	fmt.Fscan(in, &n1, &n2, &k1, &k2)
	f := make([][]int, n1+1)
	g := make([][]int, n1+1)
	for i := 0; i <= n1; i++ {
		f[i] = make([]int, n2+1)
		g[i] = make([]int, n2+1)
	}

	f[0][0], g[0][0] = 1, 1

	for i := 0; i <= n1; i++ {
		for j := 0; j <= n2; j++ {
			for x := 1; x <= min(k1, n1) && i+x <= n1; x++ {
				f[i+x][j] = (f[i+x][j] + g[i][j]) % MOD
			}
			for x := 1; x <= min(k2, n2) && j+x <= n2; x++ {
				g[i][j+x] = (g[i][j+x] + f[i][j]) % MOD
			}
		}
	}

	fmt.Fprintln(out, (f[n1][n2]+g[n1][n2]) % MOD)
}
