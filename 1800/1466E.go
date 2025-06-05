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
	var n int
	fmt.Fscan(in, &n)
	a := make([]int, n)
	cnt := [60]int{}
	for i := range a {
		fmt.Fscan(in, &a[i])
		for j := 0; j < 60; j++ {
			cnt[j] += (a[i] >> j) & 1
		}
	}

	ans := 0
	for _, x := range a {
		a, b := 0, 0
		for j := 0; j < 60; j++ {
			t := (1 << j) % MOD
			if (x>>j)&1 == 1 {
				a = (a + t*cnt[j]) % MOD
				b = (b + t*n) % MOD
			} else {
				b = (b + t*cnt[j]) % MOD
			}
		}
		ans = (ans + a*b%MOD) % MOD
	}

	fmt.Fprintln(out, ans)
}
