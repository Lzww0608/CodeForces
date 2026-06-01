package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const MOD int = 998244353

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		fmt.Fscan(in, &a[i])
	}

	F := make([]int, n+1)
	invF := make([]int, n+1)
	F[0] = 1
	for i := 1; i < len(F); i++ {
		F[i] = F[i-1] * i % MOD
	}
	invF[n] = quickPow(F[n], MOD-2)
	for i := n - 1; i >= 0; i-- {
		invF[i] = invF[i+1] * (i + 1) % MOD
	}

	ans := F[n/3] * invF[n/6] % MOD * invF[n/6] % MOD
	for i := 0; i < n; i += 3 {
		sort.Ints(a[i : i+3])
		a, b, c := a[i], a[i+1], a[i+2]
		if a == b && b == c {
			ans = ans * 3 % MOD
		} else if a == b {
			ans = ans * 2 % MOD
		}
	}

	fmt.Fprintln(out, ans)
}

func quickPow(a, b int) int {
	res := 1
	for b > 0 {
		if b&1 == 1 {
			res = res * a % MOD
		}
		a = a * a % MOD
		b >>= 1
	}
	return res
}
