package main

import (
	"bufio"
	"fmt"
	"os"
)

// https://codeforces.com/problemset/problem/448/D
func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m, k int
	fmt.Fscan(in, &n, &m, &k)

	check := func(mid int) bool {
		cnt := 0
		for i := range n {
			cnt += min(m, mid/(i+1))
		}

		return cnt >= k
	}

	l, r := 0, m*n+1
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
