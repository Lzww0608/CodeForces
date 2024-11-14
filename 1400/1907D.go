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
	type pair struct {
		l, r int
	}
	a := make([]pair, n)
	for i := range a {
		fmt.Fscan(in, &a[i].l, &a[i].r)
	}

	check := func(k int) bool {
		L, R := 0, 0
		for _, v := range a {
			L -= k
			R += k
			L = max(L, v.l)
			R = min(R, v.r)
			if L > R {
				return false
			}
		}
		return true
	}

	l, r := 0, 1_000_000_001
	for l < r {
		mid := l + ((r - l) >> 1)
		if check(mid) {
			r = mid
		} else {
			l = mid + 1
		}
		//fmt.Fprint(out, l, ",", r, " ")
	}
	fmt.Fprintln(out, l)
}
