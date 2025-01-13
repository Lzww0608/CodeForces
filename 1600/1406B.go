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
	var n int
	fmt.Fscan(in, &n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}
	sort.Ints(a)
	b := []int{1}
	c := []int{1}
	for i := 1; i < min(n, 4); i += 2 {
		b = append(b, a[i]*a[i-1]*b[len(b)-1])
	}
	for i := n - 1; i >= n-5; i-- {
		c = append(c, a[i]*c[len(c)-1])
	}
	ans := c[len(c)-1]
	for i := 1; i <= 5; i += 2 {
		ans = max(ans, c[i]*b[max(0, len(b)-i/2-1)])
	}
	fmt.Fprintln(out, ans)
}
