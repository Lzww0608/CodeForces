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
	var n, k int
	fmt.Fscan(in, &n, &k)
	a := make([]int, n)
	for i := range a {
		fmt.Fscan(in, &a[i])
	}

	mn, mx := a[0], a[n-1]
	f := make([][]int, mx+1)
	for i, x := range a {
		f[x] = append(f[x], i)
	}

	ans := mx - mn
	for d := mx; d >= 0; d-- {
		ans = min(ans, d-mn)
		for _, i := range f[d] {
			if d == a[i]/k {
				fmt.Fprintln(out, ans)
				return
			}
			v := a[i] / (a[i]/d + 1)
			f[v] = append(f[v], i)
			mn = min(mn, v)
		}
		f = f[:len(f)-1]
	}
}
