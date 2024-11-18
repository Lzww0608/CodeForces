package main

import (
	"bufio"
	"fmt"
	"math"
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
	a := make([]int, n)
	mn, mx := math.MaxInt32, math.MinInt32
	for i := range a {
		fmt.Fscan(in, &a[i])
		mx = max(mx, a[i])
		mn = min(mn, a[i])
	}
	ans := []int{}
	for mn < mx {
		if mn&1 == 1 {
			ans = append(ans, 1)
			mn = (mn + 1) / 2
			mx = (mx + 1) / 2
		} else {
			ans = append(ans, 0)
			mn >>= 1
			mx >>= 1
		}
	}
	fmt.Fprintln(out, len(ans))
	if len(ans) <= n {
		for _, x := range ans {
			fmt.Fprint(out, x, " ")
		}
		fmt.Fprintln(out)
	}

}
