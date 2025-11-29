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

const N = 31

func solve(in *bufio.Reader, out *bufio.Writer) {
	var n int
	fmt.Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		fmt.Fscan(in, &a[i])
	}

	nxt := make([][N]int, n+1)
	for j := range N {
		nxt[n][j] = n
	}

	for i := n - 1; i >= 0; i-- {
		nxt[i] = nxt[i+1]
		for j := N - 1; j >= 0; j-- {
			if (a[i]>>j)&1 == 0 {
				nxt[i][j] = i
			}
		}
	}

	var q int
	for fmt.Fscan(in, &q); q > 0; q-- {
		var l, k int
		fmt.Fscan(in, &l, &k)
		l--
		if k > a[l] {
			fmt.Fprintf(out, "%d ", -1)
			continue
		}
		ans, res := l, n
		for i := N - 1; i >= 0; i-- {
			if (k>>i)&1 == 1 {
				res = min(res, nxt[l][i])
			} else {
				ans = max(ans, min(res, nxt[l][i]))
			}
		}
		ans = max(ans, res)
		fmt.Fprintf(out, "%d ", ans)
	}
	fmt.Fprintln(out)
}
