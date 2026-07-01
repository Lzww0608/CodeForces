package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	for fmt.Fscan(in, &t); t > 0; t-- {
		solve(in, out)
	}
}

const MOD int = 1_000_000_007

func solve(in *bufio.Reader, out *bufio.Writer) {
	var s string
	fmt.Fscan(in, &s)
	n := len(s)

	ans := 0
	cur := n
	cnt := make([]int, n*2+1)
	cnt[cur] = 1

	for i := range s {
		if s[i] == '0' {
			cur++
		} else {
			cur--
		}

		ans = (ans + (n-i)*cnt[cur]) % MOD
		cnt[cur] += i + 2
		cnt[cur] %= MOD
	}

	fmt.Fprintln(out, ans)
}
