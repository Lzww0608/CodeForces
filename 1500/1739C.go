package main

import (
	"bufio"
	"fmt"
	"os"
)

const MOD int = 998244353

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
	var n int
	fmt.Fscan(in, &n)

	ans := calc(n)
	fmt.Fprintln(out, ans[0], ans[1], ans[2])
}

func calc(n int) []int {
	if n == 2 {
		return []int{1, 0, 1}
	}
	a := calc(n - 2)
	return []int{
		(choose(n-1, n/2) + a[1]) % MOD,
		(choose(n-2, n/2) + a[0]) % MOD,
		1,
	}
}

func quickPow(a, r int) int {
	res := 1
	for r > 0 {
		if r&1 == 1 {
			res = (res * a) % MOD
		}
		a = (a * a) % MOD
		r >>= 1
	}
	return res
}

func factor(n int) int {
	res := 1
	for i := 1; i <= n; i++ {
		res = (res * i) % MOD
	}
	return res
}

func choose(n, k int) int {
	if k > n || k < 0 {
		return 0
	}
	// C(n, k) = n! / (k! * (n-k)!)
	a := factor(n)
	b := (factor(k) * factor(n-k)) % MOD
	res := a * quickPow(b, MOD-2) % MOD
	return res
}
