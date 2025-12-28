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

	var n int
	fmt.Fscan(in, &n)
	a := make([]int, n)
	b := make([]int, n)
	pre := make([]int, n+1)
	for i := range n {
		fmt.Fscan(in, &a[i])
	}
	for i := range n {
		fmt.Fscan(in, &b[i])
		pre[i+1] = pre[i] + a[i]*b[i]
	}

	ans := pre[n]
	for i := range n {
		cur := pre[n]
		l, r := i-1, i+1
		for l >= 0 && r < n {
			cur -= a[l]*b[l] + a[r]*b[r]
			cur += a[l]*b[r] + a[r]*b[l]
			ans = max(ans, cur)
			l--
			r++
		}

	}

	for i := range n {
		cur := pre[n]
		l, r := i, i+1
		for l >= 0 && r < n {
			cur -= a[l]*b[l] + a[r]*b[r]
			cur += a[l]*b[r] + a[r]*b[l]
			ans = max(ans, cur)
			l--
			r++
		}

	}

	fmt.Fprintln(out, ans)
}
