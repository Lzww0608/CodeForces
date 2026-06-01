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
	var n, x int
	fmt.Fscan(in, &n, &x)
	a := make([]int, n)
	for i := range a {
		fmt.Fscan(in, &a[i])
	}
	b := make([]bool, n)
	b[n-1] = true
	for i := n - 2; i >= 0; i-- {
		if a[i] <= a[i+1] {
			b[i] = true
		} else {
			break
		}
	}

	ans := 0
	l, r := 0, 0
	for {
		r = l
		for r < n && a[r] <= x {
			r++
		}

		for k := l; k < r; k++ {
			if k == 0 {
				continue
			}

			if a[k-1] > a[k] {
				fmt.Fprintln(out, -1)
				return
			}
		}

		if r == n || b[r] {
			break
		}

		ans++
		a[r], x = x, a[r]
		l = r + 1
	}

	fmt.Fprintln(out, ans)
}
