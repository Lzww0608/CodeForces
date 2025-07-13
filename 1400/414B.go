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

	var n, k int
	fmt.Fscan(in, &k, &n)

	div := make([][]int, k+1)
	for i := 1; i <= k; i++ {
		for j := i; j <= k; j += i {
			div[j] = append(div[j], i)
		}
	}

	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, k+1)
	}
	for j := 1; j <= k; j++ {
		f[1][j] = 1
	}

	for i := 2; i <= n; i++ {
		for j := 1; j <= k; j++ {
			for _, x := range div[j] {
				f[i][j] = (f[i][j] + f[i-1][x]) % MOD
			}
		}
	}

	ans := 0
	for _, x := range f[n] {
		ans = (ans + x) % MOD
	}

	fmt.Fprintln(out, ans)
}
