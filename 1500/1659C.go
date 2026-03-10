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
	var n, a, b int
	fmt.Fscan(in, &n, &a, &b)
	c := make([]int, n)
	sum := 0
	for i := range c {
		fmt.Fscan(in, &c[i])
		sum += c[i]
	}

	ans := math.MaxInt
	cur, pre := 0, 0
	for i := range n {
		ans = min(ans, cur+b*(sum-(n-i)*pre))
		cur += (a + b) * (c[i] - pre)
		sum -= c[i]
		pre = c[i]
	}

	fmt.Fprintln(out, ans)
}
