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
	var s string
	fmt.Fscan(in, &s)
	n := len(s)

	a := []int{}
	cnt := 1
	for i := 1; i < n; i++ {
		if s[i] == s[i-1] {
			cnt++
		} else {
			a = append(a, cnt)
			cnt = 1
		}
	}
	a = append(a, cnt)

	ans := 1
	for i := 1; i <= n-len(a); i++ {
		ans = ans * i % MOD
	}

	for _, x := range a {
		ans = ans * x % MOD
	}

	fmt.Fprintln(out, n-len(a), ans)
}
