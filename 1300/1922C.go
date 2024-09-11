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

	var t, n, m, l, r int
	for fmt.Fscan(in, &t); t > 0; t-- {
		fmt.Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			fmt.Fscan(in, &a[i])
		}
		pre := make([]int, n+1)
		suf := make([]int, n+1)
		pre[1] = 1
		for i := 1; i < n; i++ {
			if i == n-1 || a[i+1]-a[i] < a[i]-a[i-1] {
				pre[i+1] = pre[i] + 1
			} else {
				pre[i+1] = pre[i] + a[i+1] - a[i]
			}
		}

		suf[n-1] = 1
		for i := n - 2; i >= 0; i-- {
			if i == 0 || a[i+1]-a[i] > a[i]-a[i-1] {
				suf[i] = suf[i+1] + 1
			} else {
				suf[i] = suf[i+1] + a[i] - a[i-1]
			}

		}

		//fmt.Fprintln(out, pre)
		fmt.Fscan(in, &m)
		for i := 0; i < m; i++ {
			fmt.Fscan(in, &l, &r)
			if l > r {
				fmt.Fprintln(out, suf[r]-suf[l])
			} else {
				fmt.Fprintln(out, pre[r-1]-pre[l-1])
			}
		}
	}
}
