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
	var n, m int
	fmt.Fscan(in, &n, &m)
	p := quickPow(2, n-1, MOD)

	var l, r, x int
	or := 0
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &l, &r, &x)
		or |= x
	}

	ans := or * p % MOD
	fmt.Fprintln(out, ans)
}

func quickPow(a, r, mod int) int {
	res := 1
	for r > 0 {
		if r&1 != 0 {
			res = (res * a) % mod
		}
		a = (a * a) % mod
		r >>= 1
	}
	return res
}
