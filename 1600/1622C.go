package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	a := make([]int, n)
	sum := 0
	for i := range a {
		fmt.Fscan(in, &a[i])
		sum += a[i]
	}
	if sum <= k {
		fmt.Fprintln(out, 0)
		return
	}
	sort.Ints(a)
	ans := sum - k

	for i, s := n-1, a[0]-ans; i > 0; i-- {
		s += a[i]
		x, y := 0, n-i+1
		if s >= 0 {
			x = s / y
		} else {
			x = (s - y + 1) / y
		}
		x = min(x, a[0])
		ans = min(ans, a[0]-x+y-1)
	}

	fmt.Fprintln(out, ans)
}
