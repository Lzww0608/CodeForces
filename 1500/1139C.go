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
	fmt.Fscan(in, &n, &k)
	fa := make([]int, n)
	sz := make([]int, n)
	for i := range fa {
		fa[i] = i
		sz[i] = 1
	}

	var find func(int) int
	find = func(x int) int {
		if fa[x] != x {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}

	merge := func(x, y int) {
		rx, ry := find(x), find(y)
		if rx == ry {
			return
		}

		if sz[rx] < sz[ry] {
			rx, ry = ry, rx
		}
		sz[rx] += sz[ry]
		fa[ry] = rx
	}

	for range n - 1 {
		var u, v, x int
		fmt.Fscan(in, &u, &v, &x)
		u--
		v--
		if x == 0 {
			merge(u, v)
		}
	}

	ans := quickPow(n, k)
	for i := range n {
		if find(i) == i {
			ans = (ans - quickPow(sz[i], k) + MOD) % MOD
		}
	}

	fmt.Fprintln(out, ans)
}

func quickPow(a, r int) int {
	if a == 1 {
		return 1
	}
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
