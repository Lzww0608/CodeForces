package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	a := make([]int, n)
	cnt := make(map[int]int)
	for i := range a {
		fmt.Fscan(in, &a[i])
		cnt[a[i]]++
	}
	sort.Ints(a)
	b := make([]int, 0, n)
	for i := 0; i < n; i++ {
		if i == 0 || a[i] != a[i-1] {
			b = append(b, a[i])
		}
	}
	n = len(b)
	//fmt.Fprintln(out, b)

	ans := 0
	sum := 1
	for l, r := 0, 0; l < n; l++ {
		for r < n && b[r]-b[l] < m {
			sum = (sum * cnt[b[r]]) % MOD
			r++
		}
		if r-l == m {
			ans = (ans + sum) % MOD
		}
		sum = (sum * modInverse(cnt[b[l]])) % MOD
	}
	fmt.Fprintln(out, ans)
}

func modInverse(x int) int {
	return quickPow(x, MOD-2)
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
