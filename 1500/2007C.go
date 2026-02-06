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
	var n, a, b int
	fmt.Fscan(in, &n, &a, &b)
	t := a
	if a != b {
		t = gcd(a, b)
	}
	c := make([]int, n)
	for i := range c {
		fmt.Fscan(in, &c[i])
		c[i] %= t
	}
	sort.Ints(c)

	ans := c[n-1] - c[0]
	for i := 0; i < n-1; i++ {
		ans = min(ans, c[i]+t-c[i+1])
	}

	fmt.Fprintln(out, ans)
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}

	return a
}
