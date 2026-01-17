package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

// 1730B
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
	for i := range a {
		fmt.Fscan(in, &a[i])
	}
	b := make([]int, n)
	for i := range b {
		fmt.Fscan(in, &b[i])
	}

	mn, mx := math.MaxInt32, math.MinInt32
	for i := range n {
		mn = min(mn, -b[i]+a[i])
		mx = max(mx, b[i]+a[i])
	}

	fmt.Fprintln(out, float64(mn+mx)/2)
}
