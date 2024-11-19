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
	var n int
	fmt.Fscan(in, &n)
	var s string
	fmt.Fscan(in, &s)
	ans, mn := 1, s
	for i := 2; i <= n; i++ {
		cur := make([]byte, 0, n)
		for j := i - 1; j < n; j++ {
			cur = append(cur, s[j])
		}
		if (n-i+1)&1 == 1 {
			for j := i - 2; j >= 0; j-- {
				cur = append(cur, s[j])
			}
		} else {
			for j := 0; j < i-1; j++ {
				cur = append(cur, s[j])
			}
		}

		t := string(cur)
		if t < mn {
			mn = t
			ans = i
		}
	}
	fmt.Fprintln(out, mn)
	fmt.Fprintln(out, ans)
}
