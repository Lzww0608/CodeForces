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
	var n, q, l, r int
	fmt.Fscan(in, &n, &q)
	a := make([]int, n)
	pre := make([]int, n+1)
	b := make([]int, n+1)
	for i := range a {
		fmt.Fscan(in, &a[i])
		if a[i] == 1 {
			b[i+1] = b[i] + 2
		} else {
			b[i+1] = b[i] + 1
		}
		pre[i+1] = pre[i] + a[i]
	}

	for ; q > 0; q-- {
		fmt.Fscan(in, &l, &r)
		if l == r || pre[r]-pre[l-1] < b[r]-b[l-1] {
			fmt.Fprintln(out, "NO")
		} else {
			fmt.Fprintln(out, "YES")
		}
	}
}
