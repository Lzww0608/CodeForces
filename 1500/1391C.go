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

	var n int
	fmt.Fscan(in, &n)
	ans := 1
	for i := 1; i <= n; i++ {
		ans = ans * i % MOD
	}

	ans = (ans - quickPow(2, n-1) + MOD) % MOD
	fmt.Fprintln(out, ans)
}

func quickPow(a, r int) int {
	res := 1
	for r > 0 {
		if r&1 == 1 {
			res = res * a % MOD
		}
		r >>= 1
		a = a * a % MOD
	}

	return res
}
