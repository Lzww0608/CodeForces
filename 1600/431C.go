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

	var n, k, d int
	fmt.Fscan(in, &n, &k, &d)
	f := make([][2]int, n+1)
	f[0] = [2]int{1, 0}
	for i := 1; i <= n; i++ {
		for j := 1; j <= min(i, k); j++ {
			if j < d {
				f[i][0] += f[i-j][0]
				f[i][1] += f[i-j][1]
			} else {
				f[i][1] += f[i-j][0] + f[i-j][1]
			}
			f[i][0] %= MOD
			f[i][1] %= MOD
		}
	}

	fmt.Fprintln(out, f[n][1])
}
