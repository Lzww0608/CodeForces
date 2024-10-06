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
	a := make([]int, n)
	b := make([]int, n)
	for i := range a {
		fmt.Fscan(in, &a[i])
		b[i] = a[i] + i
		a[i] -= i
	}
	mn, mx := 0, 0
	for i := 1; i < n; i++ {
		if a[i] >= a[mn]+1 {
			fmt.Fprintln(out, "YES")
			fmt.Fprintln(out, mn+1, i+1)
			return
		} else {
			mn = i
		}

		if b[mx] >= b[i]+1 {
			fmt.Fprintln(out, "YES")
			fmt.Fprintln(out, mx+1, i+1)
			return
		} else {
			mx = i
		}
	}
	fmt.Fprintln(out, "NO")
}
