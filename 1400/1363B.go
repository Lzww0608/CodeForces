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

func solve(in *bufio.Reader, out *bufio.Writer) {
	var s string
	fmt.Fscan(in, &s)
	n := len(s)

	pre := make([]int, n+1)
	suf := make([]int, n+1)
	for i := 0; i < n; i++ {
		if s[i] == '0' {
			pre[i+1] = pre[i] + 1
		} else {
			pre[i+1] = pre[i]
		}

		if s[n-1-i] == '0' {
			suf[n-i-1] = suf[n-i] + 1
		} else {
			suf[n-i-1] = suf[n-i]
		}
	}

	ans := n
	for i := 0; i <= n; i++ {
		ans = min(ans, min(pre[i], i-pre[i])+min(suf[i], n-i-suf[i]))
	}

	fmt.Fprintln(out, ans)
}
