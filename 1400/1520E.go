package main

import (
	"bufio"
	"fmt"
	"math"
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

	pre := make([]int, n+1)
	suf := make([]int, n+1)

	cnt := 0
	for i := range s {
		if s[i] == '*' {
			cnt++
			pre[i+1] = pre[i]
		} else {
			pre[i+1] = pre[i] + cnt
		}
	}

	cnt = 0
	for i := n - 1; i >= 0; i-- {
		if s[i] == '*' {
			cnt++
			suf[i] = suf[i+1]
		} else {
			suf[i] = suf[i+1] + cnt
		}
	}

	ans := math.MaxInt
	for i := 0; i < n; i++ {
		ans = min(ans, pre[i+1]+suf[i+1])
	}

	fmt.Fprintln(out, ans)
}
