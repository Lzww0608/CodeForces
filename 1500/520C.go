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
	var s string
	fmt.Fscan(in, &s)
	m := map[byte]int{}
	for i := 0; i < n; i++ {
		m[s[i]]++
	}
	mx := 0
	cnt := 0
	for _, v := range m {
		if mx == v {
			cnt++
		} else if mx < v {
			cnt = 1
			mx = v
		}
	}
	fmt.Fprintln(out, quickPow(cnt, n))
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
