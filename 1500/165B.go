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

	check := func(m int) bool {
		sum := 0
		cur, t := m, k
		for sum < n && cur != 0 {
			sum += cur
			cur = m / t
			t *= k
		}

		return sum >= n
	}

	l, r := 0, n
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
