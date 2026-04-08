package main

import (
	"bufio"
	"fmt"
	"os"
)

const MOD int = 1_000_000_007
const N int = 200_001

var F [N]int

func init() {
	F[0] = 1
	for i := 1; i < N; i++ {
		F[i] = F[i-1] * i % MOD
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	for fmt.Fscan(in, &t); t > 0; t-- {
		solve(in, out)
	}
}

func quickPow(a, r int) int {
	res := 1
	for r > 0 {
		if r&1 == 1 {
			res = res * a % MOD
		}
		a = a * a % MOD
		r >>= 1
	}
	return res
}

func comb(n, k int) int {
	if n < k {
		return 0
	}

	return F[n] * quickPow(F[k], MOD-2) % MOD * quickPow(F[n-k], MOD-2) % MOD
}

func solve(in *bufio.Reader, out *bufio.Writer) {
	var n, k, x int
	fmt.Fscan(in, &n, &k)
	a, b := 0, 0
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &x)
		if x == 1 {
			a++
		} else {
			b++
		}
	}

	mid := (k + 1) / 2
	if a < mid {
		fmt.Fprintln(out, 0)
		return
	}

	ans := 0
	for i := mid; i <= a && i <= k; i++ {
		ans = (ans + comb(a, i)*comb(b, k-i)) % MOD
	}
	fmt.Fprintln(out, ans)
}
