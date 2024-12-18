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
	b := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i], &b[i])
	}

	c := make([]int, n)
	ans, mn := 0, math.MaxInt
	for i := 0; i < n; i++ {
		c[i] = max(0, a[i]-b[(i-1+n)%n])
		ans += c[i]
		mn = min(mn, a[i]-c[i])
	}
	fmt.Fprintln(out, ans+mn)
}
