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

	var n, m int
	fmt.Fscan(in, &n, &m)
	f := make([]int, n+1)
	sum := 0
	for i := 1; i <= n; i++ {
		f[i] = 1
		sum += f[i]
	}

	for i := 1; i < m*2; i++ {
		tmp := 0
		for j := n; j >= 1; j-- {
			f[j], sum = sum, (sum-f[j]+MOD)%MOD
			f[j] %= MOD
			tmp += f[j]
			tmp %= MOD
		}
		sum = tmp
	}

	ans := 0
	for _, x := range f {
		ans = (ans + x) % MOD
	}

	fmt.Fprintln(out, ans)
}
