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
	var n, x, m int
	fmt.Fscan(in, &n, &x, &m)

	var a, b int
	l, r := x, x
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &a, &b)
		if b < l || a > r {
			continue
		}
		l, r = min(l, a), max(r, b)
	}

	fmt.Fprintln(out, r-l+1)
}
