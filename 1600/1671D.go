package main

import (
	"bufio"
	"fmt"
	"io"
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

func solve(in io.Reader, out io.Writer) {
	var n, x int
	fmt.Fscan(in, &n, &x)
	a := make([]int, n)
	mx, mn := math.MinInt32, math.MaxInt32
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
		mx = max(mx, a[i])
		mn = min(mn, a[i])
	}

	ans := 0
	t1, tx := min(a[0]-1, a[n-1]-1), min(x-a[0], x-a[n-1])
	for i := 1; i < n; i++ {
		ans += abs(a[i] - a[i-1])
		t1 = min(t1, a[i]-1+a[i-1]-1-abs(a[i]-a[i-1]))
		tx = min(tx, x-a[i]+x-a[i-1]-abs(a[i]-a[i-1]))
	}

	if x > mx {
		ans += tx
	}
	if mn > 1 {
		ans += t1
	}
	fmt.Fprintln(out, ans)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
