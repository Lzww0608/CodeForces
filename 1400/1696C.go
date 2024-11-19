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
	var n, m, k int
	fmt.Fscan(in, &n, &m)
	a := make([]int, n)
	for i := range a {
		fmt.Fscan(in, &a[i])
	}
	fmt.Fscan(in, &k)
	b := make([]int, k)
	for i := range b {
		fmt.Fscan(in, &b[i])
	}
	c := calc(m, a)
	d := calc(m, b)
	if len(c) != len(d) {
		fmt.Fprintln(out, "NO")
		return
	}
	for i := range c {
		if c[i] != d[i] {
			fmt.Fprintln(out, "NO")
			return
		}
	}
	fmt.Fprintln(out, "YES")
	return
}

func calc(m int, a []int) (ans [][2]int) {
	for _, x := range a {
		cnt := 1
		for x%m == 0 {
			x /= m
			cnt *= m
		}

		if len(ans) == 0 || ans[len(ans)-1][0] != x {
			ans = append(ans, [2]int{x, cnt})
		} else {
			ans[len(ans)-1][1] += cnt
		}
	}

	return
}
