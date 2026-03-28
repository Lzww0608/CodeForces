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

const inf = int(1e18)

func solve(in *bufio.Reader, out *bufio.Writer) {
	var n int
	fmt.Fscan(in, &n)

	var l, r, c int
	minL, minR, mn := inf, inf, inf
	curL, curR := 1<<30, -1
	for range n {
		fmt.Fscan(in, &l, &r, &c)

		if l < curL {
			curL, minL = l, c
			mn = inf
		} else if l == curL {
			minL = min(minL, c)
		}

		if r > curR {
			curR, minR = r, c
			mn = inf
		} else if r == curR {
			minR = min(minR, c)
		}

		if l == curL && r == curR {
			mn = min(mn, c)
		}

		ans := min(minL+minR, mn)
		fmt.Fprintln(out, ans)
	}
}
