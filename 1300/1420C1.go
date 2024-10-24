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
	var n, q int
	fmt.Fscan(in, &n, &q)
	a := make([]int, n)
	for i := range a {
		fmt.Fscan(in, &a[i])
	}
	f := make([]int, n+1)
	g := make([]int, n+1)
	for i := range a {
		f[i+1], g[i+1] = max(f[i], a[i], g[i]+a[i]), max(g[i], -a[i], f[i]-a[i])
	}
	fmt.Fprintln(out, max(g[n], f[n]))
}
