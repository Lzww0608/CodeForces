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

	var n, k int
	fmt.Fscan(in, &n, &k)
	a := make([]int, n)
	b := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &b[i])
	}

	check := func(mid int) bool {
		cnt := 0
		for i := 0; i < n; i++ {
			if a[i]*mid > b[i] {
				cnt += a[i]*mid - b[i]
				if cnt > k {
					return false
				}
			}
		}

		return true
	}

	l, r := -1, 2_000_000_001
	for l+1 < r {
		mid := l + ((r - l) >> 1)
		if check(mid) {
			l = mid
		} else {
			r = mid
		}
	}

	fmt.Fprintln(out, l)
}
