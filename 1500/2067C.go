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

	check := func(mid int) bool {
		t := n - mid
		for t > 0 {
			x := t % 10
			t /= 10

			if ((7 - x + 10) % 10) <= mid {
				return true
			}
		}

		return false
	}

	if n <= 7 {
		fmt.Fprintln(out, (9-(7-n))%9)
	} else {
		l, r := -1, 7
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
}
