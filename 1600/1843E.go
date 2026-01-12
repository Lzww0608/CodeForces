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
	var n, m int
	fmt.Fscan(in, &n, &m)
	left := make([]int, m)
	right := make([]int, m)
	for i := range m {
		fmt.Fscan(in, &left[i], &right[i])
	}

	var q int
	fmt.Fscan(in, &q)
	a := make([]int, q)
	for i := range a {
		fmt.Fscan(in, &a[i])
	}

	check := func(mid int) bool {
		pre := make([]int, n+1)
		for i := range mid + 1 {
			pre[a[i]]++
		}
		for i := 1; i <= n; i++ {
			pre[i] += pre[i-1]
		}

		for i := range m {
			x, y := left[i]-1, right[i]-1
			if (pre[y+1]-pre[x])*2 > (y - x + 1) {
				return true
			}
		}

		return false
	}

	l, r := -1, q
	for l+1 < r {
		mid := l + ((r - l) >> 1)
		if check(mid) {
			r = mid
		} else {
			l = mid
		}
	}

	if r == q {
		fmt.Fprintln(out, -1)
	} else {
		fmt.Fprintln(out, r+1)
	}
}
