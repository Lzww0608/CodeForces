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
	var s string
	fmt.Fscan(in, &s)
	a := make([]int, n)

	l, r := -1, 0
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
		r = max(r, a[i])
	}

	check := func(mid int) bool {
		cnt := 0
		f := true
		for i := range s {
			if s[i] == 'R' {
				if a[i] > mid {
					f = true
				}
			} else {
				if a[i] > mid && f {
					f = false
					cnt++
				}
			}
		}

		return cnt <= k
	}

	for l+1 < r {
		mid := l + ((r - l) >> 1)
		if check(mid) {
			r = mid
		} else {
			l = mid
		}
	}

	fmt.Fprintln(out, r)
}
