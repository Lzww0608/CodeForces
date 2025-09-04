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
	a := make([]int, m)
	cnt := make([]int, n)
	for i := range a {
		fmt.Fscan(in, &a[i])
		cnt[a[i]-1]++
	}

	check := func(mid int) bool {
		sum := 0
		for i := 0; i < n; i++ {
			sum += min(cnt[i], mid) + max(0, (mid-cnt[i])/2)
		}

		return sum >= m
	}

	l, r := 0, m*2+1
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
