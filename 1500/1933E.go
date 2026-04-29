package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	a := make([]int, n)
	pre := make([]int, n+1)
	for i := range a {
		fmt.Fscan(in, &a[i])
		pre[i+1] = pre[i] + a[i]
	}

	var q, l, u int
	for fmt.Fscan(in, &q); q > 0; q-- {
		fmt.Fscan(in, &l, &u)

		r := sort.SearchInts(pre, pre[l-1]+u)
		if r > n {
			fmt.Fprintf(out, "%d ", n)
			continue
		} else if r == l {
			fmt.Fprintf(out, "%d ", l)
			continue
		}
		x := pre[r] - pre[l-1]
		y := pre[r-1] - pre[l-1]
		if x*u-x*(x-1)/2 > y*u-y*(y-1)/2 {
			fmt.Fprintf(out, "%d ", r)
		} else {
			fmt.Fprintf(out, "%d ", r-1)
		}
	}

	fmt.Fprintln(out)
}
